package resources

import (
	"context"
	"fmt"
	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	compute "google.golang.org/api/compute/v1"
)

func ComputeSslPolicies() *schema.Table {
	return &schema.Table{
		Name:         "gcp_compute_ssl_policies",
		Description:  "Represents an SSL Policy resource  Use SSL policies to control the SSL features, such as versions and cipher suites, offered by an HTTPS or SSL Proxy load balancer For more information, read  SSL Policy Concepts (== resource_for {$api_version}",
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
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveResourceId,
			},
			{
				Name:        "creation_timestamp",
				Description: "Creation timestamp in RFC3339 text format",
				Type:        schema.TypeString,
			},
			{
				Name:        "custom_features",
				Description: "A list of features enabled when the selected profile is CUSTOM The - method returns the set of features that can be specified in this list This field must be empty if the profile is not CUSTOM",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "description",
				Description: "An optional description of this resource Provide this property when you create the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "enabled_features",
				Description: "The list of features enabled in the SSL policy",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "fingerprint",
				Description: "Fingerprint of this resource A hash of the contents stored in this object This field is used in optimistic locking This field will be ignored when inserting a SslPolicy An up-to-date fingerprint must be provided in order to update the SslPolicy, otherwise the request will fail with error 412 conditionNotMet  To see the latest fingerprint, make a get() request to retrieve an SslPolicy",
				Type:        schema.TypeString,
			},
			{
				Name:        "kind",
				Description: "[Output only] Type of the resource Always compute#sslPolicyfor SSL policies",
				Type:        schema.TypeString,
			},
			{
				Name:        "min_tls_version",
				Description: "The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer This can be one of TLS_1_0, TLS_1_1, TLS_1_2",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "Name of the resource The name must be 1-63 characters long, and comply with RFC1035 Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash",
				Type:        schema.TypeString,
			},
			{
				Name:        "profile",
				Description: "Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients This can be one of COMPATIBLE, MODERN, RESTRICTED, or CUSTOM If using CUSTOM, the set of SSL features to enable must be specified in the customFeatures field",
				Type:        schema.TypeString,
			},
			{
				Name:        "self_link",
				Description: "Server-defined URL for the resource",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "gcp_compute_ssl_policy_warnings",
				Resolver: fetchComputeSslPolicyWarnings,
				Columns: []schema.Column{
					{
						Name:        "ssl_policy_id",
						Description: "Unique ID of gcp_compute_ssl_policies table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "code",
						Description: "A warning code, if applicable For example, Compute Engine returns NO_RESULTS_ON_PAGE if there are no results in the response",
						Type:        schema.TypeString,
					},
					{
						Name:        "data",
						Description: "Metadata about this warning in key: value format",
						Type:        schema.TypeJSON,
						Resolver:    resolveComputeSslPolicyWarningData,
					},
					{
						Name:        "message",
						Description: "A human-readable description of the warning code",
						Type:        schema.TypeString,
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
		call := c.Services.Compute.SslPolicies.
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
