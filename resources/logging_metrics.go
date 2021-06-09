package resources

import (
	"context"
	"fmt"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/logging/v2"
)

func LoggingMetrics() *schema.Table {
	return &schema.Table{
		Name:         "gcp_logging_metrics",
		Resolver:     fetchLoggingMetrics,
		Multiplex:    client.ProjectMultiplex,
		IgnoreError:  client.IgnoreErrorHandler,
		DeleteFilter: client.DeleteProjectFilter,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "exponential_buckets_options_growth_factor",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("BucketOptions.ExponentialBuckets.GrowthFactor"),
			},
			{
				Name:     "exponential_buckets_options_num_finite_buckets",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("BucketOptions.ExponentialBuckets.NumFiniteBuckets"),
			},
			{
				Name:     "exponential_buckets_options_scale",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("BucketOptions.ExponentialBuckets.Scale"),
			},
			{
				Name:     "linear_buckets_options_num_finite_buckets",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("BucketOptions.LinearBuckets.NumFiniteBuckets"),
			},
			{
				Name:     "linear_buckets_options_offset",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("BucketOptions.LinearBuckets.Offset"),
			},
			{
				Name:     "linear_buckets_options_width",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("BucketOptions.LinearBuckets.Width"),
			},
			{
				Name: "create_time",
				Type: schema.TypeString,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "filter",
				Type: schema.TypeString,
			},
			{
				Name: "label_extractors",
				Type: schema.TypeJSON,
			},
			{
				Name:     "metric_descriptor_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MetricDescriptor.Description"),
			},
			{
				Name:     "metric_descriptor_display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MetricDescriptor.DisplayName"),
			},
			{
				Name:     "metric_descriptor_launch_stage",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MetricDescriptor.LaunchStage"),
			},
			{
				Name:     "metric_descriptor_metadata_ingest_delay",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MetricDescriptor.Metadata.IngestDelay"),
			},
			{
				Name:     "metric_descriptor_metadata_launch_stage",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MetricDescriptor.Metadata.LaunchStage"),
			},
			{
				Name:     "metric_descriptor_metadata_sample_period",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MetricDescriptor.Metadata.SamplePeriod"),
			},
			{
				Name:     "metric_descriptor_metric_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MetricDescriptor.MetricKind"),
			},
			{
				Name:     "metric_descriptor_monitored_resource_types",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("MetricDescriptor.MonitoredResourceTypes"),
			},
			{
				Name:     "metric_descriptor_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MetricDescriptor.Name"),
			},
			{
				Name:     "metric_descriptor_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MetricDescriptor.Type"),
			},
			{
				Name:     "metric_descriptor_unit",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MetricDescriptor.Unit"),
			},
			{
				Name:     "metric_descriptor_value_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MetricDescriptor.ValueType"),
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "update_time",
				Type: schema.TypeString,
			},
			{
				Name: "value_extractor",
				Type: schema.TypeString,
			},
			{
				Name: "version",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "gcp_logging_metric_descriptor_labels",
				Resolver: fetchLoggingMetricDescriptorLabels,
				Columns: []schema.Column{
					{
						Name:     "metric_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "description",
						Type: schema.TypeString,
					},
					{
						Name: "key",
						Type: schema.TypeString,
					},
					{
						Name: "value_type",
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
func fetchLoggingMetrics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Logging.Projects.Metrics.
			List(fmt.Sprintf("projects/%s", c.ProjectId)).
			Context(ctx).
			PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		res <- output.Metrics
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func fetchLoggingMetricDescriptorLabels(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(*logging.LogMetric)
	if !ok {
		return fmt.Errorf("expected *logging.LogMetric but got %T", p)
	}

	if p.MetricDescriptor == nil {
		return nil
	}

	res <- p.MetricDescriptor.Labels
	return nil
}
