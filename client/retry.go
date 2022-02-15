package client

import (
	"context"
	"errors"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/googleapis/gax-go/v2"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/api/googleapi"
)

func shouldRetryFunc(log hclog.Logger) func(err error) bool {
	return func(err error) bool {
		if IgnoreErrorHandler(err) {
			reason := ""
			var gerr *googleapi.Error
			if errors.As(err, &gerr) && len(gerr.Errors) > 0 {
				reason = gerr.Errors[0].Reason
			}

			log.Debug("retrier not retrying: ignore error", "err", err, "err_reason", reason)
			return false
		}

		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			log.Debug("retrier not retrying", "err", err)
			return false
		}

		log.Debug("retrying error", "err", err)
		return true
	}
}

// RetryingResolver runs the TableResolver with retry. Not very good as it could cause multiple resources with multi-page retries (retrying after fetching some resources)
func RetryingResolver(f schema.TableResolver) schema.TableResolver {
	return func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
		cl := meta.(*Client)
		return gax.Invoke(ctx, func(ctx context.Context, _ gax.CallSettings) error {
			return f(ctx, meta, parent, res)
		}, gax.WithRetry(func() gax.Retryer {
			return gax.OnErrorFunc(cl.backoff.Gax, shouldRetryFunc(cl.logger))
		}))
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
		return gax.OnErrorFunc(lb.Backoff().Gax, shouldRetryFunc(lb.Logger()))
	}))
	return val, err
}
