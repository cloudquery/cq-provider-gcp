package cloudrun

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	run "google.golang.org/api/run/v1"

	"github.com/cloudquery/cq-provider-gcp/client"
)

//go:generate cq-gen --resource services --config gen.hcl --output .
func Services() *schema.Table {
	return &schema.Table{
		Name:         "gcp_cloudrun_services",
		Description:  "Service acts as a top-level container that manages a set of Routes and Configurations which implement a network service",
		Resolver:     fetchCloudrunServices,
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
				Name:        "api_version",
				Description: "The API version for this call such as \"servingknativedev/v1\"",
				Type:        schema.TypeString,
			},
			{
				Name:        "kind",
				Description: "The kind of resource, in this case \"Service\"",
				Type:        schema.TypeString,
			},
			{
				Name:        "metadata_annotations",
				Description: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Metadata.Annotations"),
			},
			{
				Name:        "metadata_cluster_name",
				Description: "Not supported by Cloud Run The name of the cluster which the object belongs to",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Metadata.ClusterName"),
			},
			{
				Name:        "metadata_creation_timestamp",
				Description: "CreationTimestamp is a timestamp representing the server time when this object was created",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Metadata.CreationTimestamp"),
			},
			{
				Name:        "metadata_deletion_grace_period_seconds",
				Description: "Not supported by Cloud Run Number of seconds allowed for this object to gracefully terminate before it will be removed from the system",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Metadata.DeletionGracePeriodSeconds"),
			},
			{
				Name:        "metadata_deletion_timestamp",
				Description: "Not supported by Cloud Run DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Metadata.DeletionTimestamp"),
			},
			{
				Name:        "metadata_finalizers",
				Description: "Not supported by Cloud Run Must be empty before the object is deleted from the registry",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Metadata.Finalizers"),
			},
			{
				Name:        "metadata_generate_name",
				Description: "Not supported by Cloud Run GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Metadata.GenerateName"),
			},
			{
				Name:        "metadata_generation",
				Description: "A sequence number representing a specific generation of the desired state",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Metadata.Generation"),
			},
			{
				Name:        "metadata_labels",
				Description: "Map of string keys and values that can be used to organize and categorize (scope and select) objects",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Metadata.Labels"),
			},
			{
				Name:        "metadata_name",
				Description: "Name must be unique within a namespace, within a Cloud Run region",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Metadata.Name"),
			},
			{
				Name:        "metadata_namespace",
				Description: "Namespace defines the space within each name must be unique, within a Cloud Run region",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Metadata.Namespace"),
			},
			{
				Name:        "metadata_resource_version",
				Description: "Optional",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Metadata.ResourceVersion"),
			},
			{
				Name:        "metadata_self_link",
				Description: "SelfLink is a URL representing this object Populated by the system",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Metadata.SelfLink"),
			},
			{
				Name:        "metadata_uid",
				Description: "UID is the unique in time and space value for this object",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Metadata.Uid"),
			},
			{
				Name:        "spec",
				Description: "Spec holds the desired state of the Service (from the client)",
				Type:        schema.TypeJSON,
			},
			{
				Name:     "status_address_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status.Address.Url"),
			},
			{
				Name:        "status_latest_created_revision_name",
				Description: "From ConfigurationStatus LatestCreatedRevisionName is the last revision that was created from this Service's Configuration",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.LatestCreatedRevisionName"),
			},
			{
				Name:        "status_latest_ready_revision_name",
				Description: "From ConfigurationStatus LatestReadyRevisionName holds the name of the latest Revision stamped out from this Service's Configuration that has had its \"Ready\" condition become \"True\"",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.LatestReadyRevisionName"),
			},
			{
				Name:        "status_observed_generation",
				Description: "ObservedGeneration is the 'Generation' of the Route that was last processed by the controller",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Status.ObservedGeneration"),
			},
			{
				Name:        "status_url",
				Description: "From RouteStatus",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.Url"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "gcp_cloudrun_service_metadata_owner_references",
				Description: "OwnerReference contains enough information to let you identify an owning object",
				Resolver:    fetchCloudrunServiceMetadataOwnerReferences,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_cloudrun_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "api_version",
						Description: "API version of the referent",
						Type:        schema.TypeString,
					},
					{
						Name:        "block_owner_deletion",
						Description: "If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed",
						Type:        schema.TypeBool,
					},
					{
						Name:        "controller",
						Description: "If true, this reference points to the managing controller",
						Type:        schema.TypeBool,
					},
					{
						Name:        "kind",
						Description: "Kind of the referent",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "Name of the referent",
						Type:        schema.TypeString,
					},
					{
						Name:        "uid",
						Description: "UID of the referent",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "gcp_cloudrun_service_status_conditions",
				Description: "Condition defines a generic condition for a Resource",
				Resolver:    fetchCloudrunServiceStatusConditions,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_cloudrun_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "last_transition_time",
						Description: "Optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "message",
						Description: "Optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "reason",
						Description: "Optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "severity",
						Description: "Optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "status",
						Description: "Status of the condition, one of True, False, Unknown",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "type is used to communicate the status of the reconciliation process",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "gcp_cloudrun_service_status_traffic",
				Description: "TrafficTarget holds a single entry of the routing table for a Route",
				Resolver:    fetchCloudrunServiceStatusTraffics,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_cloudrun_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "configuration_name",
						Description: "ConfigurationName of a configuration to whose latest revision we will send this portion of traffic",
						Type:        schema.TypeString,
					},
					{
						Name:        "latest_revision",
						Description: "Optional",
						Type:        schema.TypeBool,
					},
					{
						Name:        "percent",
						Description: "Percent specifies percent of the traffic to this Revision or Configuration",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "revision_name",
						Description: "RevisionName of a specific revision to which to send this portion of traffic",
						Type:        schema.TypeString,
					},
					{
						Name:        "tag",
						Description: "Optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "url",
						Description: "Output only",
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

func fetchCloudrunServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	var nextPageToken string
	for {
		if c.Services == nil || c.Services.CloudRun == nil || c.Services.CloudRun.Projects == nil || c.Services.CloudRun.Projects.Locations == nil || c.Services.CloudRun.Projects.Locations.Services == nil {
			return nil
		}
		call := c.Services.CloudRun.Projects.Locations.Services.List("projects/" + c.ProjectId + "/locations/-").Continue(nextPageToken)
		list, err := c.RetryingDo(ctx, call)
		if err != nil {
			return diag.WrapError(err)
		}
		output := list.(*run.ListServicesResponse)

		if output.Items == nil {
			return nil
		}

		res <- output.Items
		if output.Metadata == nil || output.Metadata.Continue == "" {
			break
		}
		nextPageToken = output.Metadata.Continue
	}
	return nil
}
func fetchCloudrunServiceMetadataOwnerReferences(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*run.Service)
	if r.Metadata == nil {
		return nil
	}
	res <- r.Metadata.OwnerReferences
	return nil
}
func fetchCloudRunServiceSpecTemplateMetadataOwnerReferences(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*run.Service)
	if r.Spec == nil || r.Spec.Template == nil || r.Spec.Template.Metadata == nil {
		return nil
	}
	res <- r.Spec.Template.Metadata.OwnerReferences
	return nil
}
func fetchCloudRunServiceSpecTemplateSpecContainers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*run.Service)
	if r.Spec == nil || r.Spec.Template == nil || r.Spec.Template.Spec == nil {
		return nil
	}
	res <- r.Spec.Template.Spec.Containers
	return nil
}
func fetchCloudRunServiceSpecTemplateSpecContainerEnvs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*run.Container)
	res <- r.Env
	return nil
}
func fetchCloudRunServiceSpecTemplateSpecContainerEnvFroms(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*run.Container)
	res <- r.EnvFrom
	return nil
}
func fetchCloudRunServiceSpecTemplateSpecContainerLivenessProbeHttpGetHttpHeaders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*run.Container)
	if r.LivenessProbe == nil || r.LivenessProbe.HttpGet == nil {
		return nil
	}
	res <- r.LivenessProbe.HttpGet.HttpHeaders
	return nil
}
func fetchCloudRunServiceSpecTemplateSpecContainerStartupProbeHttpGetHttpHeaders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*run.Container)
	if r.StartupProbe == nil || r.StartupProbe.HttpGet == nil {
		return nil
	}
	res <- r.StartupProbe.HttpGet.HttpHeaders
	return nil
}
func fetchCloudRunServiceSpecTemplateSpecContainerVolumeMounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*run.Container)
	res <- r.VolumeMounts
	return nil
}
func fetchCloudRunServiceSpecTemplateSpecVolumes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*run.Service)
	if r.Spec == nil || r.Spec.Template == nil || r.Spec.Template.Spec == nil {
		return nil
	}
	res <- r.Spec.Template.Spec.Volumes
	return nil
}
func fetchCloudRunServiceSpecTemplateSpecVolumeConfigMapItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*run.Volume)
	if r.ConfigMap == nil {
		return nil
	}
	res <- r.ConfigMap.Items
	return nil
}
func fetchCloudRunServiceSpecTemplateSpecVolumeSecretItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*run.Volume)
	if r.Secret == nil {
		return nil
	}
	res <- r.Secret.Items
	return nil
}
func fetchCloudRunServiceSpecTraffics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*run.Service)
	if r.Spec == nil {
		return nil
	}
	res <- r.Spec.Traffic
	return nil
}
func fetchCloudrunServiceStatusConditions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*run.Service)
	if r.Status == nil {
		return nil
	}
	res <- r.Status.Conditions
	return nil
}
func fetchCloudrunServiceStatusTraffics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*run.Service)
	if r.Status == nil {
		return nil
	}
	res <- r.Status.Traffic
	return nil
}
