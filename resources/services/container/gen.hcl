service          = "gcp"
output_directory = "."
add_generate     = true


description_modifier "remove_read_only" {
  words = ["[Output Only] "]
}

description_modifier "remove_field_name" {
  regex = ".+: "
}


resource "gcp" "container" "clusters" {
  path = "google.golang.org/api/container/v1.Cluster"
  ignoreError "IgnoreError" {
    path = "github.com/cloudquery/cq-provider-gcp/client.IgnoreErrorHandler"
  }
  multiplex "ProjectMultiplex" {
    path = "github.com/cloudquery/cq-provider-gcp/client.ProjectMultiplex"
  }
  deleteFilter "ProjectDeleteFilter" {
    path = "github.com/cloudquery/cq-provider-gcp/client.DeleteProjectFilter"
  }

  userDefinedColumn "project_id" {
    type        = "string"
    description = "GCP Project Id of the resource"
    resolver "resolveResourceProject" {
      path = "github.com/cloudquery/cq-provider-gcp/client.ResolveProject"
    }
  }

  column "autoscaling_autoprovisioning_node_pool_defaults" {
    type = "json"
  }

  column "node_config" {
    type = "json"
  }

  column "node_pools_config_accelerators" {
    type = "json"
  }

  column "resource_usage_export_config" {
    type = "json"
  }

  options {
    primary_keys = [" id "]
  }
}

