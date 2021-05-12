package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	crm "google.golang.org/api/cloudresourcemanager/v3"
)

func CrmProjects() *schema.Table {
	return &schema.Table{
		Name:         "gcp_crm_projects",
		Resolver:     fetchCrmProjects,
		Multiplex:    client.ProjectMultiplex,
		DeleteFilter: client.DeleteProjectFilter,
		IgnoreError:  client.IgnoreErrorHandler,
		Columns: []schema.Column{
			{
				Name: "create_time",
				Type: schema.TypeString,
			},
			{
				Name: "delete_time",
				Type: schema.TypeString,
			},
			{
				Name: "display_name",
				Type: schema.TypeString,
			},
			{
				Name: "etag",
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
				Name: "parent",
				Type: schema.TypeString,
			},
			{
				Name: "project_id",
				Type: schema.TypeString,
			},
			{
				Name: "state",
				Type: schema.TypeString,
			},
			{
				Name: "update_time",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchCrmProjects(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	var resp *crm.ListProjectsResponse
	var err error
	c := meta.(*client.Client)
	nextPageToken := ""
	call := c.Services.Crm.Projects.List().Context(ctx)
	for {
		call.PageToken(nextPageToken)
		// Google API when quota exceeded can return both QuotaExceeded and Forbidden
		retryErr := c.RetryWithDefaultBackoffIgnoreErrors(ctx, func() (bool, error) {
			resp, err = call.Do()
			return true, err
		}, map[int]bool{client.QuotaExceeded: true, client.Forbidden: true})
		if retryErr != nil {
			return retryErr
		}
		res <- resp.Projects
		if resp.NextPageToken == "" {
			break
		}
		nextPageToken = resp.NextPageToken
	}
	return nil
}
