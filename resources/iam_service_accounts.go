package resources

import (
	"context"
	"fmt"
	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/iam/v1"
)

func IamServiceAccounts() *schema.Table {
	return &schema.Table{
		Name:         "gcp_iam_service_accounts",
		Resolver:     fetchIamServiceAccounts,
		Multiplex:    client.ProjectMultiplex,
		IgnoreError:  client.IgnoreErrorHandler,
		DeleteFilter: client.DeleteProjectFilter,
		Columns: []schema.Column{
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "disabled",
				Type: schema.TypeBool,
			},
			{
				Name: "display_name",
				Type: schema.TypeString,
			},
			{
				Name: "email",
				Type: schema.TypeString,
			},
			{
				Name: "etag",
				Type: schema.TypeString,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "oauth2_client_id",
				Type: schema.TypeString,
			},
			{
				Name: "project_id",
				Type: schema.TypeString,
			},
			{
				Name: "unique_id",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "gcp_iam_service_account_keys",
				Resolver: fetchIamServiceAccountKeys,
				Columns: []schema.Column{
					{
						Name:     "service_account_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "key_algorithm",
						Type: schema.TypeString,
					},
					{
						Name: "key_origin",
						Type: schema.TypeString,
					},
					{
						Name: "key_type",
						Type: schema.TypeString,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name:     "valid_after_time",
						Type:     schema.TypeTimestamp,
						Resolver: client.ISODateResolver("ValidAfterTime"),
					},
					{
						Name:     "valid_before_time",
						Type:     schema.TypeTimestamp,
						Resolver: client.ISODateResolver("ValidBeforeTime"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIamServiceAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Iam.Projects.ServiceAccounts.List("projects/" + c.ProjectId).Context(ctx)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}
		res <- output.Accounts
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func fetchIamServiceAccountKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	p, ok := parent.Item.(*iam.ServiceAccount)
	if !ok {
		return fmt.Errorf("expected *iam.ServiceAccount but got %T", p)
	}
	call := c.Services.Iam.Projects.ServiceAccounts.Keys.List(p.Name).Context(ctx)

	output, err := call.Do()
	if err != nil {
		return err
	}
	res <- output.Keys
	return nil
}
