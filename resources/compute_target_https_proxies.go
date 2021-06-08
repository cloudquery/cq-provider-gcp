package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ComputeTargetHTTPSProxies() *schema.Table {
	return &schema.Table{
		Name:         "gcp_compute_target_https_proxies",
		Resolver:     fetchComputeTargetHttpsProxies,
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
				Name: "authorization_policy",
				Type: schema.TypeString,
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
				Name: "fingerprint",
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
				Name: "proxy_bind",
				Type: schema.TypeBool,
			},
			{
				Name: "quic_override",
				Type: schema.TypeString,
			},
			{
				Name: "region",
				Type: schema.TypeString,
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
			},
			{
				Name: "server_tls_policy",
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
			{
				Name: "url_map",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeTargetHttpsProxies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Compute.TargetHttpsProxies.List(c.ProjectId).Context(ctx)
		call.PageToken(nextPageToken)
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
