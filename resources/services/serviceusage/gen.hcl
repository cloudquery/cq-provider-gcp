service          = "gcp"
output_directory = "."
add_generate     = true


description_modifier "remove_read_only" {
  words = ["[Output Only] "]
}

description_modifier "remove_field_name" {
  regex = ".+: "
}


resource "gcp" "serviceusage" "services" {
  path = "google.golang.org/api/serviceusage/v1.GoogleApiServiceusageV1Service"
  ignoreError "IgnoreError" {
    path = "github.com/cloudquery/cq-provider-gcp/client.IgnoreErrorHandler"
  }
  multiplex "ProjectMultiplex" {
    path = "github.com/cloudquery/cq-provider-gcp/client.ProjectMultiplex"
  }
  deleteFilter "ProjectDeleteFilter" {
    path = "github.com/cloudquery/cq-provider-gcp/client.DeleteProjectFilter"
  }


  relation "gcp" "serviceusage" "config_documentation_pages" {
    column "subpages" {
      type = "json"
    }
  }
}

