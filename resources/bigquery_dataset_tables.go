package resources

import (
	"context"
	"fmt"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/bigquery/v2"
)

func BigqueryDatasetTables() *schema.Table {
	return &schema.Table{
		Name:     "gcp_bigquery_dataset_tables",
		Resolver: fetchBigqueryDatasetTables,
		Columns: []schema.Column{
			{
				Name:     "dataset_id",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIdResolver,
			},
			{
				Name:     "clustering_fields",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Clustering.Fields"),
			},
			{
				Name: "creation_time",
				Type: schema.TypeBigInt,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name:     "encryption_configuration_kms_key_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EncryptionConfiguration.KmsKeyName"),
			},
			{
				Name: "etag",
				Type: schema.TypeString,
			},
			{
				Name: "expiration_time",
				Type: schema.TypeBigInt,
			},
			{
				Name:     "external_data_configuration_autodetect",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ExternalDataConfiguration.Autodetect"),
			},
			{
				Name:     "external_data_configuration_compression",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ExternalDataConfiguration.Compression"),
			},
			{
				Name:     "external_data_configuration_connection_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ExternalDataConfiguration.ConnectionId"),
			},
			{
				Name:     "external_data_configuration_ignore_unknown_values",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ExternalDataConfiguration.IgnoreUnknownValues"),
			},
			{
				Name:     "external_data_configuration_max_bad_records",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ExternalDataConfiguration.MaxBadRecords"),
			},
			{
				Name:     "external_data_configuration_schema",
				Type:     schema.TypeJSON,
				Resolver: resolveBigqueryDatasetTableExternalDataConfigurationSchema,
			},
			{
				Name:     "external_data_configuration_source_format",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ExternalDataConfiguration.SourceFormat"),
			},
			{
				Name:     "external_data_configuration_source_uris",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ExternalDataConfiguration.SourceUris"),
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
				Name:     "materialized_view_enable_refresh",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("MaterializedView.EnableRefresh"),
			},
			{
				Name:     "materialized_view_last_refresh_time",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("MaterializedView.LastRefreshTime"),
			},
			{
				Name:     "materialized_view_query",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MaterializedView.Query"),
			},
			{
				Name:     "materialized_view_refresh_interval_ms",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("MaterializedView.RefreshIntervalMs"),
			},
			{
				Name:     "model_options_labels",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Model.ModelOptions.Labels"),
			},
			{
				Name:     "model_options_loss_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Model.ModelOptions.LossType"),
			},
			{
				Name:     "model_options_model_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Model.ModelOptions.ModelType"),
			},
			{
				Name: "num_bytes",
				Type: schema.TypeBigInt,
			},
			{
				Name: "num_long_term_bytes",
				Type: schema.TypeBigInt,
			},
			{
				Name: "num_physical_bytes",
				Type: schema.TypeBigInt,
			},
			{
				Name: "num_rows",
				Type: schema.TypeBigInt,
			},
			{
				Name:     "range_partitioning_field",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RangePartitioning.Field"),
			},
			{
				Name:     "range_partitioning_range_end",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("RangePartitioning.Range.End"),
			},
			{
				Name:     "range_partitioning_range_interval",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("RangePartitioning.Range.Interval"),
			},
			{
				Name:     "range_partitioning_range_start",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("RangePartitioning.Range.Start"),
			},
			{
				Name: "require_partition_filter",
				Type: schema.TypeBool,
			},
			{
				Name:     "schema",
				Type:     schema.TypeJSON,
				Resolver: resolveBigqueryDatasetTableSchema,
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
			},
			{
				Name:     "streaming_buffer_estimated_bytes",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("StreamingBuffer.EstimatedBytes"),
			},
			{
				Name:     "streaming_buffer_estimated_rows",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("StreamingBuffer.EstimatedRows"),
			},
			{
				Name:     "streaming_buffer_oldest_entry_time",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("StreamingBuffer.OldestEntryTime"),
			},
			{
				Name:     "time_partitioning_expiration_ms",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("TimePartitioning.ExpirationMs"),
			},
			{
				Name:     "time_partitioning_field",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TimePartitioning.Field"),
			},
			{
				Name:     "time_partitioning_require_partition_filter",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("TimePartitioning.RequirePartitionFilter"),
			},
			{
				Name:     "time_partitioning_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TimePartitioning.Type"),
			},
			{
				Name: "type",
				Type: schema.TypeString,
			},
			{
				Name:     "view_query",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("View.Query"),
			},
			{
				Name:     "view_use_legacy_sql",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("View.UseLegacySql"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "gcp_bigquery_dataset_table_dataset_model_training_runs",
				Resolver: fetchBigqueryDatasetTableDatasetModelTrainingRuns,
				Columns: []schema.Column{
					{
						Name:     "dataset_table_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "start_time",
						Type: schema.TypeString,
					},
					{
						Name: "state",
						Type: schema.TypeString,
					},
					{
						Name:     "training_options_early_stop",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("TrainingOptions.EarlyStop"),
					},
					{
						Name:     "training_options_l1_reg",
						Type:     schema.TypeFloat,
						Resolver: schema.PathResolver("TrainingOptions.L1Reg"),
					},
					{
						Name:     "training_options_l2_reg",
						Type:     schema.TypeFloat,
						Resolver: schema.PathResolver("TrainingOptions.L2Reg"),
					},
					{
						Name:     "training_options_learn_rate",
						Type:     schema.TypeFloat,
						Resolver: schema.PathResolver("TrainingOptions.LearnRate"),
					},
					{
						Name:     "training_options_learn_rate_strategy",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("TrainingOptions.LearnRateStrategy"),
					},
					{
						Name:     "training_options_line_search_init_learn_rate",
						Type:     schema.TypeFloat,
						Resolver: schema.PathResolver("TrainingOptions.LineSearchInitLearnRate"),
					},
					{
						Name:     "training_options_max_iteration",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("TrainingOptions.MaxIteration"),
					},
					{
						Name:     "training_options_min_rel_progress",
						Type:     schema.TypeFloat,
						Resolver: schema.PathResolver("TrainingOptions.MinRelProgress"),
					},
					{
						Name:     "training_options_warm_start",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("TrainingOptions.WarmStart"),
					},
				},
			},
			{
				Name:     "gcp_bigquery_dataset_table_view_user_defined_function_resources",
				Resolver: fetchBigqueryDatasetTableViewUserDefinedFunctionResources,
				Columns: []schema.Column{
					{
						Name:     "dataset_table_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "inline_code",
						Type: schema.TypeString,
					},
					{
						Name: "resource_uri",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchBigqueryDatasetTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(*bigquery.Dataset)
	if !ok {
		return fmt.Errorf("expected *bigquery.Dataset but got %T", p)
	}
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.BigQuery.Tables.List(c.ProjectId, p.DatasetReference.DatasetId).Context(ctx).PageToken(nextPageToken)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		for _, t := range output.Tables {
			call := c.Services.BigQuery.Tables.Get(c.ProjectId, p.DatasetReference.DatasetId, t.TableReference.TableId)
			table, err := call.Do()
			if err != nil {
				return err
			}
			res <- table
		}

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}

func resolveBigqueryDatasetTableExternalDataConfigurationSchema(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(*bigquery.Table)
	if !ok {
		return fmt.Errorf("expected *bigquery.Table but got %T", p)
	}

	if p.ExternalDataConfiguration == nil || p.ExternalDataConfiguration.Schema == nil {
		return nil
	}

	schema := make(map[string]interface{})
	for _, f := range p.ExternalDataConfiguration.Schema.Fields {
		schema[f.Name] = f.Type
	}
	return resource.Set(c.Name, schema)
}

func resolveBigqueryDatasetTableSchema(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(*bigquery.Table)
	if !ok {
		return fmt.Errorf("expected *bigquery.Table but got %T", p)
	}

	if p.Schema == nil {
		return nil
	}

	schema := make(map[string]interface{})
	for _, f := range p.Schema.Fields {
		schema[f.Name] = f.Type
	}
	return resource.Set(c.Name, schema)
}

func fetchBigqueryDatasetTableDatasetModelTrainingRuns(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(*bigquery.Table)
	if !ok {
		return fmt.Errorf("expected *bigquery.Table but got %T", p)
	}

	if p.Model == nil {
		return nil
	}

	res <- p.Model.TrainingRuns
	return nil
}

func fetchBigqueryDatasetTableViewUserDefinedFunctionResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(*bigquery.Table)
	if !ok {
		return fmt.Errorf("expected *bigquery.Table but got %T", p)
	}

	if p.View == nil {
		return nil
	}

	res <- p.View.UserDefinedFunctionResources
	return nil
}
