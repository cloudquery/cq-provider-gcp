package client

import (
	"context"
	"testing"

	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
)

type mockResolver bool

type myErr string

const mockError = myErr("test")

func (m myErr) Error() string { return string(m) }

func (r *mockResolver) resolve(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	*r = true
	return mockError
}

func TestRequireEnabledServices(t *testing.T) {
	tests := []struct {
		name             string
		enabledServices  map[string]map[GcpService]struct{}
		requiredServices []GcpService
		wantResolverCall bool
		wantErr          error
	}{
		{
			"service enabled",
			map[string]map[GcpService]struct{}{"project1": {"service1": struct{}{}}},
			[]GcpService{"service1"},
			true,
			mockError,
		},
		{
			"service disabled",
			map[string]map[GcpService]struct{}{"project1": {}},
			[]GcpService{"service1"},
			false,
			nil,
		},
		{
			"service enabled in another project",
			map[string]map[GcpService]struct{}{"project1": {}, "project2": {"service1": struct{}{}}},
			[]GcpService{"service1"},
			false,
			nil,
		},
		{
			"other service enabled",
			map[string]map[GcpService]struct{}{"project1": {"service2": struct{}{}}},
			[]GcpService{"service1"},
			false,
			nil,
		},
		{
			"one of required services is disabled",
			map[string]map[GcpService]struct{}{"project1": {"service2": struct{}{}}},
			[]GcpService{"service1", "service2"},
			false,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r mockResolver
			f := RequireEnabledServices(r.resolve, tt.requiredServices...)
			cl := Client{
				ProjectId:       "project1",
				EnabledServices: tt.enabledServices,
				logger: logging.New(&hclog.LoggerOptions{
					Level: hclog.Warn,
				}),
			}
			err := f(context.Background(), &cl, nil, nil)
			assert.Equal(t, tt.wantResolverCall, bool(r))
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
