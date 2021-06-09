package resources

import (
	"context"
	"fmt"
	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/dns/v1"
)

func DNSManagedZones() *schema.Table {
	return &schema.Table{
		Name:         "gcp_dns_managed_zones",
		Resolver:     fetchDnsManagedZones,
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
				Name: "creation_time",
				Type: schema.TypeString,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "dns_name",
				Type: schema.TypeString,
			},
			{
				Name:     "dnssec_config_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DnssecConfig.Kind"),
			},
			{
				Name:     "dnssec_config_non_existence",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DnssecConfig.NonExistence"),
			},
			{
				Name:     "dnssec_config_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DnssecConfig.State"),
			},
			{
				Name:     "forwarding_config_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ForwardingConfig.Kind"),
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
				Name: "labels",
				Type: schema.TypeJSON,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "name_server_set",
				Type: schema.TypeString,
			},
			{
				Name: "name_servers",
				Type: schema.TypeStringArray,
			},
			{
				Name:     "peering_config_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PeeringConfig.Kind"),
			},
			{
				Name:     "peering_config_target_network_deactivate_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PeeringConfig.TargetNetwork.DeactivateTime"),
			},
			{
				Name:     "peering_config_target_network_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PeeringConfig.TargetNetwork.Kind"),
			},
			{
				Name:     "peering_config_target_network_network_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PeeringConfig.TargetNetwork.NetworkUrl"),
			},
			{
				Name:     "private_visibility_config_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PrivateVisibilityConfig.Kind"),
			},
			{
				Name:     "reverse_lookup_config_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReverseLookupConfig.Kind"),
			},
			{
				Name:     "service_directory_config_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceDirectoryConfig.Kind"),
			},
			{
				Name:     "service_directory_config_namespace_deletion_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceDirectoryConfig.Namespace.DeletionTime"),
			},
			{
				Name:     "service_directory_config_namespace_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceDirectoryConfig.Namespace.Kind"),
			},
			{
				Name:     "service_directory_config_namespace_namespace_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceDirectoryConfig.Namespace.NamespaceUrl"),
			},
			{
				Name: "visibility",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "gcp_dns_managed_zone_dnssec_config_default_key_specs",
				Resolver: fetchDnsManagedZoneDnssecConfigDefaultKeySpecs,
				Columns: []schema.Column{
					{
						Name:     "managed_zone_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "algorithm",
						Type: schema.TypeString,
					},
					{
						Name: "key_length",
						Type: schema.TypeBigInt,
					},
					{
						Name: "key_type",
						Type: schema.TypeString,
					},
					{
						Name: "kind",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "gcp_dns_managed_zone_forwarding_config_target_name_servers",
				Resolver: fetchDnsManagedZoneForwardingConfigTargetNameServers,
				Columns: []schema.Column{
					{
						Name:     "managed_zone_id",
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
				Name:     "gcp_dns_managed_zone_private_visibility_config_networks",
				Resolver: fetchDnsManagedZonePrivateVisibilityConfigNetworks,
				Columns: []schema.Column{
					{
						Name:     "managed_zone_id",
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
func fetchDnsManagedZones(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Dns.ManagedZones.List(c.ProjectId).Context(ctx).PageToken(nextPageToken)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		res <- output.ManagedZones

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func fetchDnsManagedZoneDnssecConfigDefaultKeySpecs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(*dns.ManagedZone)
	if !ok {
		return fmt.Errorf("expected *dns.ManagedZone but got %T", p)
	}

	if p.DnssecConfig == nil {
		return nil
	}

	res <- p.DnssecConfig.DefaultKeySpecs
	return nil
}
func fetchDnsManagedZoneForwardingConfigTargetNameServers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(*dns.ManagedZone)
	if !ok {
		return fmt.Errorf("expected *dns.ManagedZone but got %T", p)
	}

	if p.ForwardingConfig == nil {
		return nil
	}

	res <- p.ForwardingConfig.TargetNameServers
	return nil
}
func fetchDnsManagedZonePrivateVisibilityConfigNetworks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(*dns.ManagedZone)
	if !ok {
		return fmt.Errorf("expected *dns.ManagedZone but got %T", p)
	}

	if p.PrivateVisibilityConfig == nil {
		return nil
	}

	res <- p.PrivateVisibilityConfig.Networks
	return nil
}
