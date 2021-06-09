package resources

import (
	"context"
	"fmt"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/compute/v1"
)

func ComputeProjects() *schema.Table {
	return &schema.Table{
		Name:         "gcp_compute_projects",
		Resolver:     fetchComputeProjects,
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
				Name:     "common_instance_metadata_fingerprint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CommonInstanceMetadata.Fingerprint"),
			},
			{
				Name:     "common_instance_metadata_items",
				Type:     schema.TypeJSON,
				Resolver: resolveComputeProjectCommonInstanceMetadataItems,
			},
			{
				Name:     "common_instance_metadata_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CommonInstanceMetadata.Kind"),
			},
			{
				Name: "creation_timestamp",
				Type: schema.TypeString,
			},
			{
				Name: "default_network_tier",
				Type: schema.TypeString,
			},
			{
				Name: "default_service_account",
				Type: schema.TypeString,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "enabled_features",
				Type: schema.TypeStringArray,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name: "kind",
				Type: schema.TypeString,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
			},
			{
				Name:     "usage_export_location_bucket_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UsageExportLocation.BucketName"),
			},
			{
				Name:     "usage_export_location_report_name_prefix",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UsageExportLocation.ReportNamePrefix"),
			},
			{
				Name: "xpn_project_status",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "gcp_compute_project_quotas",
				Resolver: fetchComputeProjectQuotas,
				Columns: []schema.Column{
					{
						Name:     "project_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "limit",
						Type: schema.TypeFloat,
					},
					{
						Name: "metric",
						Type: schema.TypeString,
					},
					{
						Name: "owner",
						Type: schema.TypeString,
					},
					{
						Name: "usage",
						Type: schema.TypeFloat,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeProjects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	call := c.Services.Compute.Projects.
		Get(c.ProjectId).
		Context(ctx)
	output, err := call.Do()
	if err != nil {
		return err
	}
	res <- output
	return nil
}
func resolveComputeProjectCommonInstanceMetadataItems(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(*compute.Project)
	if !ok {
		return fmt.Errorf("expected *compute.Project but got %T", p)
	}
	m := make(map[string]interface{})
	for _, i := range p.CommonInstanceMetadata.Items {
		m[i.Key] = i.Value
	}
	return resource.Set(c.Name, m)
}
func fetchComputeProjectQuotas(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(*compute.Project)
	if !ok {
		return fmt.Errorf("expected *compute.Project but got %T", p)
	}
	res <- p.Quotas
	return nil
}
