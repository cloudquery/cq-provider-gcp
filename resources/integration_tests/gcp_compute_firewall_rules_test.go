package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-gcp/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationComputeForwardingRules(t *testing.T) {
	testIntegrationHelper(t, resources.ComputeForwardingRules(), []string{"gcp_compute_forwarding_rules.tf", "network.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.ComputeForwardingRules().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("forwarding-rule-%s%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":                   fmt.Sprintf("forwarding-rule-%s%s", res.Prefix, res.Suffix),
						"load_balancing_scheme":  "INTERNAL",
						"is_mirroring_collector": false,
						"ip_protocol":            "TCP",
						"all_ports":              true,
						"allow_global_access":    true,
						"network_tier":           "PREMIUM",
						//"labels": map[string]interface{}{
						//	"test": "test",
						//},
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "gcp_compute_backend_service_backends",
					ForeignKeyName: "backend_service_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"failover":                     false,
								"balancing_mode":               "CONNECTION",
								"max_utilization":              float64(0),
								"max_connections":              float64(0),
								"max_connections_per_endpoint": float64(0),
								"max_connections_per_instance": float64(0),
								"0.0":                          float64(0),
								"capacity_scaler":              float64(0),
								"max_rate_per_endpoint":        float64(0),
								"max_rate_per_instance":        float64(0),
							},
						},
					},
				},
			},
		}
	})
}
