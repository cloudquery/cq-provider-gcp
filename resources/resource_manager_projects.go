package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/cloudresourcemanager/v3"
)

func ResourceManagerProjects() *schema.Table {
	return &schema.Table{
		Name:         "gcp_resource_manager_projects",
		Description:  "Project: A project is a high-level Google Cloud entity",
		Resolver:     fetchResourceManagerProjects,
		Multiplex:    client.ProjectMultiplex,
		IgnoreError:  client.IgnoreErrorHandler,
		DeleteFilter: client.DeleteProjectFilter,
		Columns: []schema.Column{
			{
				Name:     "policy",
				Type:     schema.TypeJSON,
				Resolver: resolveResourceManagerProjectPolicy,
			},
			{
				Name:        "create_time",
				Description: "CreateTime: Output only",
				Type:        schema.TypeTimestamp,
				Resolver:    resolveResourceManagerProjectCreateTime,
			},
			{
				Name:        "delete_time",
				Description: "DeleteTime: Output only",
				Type:        schema.TypeTimestamp,
				Resolver:    resolveResourceManagerProjectDeleteTime,
			},
			{
				Name:        "display_name",
				Description: "DisplayName: Optional",
				Type:        schema.TypeString,
			},
			{
				Name:        "etag",
				Description: "Etag: Output only",
				Type:        schema.TypeString,
			},
			{
				Name:        "labels",
				Description: "Labels: Optional",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "name",
				Description: "Name: Output only",
				Type:        schema.TypeString,
			},
			{
				Name:        "parent",
				Description: "Parent: Optional",
				Type:        schema.TypeString,
			},
			{
				Name:        "project_id",
				Description: "ProjectId: Immutable",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "State: Output only",
				Type:        schema.TypeString,
			},
			{
				Name:        "update_time",
				Description: "UpdateTime: Output only",
				Type:        schema.TypeTimestamp,
				Resolver:    resolveResourceManagerProjectUpdateTime,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchResourceManagerProjects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	call := c.Services.ResourceManager.Projects.
		Get("projects/" + c.ProjectId).
		Context(ctx)
	output, err := call.Do()
	if err != nil {
		return err
	}
	res <- output
	return nil
}
func resolveResourceManagerProjectPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	client := meta.(*client.Client)
	p, ok := resource.Item.(*cloudresourcemanager.Project)
	if !ok {
		return fmt.Errorf("expected *cloudresourcemanager.Project but got %T", p)
	}

	call := client.Services.ResourceManager.Projects.
		GetIamPolicy("projects/"+p.ProjectId, &cloudresourcemanager.GetIamPolicyRequest{}).
		Context(ctx)
	output, err := call.Do()
	if err != nil {
		return err
	}
	var policy map[string]interface{}
	data, err := json.Marshal(output)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &policy); err != nil {
		return err
	}

	return resource.Set(c.Name, policy)
}
func resolveResourceManagerProjectCreateTime(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(*cloudresourcemanager.Project)
	if !ok {
		return fmt.Errorf("expected *cloudresourcemanager.Project but got %T", p)
	}

	date, err := client.ParseISODate(p.CreateTime)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, date)
}
func resolveResourceManagerProjectDeleteTime(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(*cloudresourcemanager.Project)
	if !ok {
		return fmt.Errorf("expected *cloudresourcemanager.Project but got %T", p)
	}

	date, err := client.ParseISODate(p.DeleteTime)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, date)
}
func resolveResourceManagerProjectUpdateTime(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(*cloudresourcemanager.Project)
	if !ok {
		return fmt.Errorf("expected *cloudresourcemanager.Project but got %T", p)
	}

	date, err := client.ParseISODate(p.UpdateTime)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, date)
}
