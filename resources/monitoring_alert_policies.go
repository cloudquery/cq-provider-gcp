package resources

import (
	"context"
	"fmt"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/monitoring/v3"
)

func MonitoringAlertPolicies() *schema.Table {
	return &schema.Table{
		Name:         "gcp_monitoring_alert_policies",
		Resolver:     fetchMonitoringAlertPolicies,
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
				Name: "combiner",
				Type: schema.TypeString,
			},
			{
				Name:     "creation_record_mutate_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreationRecord.MutateTime"),
			},
			{
				Name:     "creation_record_mutated_by",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreationRecord.MutatedBy"),
			},
			{
				Name: "display_name",
				Type: schema.TypeString,
			},
			{
				Name:     "documentation_content",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Documentation.Content"),
			},
			{
				Name:     "documentation_mime_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Documentation.MimeType"),
			},
			{
				Name: "enabled",
				Type: schema.TypeBool,
			},
			{
				Name:     "mutate_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MutationRecord.MutateTime"),
			},
			{
				Name:     "mutated_by",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MutationRecord.MutatedBy"),
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "notification_channels",
				Type: schema.TypeStringArray,
			},
			{
				Name: "user_labels",
				Type: schema.TypeJSON,
			},
			{
				Name:     "validity_code",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Validity.Code"),
			},
			{
				Name:     "validity_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Validity.Message"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "gcp_monitoring_alert_policy_conditions",
				Resolver: fetchMonitoringAlertPolicyConditions,
				Columns: []schema.Column{
					{
						Name:     "alert_policy_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "absent_duration",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ConditionAbsent.Duration"),
					},
					{
						Name:     "absent_filter",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ConditionAbsent.Filter"),
					},
					{
						Name:     "absent_trigger_count",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("ConditionAbsent.Trigger.Count"),
					},
					{
						Name:     "absent_trigger_percent",
						Type:     schema.TypeFloat,
						Resolver: schema.PathResolver("ConditionAbsent.Trigger.Percent"),
					},
					{
						Name:     "monitoring_query_language_duration",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ConditionMonitoringQueryLanguage.Duration"),
					},
					{
						Name:     "monitoring_query_language_query",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ConditionMonitoringQueryLanguage.Query"),
					},
					{
						Name:     "monitoring_query_language_trigger_count",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("ConditionMonitoringQueryLanguage.Trigger.Count"),
					},
					{
						Name:     "monitoring_query_language_trigger_percent",
						Type:     schema.TypeFloat,
						Resolver: schema.PathResolver("ConditionMonitoringQueryLanguage.Trigger.Percent"),
					},
					{
						Name:     "threshold_comparison",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ConditionThreshold.Comparison"),
					},
					{
						Name:     "threshold_denominator_filter",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ConditionThreshold.DenominatorFilter"),
					},
					{
						Name:     "threshold_duration",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ConditionThreshold.Duration"),
					},
					{
						Name:     "threshold_filter",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ConditionThreshold.Filter"),
					},
					{
						Name:     "threshold_value",
						Type:     schema.TypeFloat,
						Resolver: schema.PathResolver("ConditionThreshold.ThresholdValue"),
					},
					{
						Name:     "threshold_trigger_count",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("ConditionThreshold.Trigger.Count"),
					},
					{
						Name:     "threshold_trigger_percent",
						Type:     schema.TypeFloat,
						Resolver: schema.PathResolver("ConditionThreshold.Trigger.Percent"),
					},
					{
						Name: "display_name",
						Type: schema.TypeString,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:     "gcp_monitoring_alert_policy_condition_absent_aggregations",
						Resolver: fetchMonitoringAlertPolicyConditionAbsentAggregations,
						Columns: []schema.Column{
							{
								Name:     "alert_policy_condition_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "alignment_period",
								Type: schema.TypeString,
							},
							{
								Name: "cross_series_reducer",
								Type: schema.TypeString,
							},
							{
								Name: "group_by_fields",
								Type: schema.TypeStringArray,
							},
							{
								Name: "per_series_aligner",
								Type: schema.TypeString,
							},
						},
					},
					{
						Name:     "gcp_monitoring_alert_policy_condition_threshold_aggregations",
						Resolver: fetchMonitoringAlertPolicyConditionThresholdAggregations,
						Columns: []schema.Column{
							{
								Name:     "alert_policy_condition_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "alignment_period",
								Type: schema.TypeString,
							},
							{
								Name: "cross_series_reducer",
								Type: schema.TypeString,
							},
							{
								Name: "group_by_fields",
								Type: schema.TypeStringArray,
							},
							{
								Name: "per_series_aligner",
								Type: schema.TypeString,
							},
						},
					},
					{
						Name:     "gcp_monitoring_alert_policy_condition_denominator_aggregations",
						Resolver: fetchMonitoringAlertPolicyConditionDenominatorAggregations,
						Columns: []schema.Column{
							{
								Name:     "alert_policy_condition_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "alignment_period",
								Type: schema.TypeString,
							},
							{
								Name: "cross_series_reducer",
								Type: schema.TypeString,
							},
							{
								Name: "group_by_fields",
								Type: schema.TypeStringArray,
							},
							{
								Name: "per_series_aligner",
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchMonitoringAlertPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Monitoring.Projects.AlertPolicies.
			List(fmt.Sprintf("projects/%s", c.ProjectId)).
			Context(ctx).
			PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		res <- output.AlertPolicies

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func fetchMonitoringAlertPolicyConditions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(*monitoring.AlertPolicy)
	if !ok {
		return fmt.Errorf("expected *monitoring.AlertPolicy but got %T", p)
	}

	res <- p.Conditions
	return nil
}
func fetchMonitoringAlertPolicyConditionAbsentAggregations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(*monitoring.Condition)
	if !ok {
		return fmt.Errorf("expected *monitoring.Condition but got %T", p)
	}

	if p.ConditionAbsent == nil {
		return nil
	}
	res <- p.ConditionAbsent.Aggregations
	return nil
}
func fetchMonitoringAlertPolicyConditionThresholdAggregations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(*monitoring.Condition)
	if !ok {
		return fmt.Errorf("expected *monitoring.Condition but got %T", p)
	}

	if p.ConditionThreshold == nil {
		return nil
	}
	res <- p.ConditionThreshold.Aggregations
	return nil
}
func fetchMonitoringAlertPolicyConditionDenominatorAggregations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(*monitoring.Condition)
	if !ok {
		return fmt.Errorf("expected *monitoring.Condition but got %T", p)
	}

	if p.ConditionThreshold == nil {
		return nil
	}
	res <- p.ConditionThreshold.DenominatorAggregations
	return nil
}
