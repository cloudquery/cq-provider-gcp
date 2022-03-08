package client

import (
	"context"
	"errors"
	"net/http"

	"github.com/googleapis/gax-go/v2"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/api/googleapi"
)

func shouldRetryFunc(log hclog.Logger, maxRetries int) func(err error) bool {
	totalRetries := 0
	return func(err error) bool {
		totalRetries += 1
		if totalRetries > maxRetries {
			log.Debug("retrier not retrying, reached max retries", "err", err, "max_retries", maxRetries)
			return false
		}
		var gerr *googleapi.Error
		if ok := errors.As(err, &gerr); ok {
			if gerr.Code == http.StatusForbidden {
				var reason string
				if len(gerr.Errors) > 0 {
					reason = gerr.Errors[0].Reason
				}
				log.Debug("retrier not retrying: ignore error", "err", err, "err_reason", reason, "total_retries", totalRetries, "max_retries", maxRetries)
				return false
			}
		}

		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			log.Debug("retrier not retrying", "err", err, "total_retries", totalRetries, "max_retries", maxRetries)
			return false
		}

		log.Debug("retrying api call", "err", err, "total_retries", totalRetries, "max_retries", maxRetries)
		return true
	}
}

type loggingBackoffer interface {
	Logger() hclog.Logger
	Backoff() BackoffSettings
}

// Retryer runs the given doFunc with retry
func Retryer[T any](ctx context.Context, lb loggingBackoffer, doFunc func(...googleapi.CallOption) (T, error), opts ...googleapi.CallOption) (T, error) {
	var val T
	err := gax.Invoke(ctx, func(ctx context.Context, _ gax.CallSettings) error {
		var err error
		val, err = doFunc(opts...)
		return err
	}, gax.WithRetry(func() gax.Retryer {
		bo := lb.Backoff()
		return gax.OnErrorFunc(bo.Gax, shouldRetryFunc(lb.Logger(), bo.MaxRetries))
	}))
	return val, err
}
