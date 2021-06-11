package resources

import (
	"context"
	"fmt"
	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/logging/v2"
)

func LoggingSinks() *schema.Table {
	return &schema.Table{
		Name:         "gcp_logging_sinks",
		Resolver:     fetchLoggingSinks,
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
				Name:     "bigquery_options_use_partitioned_tables",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("BigqueryOptions.UsePartitionedTables"),
			},
			{
				Name:     "bigquery_options_uses_timestamp_column_partitioning",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("BigqueryOptions.UsesTimestampColumnPartitioning"),
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
				Name: "destination",
				Type: schema.TypeString,
			},
			{
				Name: "disabled",
				Type: schema.TypeBool,
			},
			{
				Name: "filter",
				Type: schema.TypeString,
			},
			{
				Name: "include_children",
				Type: schema.TypeBool,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "output_version_format",
				Type: schema.TypeString,
			},
			{
				Name: "update_time",
				Type: schema.TypeString,
			},
			{
				Name: "writer_identity",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "gcp_logging_sink_exclusions",
				Resolver: fetchLoggingSinkExclusions,
				Columns: []schema.Column{
					{
						Name:     "sink_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
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
						Name: "disabled",
						Type: schema.TypeBool,
					},
					{
						Name: "filter",
						Type: schema.TypeString,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "update_time",
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
func fetchLoggingSinks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Logging.Sinks.
			List(fmt.Sprintf("projects/%s", c.ProjectId)).
			Context(ctx).
			PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		res <- output.Sinks
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func fetchLoggingSinkExclusions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(*logging.LogSink)
	if !ok {
		return fmt.Errorf("expected *logging.LogSink but got %T", p)
	}

	res <- p.Exclusions
	return nil
}
