package resources

import (
	"context"
	"fmt"
	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	iam "google.golang.org/api/iam/v1"
	"time"
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
						Resolver: resolveIamServiceAccountKeyValidAfterTime,
					},
					{
						Name:     "valid_before_time",
						Type:     schema.TypeTimestamp,
						Resolver: resolveIamServiceAccountKeyValidBeforeTime,
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
func resolveIamServiceAccountKeyValidAfterTime(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(*iam.ServiceAccountKey)
	if !ok {
		return fmt.Errorf("expected *iam.ServiceAccountKey but got %T", p)
	}
	location, err := time.LoadLocation("UTC")
	if err != nil {
		return err
	}
	date, err := time.ParseInLocation(time.RFC3339, p.ValidAfterTime, location)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, date)
}
func resolveIamServiceAccountKeyValidBeforeTime(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(*iam.ServiceAccountKey)
	if !ok {
		return fmt.Errorf("expected *iam.ServiceAccountKey but got %T", p)
	}
	location, err := time.LoadLocation("UTC")
	if err != nil {
		return err
	}
	date, err := time.ParseInLocation(time.RFC3339, p.ValidBeforeTime, location)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, date)
}
