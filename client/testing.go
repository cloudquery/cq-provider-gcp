package client

import (
	"testing"

	"github.com/cloudquery/cq-provider-sdk/plugins"
	"github.com/cloudquery/cq-provider-sdk/schema"
)

type TestOptions struct {
	SkipEmptyJsonB bool
}

func GcpMockTestHelper(t *testing.T, table *schema.Table, createService func() (*Services, error), options TestOptions) {
	t.Helper()

	table.IgnoreInTests = false

	providertest.TestResource(t, providertest.ResourceTestCase{
		Provider: &plugins.SourcePlugin{
			Name:    "gcp_mock_test_provider",
			Version: "development",
			// Configure: func(logger hclog.Logger, i interface{}) (schema.ClientMeta, diag.Diagnostics) {
			// 	svc, err := createService()
			// 	if err != nil {
			// 		return nil, diag.FromError(err, diag.INTERNAL)
			// 	}
			// 	c := NewGcpClient(logging.New(&hclog.LoggerOptions{
			// 		Level: hclog.Warn,
			// 	}), BackoffSettings{}, []string{"testProject"}, svc)
			// 	return c, nil
			// },
			Tables: []*schema.Table{
				table,
			},
			// Config: func() provider.Config {
			// 	return &Config{}
			// },
		},
		Config: "",
	})
}
