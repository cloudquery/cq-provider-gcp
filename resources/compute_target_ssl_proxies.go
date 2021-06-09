package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ComputeTargetSslProxies() *schema.Table {
	return &schema.Table{
		Name:         "gcp_compute_target_ssl_proxies",
		Resolver:     fetchComputeTargetSslProxies,
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
				Name: "creation_timestamp",
				Type: schema.TypeString,
			},
			{
				Name: "description",
				Type: schema.TypeString,
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
				Name: "proxy_header",
				Type: schema.TypeString,
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
			},
			{
				Name: "service",
				Type: schema.TypeString,
			},
			{
				Name: "ssl_certificates",
				Type: schema.TypeStringArray,
			},
			{
				Name: "ssl_policy",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeTargetSslProxies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Compute.TargetSslProxies.
			List(c.ProjectId).
			Context(ctx).
			PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		res <- output.Items

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
