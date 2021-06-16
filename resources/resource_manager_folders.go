package resources

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	cloudresourcemanager "google.golang.org/api/cloudresourcemanager/v3"
)

func ResourceManagerFolders() *schema.Table {
	return &schema.Table{
		Name:         "gcp_resource_manager_folders",
		Description:  "Folder: A folder in an organization's resource hierarchy, used to organize that organization's resources. ",
		Resolver:     fetchResourceManagerFolders,
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
				Name:     "policy",
				Type:     schema.TypeJSON,
				Resolver: resolveResourceManagerFolderPolicy,
			},
			{
				Name:        "create_time",
				Description: "CreateTime: Output only",
				Type:        schema.TypeTimestamp,
				Resolver:    resolveResourceManagerFolderCreateTime,
			},
			{
				Name:        "delete_time",
				Description: "DeleteTime: Output only",
				Type:        schema.TypeTimestamp,
				Resolver:    resolveResourceManagerFolderDeleteTime,
			},
			{
				Name:        "display_name",
				Description: "DisplayName: The folder's display name",
				Type:        schema.TypeString,
			},
			{
				Name:        "etag",
				Description: "Etag: Output only",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "Name: Output only",
				Type:        schema.TypeString,
			},
			{
				Name:        "parent",
				Description: "Parent: Required",
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
				Resolver:    resolveResourceManagerFolderUpdateTime,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchResourceManagerFolders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	//todo service account needs specific permissions to list folders https://cloud.google.com/resource-manager/docs/creating-managing-folders#folder-permissions
	call := c.Services.ResourceManager.Folders.
		List().
		Context(ctx)
	output, err := call.Do()
	if err != nil {
		return err
	}
	res <- output.Folders
	return nil
}
func resolveResourceManagerFolderPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	client := meta.(*client.Client)
	p, ok := resource.Item.(*cloudresourcemanager.Folder)
	if !ok {
		return fmt.Errorf("expected *cloudresourcemanager.Folder but got %T", p)
	}

	call := client.Services.ResourceManager.Projects.
		GetIamPolicy("folders/"+p.Name, &cloudresourcemanager.GetIamPolicyRequest{}).
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
func resolveResourceManagerFolderCreateTime(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(*cloudresourcemanager.Folder)
	if !ok {
		return fmt.Errorf("expected *cloudresourcemanager.Folder but got %T", p)
	}

	date, err := client.ParseISODate(p.CreateTime)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, date)
}
func resolveResourceManagerFolderDeleteTime(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(*cloudresourcemanager.Folder)
	if !ok {
		return fmt.Errorf("expected *cloudresourcemanager.Folder but got %T", p)
	}

	date, err := client.ParseISODate(p.DeleteTime)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, date)
}
func resolveResourceManagerFolderUpdateTime(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(*cloudresourcemanager.Folder)
	if !ok {
		return fmt.Errorf("expected *cloudresourcemanager.Folder but got %T", p)
	}

	date, err := client.ParseISODate(p.UpdateTime)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, date)
}
