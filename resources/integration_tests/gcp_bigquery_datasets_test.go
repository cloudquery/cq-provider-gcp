package integration_tests

//todo not enough permissions to finish the task
//func TestIntegrationBigqueryDatasets(t *testing.T) {
//	testIntegrationHelper(t, resources.BigqueryDatasets(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
//		return providertest.ResourceIntegrationVerification{
//			Name: resources.BigqueryDatasets().Name,
//			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
//				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("ds_%s%s", res.Prefix, res.Suffix)})
//			},
//			ExpectedValues: []providertest.ExpectedValue{
//				{
//					Count: 1,
//					Data: map[string]interface{}{
//						"tags": map[string]interface{}{
//							"Type":   "integration_test",
//							"Name":   fmt.Sprintf("nat-%s-%s", res.Prefix, res.Suffix),
//							"TestId": res.Suffix,
//						},
//					},
//				},
//			},
//		}
//	})
//}
