package resources

import (
	"context"
	"fmt"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/compute/v1"
)

func ComputeSslPolicies() *schema.Table {
	return &schema.Table{
		Name:         "gcp_compute_ssl_policies",
		Resolver:     fetchComputeSslPolicies,
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
				Name: "custom_features",
				Type: schema.TypeStringArray,
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
				Name: "min_tls_version",
				Type: schema.TypeString,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "profile",
				Type: schema.TypeString,
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "gcp_compute_ssl_policy_warnings",
				Resolver: fetchComputeSslPolicyWarnings,
				Columns: []schema.Column{
					{
						Name:     "ssl_policy_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "code",
						Type: schema.TypeString,
					},
					{
						Name:     "data",
						Type:     schema.TypeJSON,
						Resolver: resolveComputeSslPolicyWarningData,
					},
					{
						Name: "message",
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
func fetchComputeSslPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Compute.SslPolicies.List(c.ProjectId).Context(ctx)
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

func fetchComputeSslPolicyWarnings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(*compute.SslPolicy)
	if !ok {
		return fmt.Errorf("expected *compute.SslPolicy but got %T", p)
	}
	res <- p.Warnings
	return nil
}

func resolveComputeSslPolicyWarningData(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(*compute.SslPolicyWarnings)
	if !ok {
		return fmt.Errorf("expected *compute.SslPolicy but got %T", p)
	}
	data := make(map[string]string)
	for _, v := range p.Data {
		data[v.Key] = v.Value
	}
	return resource.Set(c.Name, data)
}
