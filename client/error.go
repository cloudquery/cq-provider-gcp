package client

import (
	"errors"
	"net/http"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/execution"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/googleapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func IgnoreErrorHandler(err error) bool {
	var gerr *googleapi.Error
	if ok := errors.As(err, &gerr); ok {
		if gerr.Code == http.StatusForbidden && len(gerr.Errors) > 0 {
			switch gerr.Errors[0].Reason {
			case "accessNotConfigured", "forbidden", "SERVICE_DISABLED":
				return true
			}
		}
	}
	return false
}

type diagValue struct {
	severity diag.Severity
	typ      diag.DiagnosticType
	summary  string
}

var grpcCodeToDiag = map[codes.Code]diagValue{
	codes.PermissionDenied:  {diag.WARNING, diag.ACCESS, "Access denied"},
	codes.Unauthenticated:   {diag.WARNING, diag.ACCESS, "Authentication failure"},
	codes.ResourceExhausted: {diag.WARNING, diag.THROTTLE, "Resource exhausted (quota etc)"},
	codes.Unimplemented:     {diag.IGNORE, diag.RESOLVING, "Operation not implemented or not supported"},
}

var httpCodeToGRPCCode = map[int]codes.Code{
	http.StatusForbidden:       codes.PermissionDenied,
	http.StatusUnauthorized:    codes.Unauthenticated,
	http.StatusTooManyRequests: codes.ResourceExhausted,
	http.StatusNotImplemented:  codes.Unimplemented,
}

func ErrorClassifier(meta schema.ClientMeta, resourceName string, err error) diag.Diagnostics {
	// https://pkg.go.dev/cloud.google.com/go#hdr-Inspecting_errors:
	// Most of the errors returned by the generated clients can be converted into a `grpc.Status`
	if err == nil {
		return nil
	}

	if s, ok := status.FromError(err); ok {
		if v, ok := grpcCodeToDiag[s.Code()]; ok {
			return execution.FromError(err, execution.WithSeverity(v.severity), execution.WithType(v.typ), execution.WithResource(resourceName), execution.WithSummary(v.summary), execution.WithDetails(s.Message()))
		}
	}

	// as a fallback, try to convert the error to *googleapi.Error
	var gerr *googleapi.Error
	if ok := errors.As(err, &gerr); ok {
		if grpcCode, ok := httpCodeToGRPCCode[gerr.Code]; ok {
			if v, ok := grpcCodeToDiag[grpcCode]; ok {
				return execution.FromError(err, execution.WithSeverity(v.severity), execution.WithType(v.typ), execution.WithResource(resourceName), execution.WithSummary(v.summary))
			}
		}
	}

	// failure to classify
	return nil
}
