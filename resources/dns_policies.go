package resources

import (
	"context"
	"fmt"
	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	dns "google.golang.org/api/dns/v1"
)

func DNSPolicies() *schema.Table {
	return &schema.Table{
		Name:         "gcp_dns_policies",
		Resolver:     fetchDnsPolicies,
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
				Name:     "alternative_name_server_config_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AlternativeNameServerConfig.Kind"),
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "enable_inbound_forwarding",
				Type: schema.TypeBool,
			},
			{
				Name: "enable_logging",
				Type: schema.TypeBool,
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
		},
		Relations: []*schema.Table{
			{
				Name:     "gcp_dns_policy_alternative_name_server_target_name_servers",
				Resolver: fetchDnsPolicyAlternativeNameServerTargetNameServers,
				Columns: []schema.Column{
					{
						Name:     "policy_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "forwarding_path",
						Type: schema.TypeString,
					},
					{
						Name: "ipv4_address",
						Type: schema.TypeString,
					},
					{
						Name: "kind",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "gcp_dns_policy_networks",
				Resolver: fetchDnsPolicyNetworks,
				Columns: []schema.Column{
					{
						Name:     "policy_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "kind",
						Type: schema.TypeString,
					},
					{
						Name: "network_url",
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
func fetchDnsPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Dns.Policies.
			List(c.ProjectId).
			Context(ctx).
			PageToken(nextPageToken)

		output, err := call.Do()
		if err != nil {
			return err
		}

		res <- output.Policies

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func fetchDnsPolicyAlternativeNameServerTargetNameServers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(*dns.Policy)
	if !ok {
		return fmt.Errorf("expected *dns.Policy but got %T", p)
	}

	if p.AlternativeNameServerConfig == nil {
		return nil
	}

	res <- p.AlternativeNameServerConfig.TargetNameServers
	return nil
}
func fetchDnsPolicyNetworks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(*dns.Policy)
	if !ok {
		return fmt.Errorf("expected *dns.Policy but got %T", p)
	}
	res <- p.Networks
	return nil
}
