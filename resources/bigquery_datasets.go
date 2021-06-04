package resources

import (
	"context"
	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func BigqueryDatasets() *schema.Table {
	return &schema.Table{
		Name:         "gcp_bigquery_datasets",
		Resolver:     fetchBigqueryDatasets,
		Multiplex:    client.ProjectMultiplex,
		DeleteFilter: client.DeleteProjectFilter,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name: "creation_time",
				Type: schema.TypeBigInt,
			},
			{
				Name:     "default_encryption_configuration_kms_key_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultEncryptionConfiguration.KmsKeyName"),
			},
			{
				Name: "default_partition_expiration_ms",
				Type: schema.TypeBigInt,
			},
			{
				Name: "default_table_expiration_ms",
				Type: schema.TypeBigInt,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "etag",
				Type: schema.TypeString,
			},
			{
				Name: "friendly_name",
				Type: schema.TypeString,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name: "kind",
				Type: schema.TypeString,
			},
			{
				Name: "labels",
				Type: schema.TypeJSON,
			},
			{
				Name: "last_modified_time",
				Type: schema.TypeBigInt,
			},
			{
				Name: "location",
				Type: schema.TypeString,
			},
			{
				Name:     "satisfies_pzs",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SatisfiesPZS"),
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			BigqueryDatasetAccesses(),
			BigqueryDatasetTables(),
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchBigqueryDatasets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.BigQuery.Datasets.List(c.ProjectId).Context(ctx).PageToken(nextPageToken)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		for _, d := range output.Datasets {
			call := c.Services.BigQuery.Datasets.Get(c.ProjectId, d.DatasetReference.DatasetId)
			dataset, err := call.Do()
			if err != nil {
				return err
			}
			res <- dataset
		}

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
