package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ComputeTargetSslProxies() *schema.Table {
	return &schema.Table{
		Name:         "gcp_compute_target_ssl_proxies",
		Description:  "Represents a Target SSL Proxy resource  A target SSL proxy is a component of a SSL Proxy load balancer Global forwarding rules reference a target SSL proxy, and the target proxy then references an external backend service For more information, read Using Target Proxies (== resource_for {$api_version}",
		Resolver:     fetchComputeTargetSslProxies,
		Multiplex:    client.ProjectMultiplex,
		IgnoreError:  client.IgnoreErrorHandler,
		DeleteFilter: client.DeleteProjectFilter,
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "resource_id",
				Description: "Original Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveResourceId,
			},
			{
				Name:        "creation_timestamp",
				Description: "Creation timestamp in RFC3339 text format",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "An optional description of this resource Provide this property when you create the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "kind",
				Description: "Type of the resource Always compute#targetSslProxy for target SSL proxies",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "Name of the resource Provided by the client when the resource is created The name must be 1-63 characters long, and comply with RFC1035 Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash",
				Type:        schema.TypeString,
			},
			{
				Name:        "proxy_header",
				Description: "Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1 The default is NONE",
				Type:        schema.TypeString,
			},
			{
				Name:        "self_link",
				Description: "Server-defined URL for the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "service",
				Description: "URL to the BackendService resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "ssl_certificates",
				Description: "URLs to SslCertificate resources that are used to authenticate connections to Backends At least one SSL certificate must be specified Currently, you may specify up to 15 SSL certificates",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "ssl_policy",
				Description: "URL of SslPolicy resource that will be associated with the TargetSslProxy resource If not set, the TargetSslProxy resource will not have any SSL policy configured",
				Type:        schema.TypeString,
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
