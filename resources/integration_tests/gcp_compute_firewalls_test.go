package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-gcp/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationComputeFirewalls(t *testing.T) {
	testIntegrationHelper(t, resources.ComputeFirewalls(), []string{"gcp_compute_firewalls.tf", "network.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.ComputeFirewalls().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("google-compute-firewalls-firewall-%s", res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":      fmt.Sprintf("google-compute-firewalls-firewall-%s", res.Suffix),
						"disabled":  false,
						"direction": "INGRESS",
						"source_tags": []interface{}{
							"web",
						},
					},
				},
			},
		}
	})
}
