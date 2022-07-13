service          = "gcp"
output_directory = "."
add_generate     = true


description_modifier "remove_read_only" {
  words = ["[Output Only] "]
}

description_modifier "remove_optional" {
  words = ["(Optional) "]
}

description_modifier "remove_field_name" {
  regex = "^.+: "
}

resource "gcp" "cloudrun" "services" {
  path = "google.golang.org/api/run/v1.Service"

  multiplex "ProjectMultiplex" {
    path = "github.com/cloudquery/cq-provider-gcp/client.ProjectMultiplex"
  }
  deleteFilter "DeleteFilter" {
    path = "github.com/cloudquery/cq-provider-gcp/client.DeleteProjectFilter"
  }
  ignoreError "IgnoreError" {
    path = "github.com/cloudquery/cq-provider-gcp/client.IgnoreErrorHandler"
  }

  userDefinedColumn "project_id" {
    type = "string"
    resolver "resolveResourceProject" {
      path = "github.com/cloudquery/cq-provider-gcp/client.ResolveProject"
    }
  }

  column "create_time" {
    type = "timestamp"
    resolver "ISODateResolver" {
      path = "github.com/cloudquery/cq-provider-gcp/client.ISODateResolver"
      path_resolver = true
    }
  }

  column "delete_time" {
    type = "timestamp"
    resolver "ISODateResolver" {
      path = "github.com/cloudquery/cq-provider-gcp/client.ISODateResolver"
      path_resolver = true
    }
  }

  column "update_time" {
    type = "timestamp"
    resolver "ISODateResolver" {
      path = "github.com/cloudquery/cq-provider-gcp/client.ISODateResolver"
      path_resolver = true
    }
  }

  relation "gcp" "cloudrun" "spec_template_spec_containers" {
    path = "google.golang.org/api/run/v1.Container"
    description = "A single application container"

    column "ports" {
      type = "json"
      generate_resolver = false
    }

    column "readiness_probe_http_get_http_headers" {
      type = "json"
      generate_resolver = false
    }
  }
}

