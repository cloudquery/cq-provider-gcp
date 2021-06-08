package resources

import (
	"context"
	"fmt"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/bigquery/v2"
)

func BigqueryDatasetAccesses() *schema.Table {
	return &schema.Table{
		Name:     "gcp_bigquery_dataset_accesses",
		Resolver: fetchBigqueryDatasetAccesses,
		Columns: []schema.Column{
			{
				Name:     "dataset_id",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIdResolver,
			},
			{
				Name:     "target_types",
				Type:     schema.TypeStringArray,
				Resolver: resolveBigqueryDatasetAccessTargetTypes,
			},
			{
				Name: "domain",
				Type: schema.TypeString,
			},
			{
				Name: "group_by_email",
				Type: schema.TypeString,
			},
			{
				Name: "iam_member",
				Type: schema.TypeString,
			},
			{
				Name: "role",
				Type: schema.TypeString,
			},
			{
				Name:     "routine_dataset_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Routine.DatasetId"),
			},
			{
				Name:     "routine_project_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Routine.ProjectId"),
			},
			{
				Name:     "routine_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Routine.RoutineId"),
			},
			{
				Name: "special_group",
				Type: schema.TypeString,
			},
			{
				Name: "user_by_email",
				Type: schema.TypeString,
			},
			{
				Name:     "view_dataset_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("View.DatasetId"),
			},
			{
				Name:     "view_project_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("View.ProjectId"),
			},
			{
				Name:     "view_table_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("View.TableId"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchBigqueryDatasetAccesses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(*bigquery.Dataset)
	if !ok {
		return fmt.Errorf("expected bigquery.Dataset but got %T", p)
	}
	res <- p.Access
	return nil
}

func resolveBigqueryDatasetAccessTargetTypes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(*bigquery.DatasetAccess)
	if !ok {
		return fmt.Errorf("expected bigquery.DatasetAccess but got %T", p)
	}
	if p.Dataset == nil {
		return nil
	}
	result := make([]string, 0, len(p.Dataset.TargetTypes))
	for _, t := range p.Dataset.TargetTypes {
		result = append(result, t.TargetType)
	}
	return resource.Set(c.Name, result)
}
