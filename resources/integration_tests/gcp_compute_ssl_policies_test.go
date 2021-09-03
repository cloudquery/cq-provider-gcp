package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-gcp/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationComputeSslPolicies(t *testing.T) {
	testIntegrationHelper(t, resources.ComputeSslPolicies(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.ComputeSslPolicies().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("ssl-policies-policy-%s%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":            fmt.Sprintf("ssl-policies-policy-%s%s", res.Prefix, res.Suffix),
						"description":     "",
						"kind":            "compute#sslPolicy",
						"min_tls_version": "TLS_1_2",
						"profile":         "CUSTOM",
						"custom_features": []interface{}{
							"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
							"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
						},
					},
				},
			},
		}
	})
}
