service          = "gcp"
output_directory = "."
add_generate     = true


description_modifier "remove_read_only" {
  words = ["[Output Only] "]
}

description_modifier "remove_field_name" {
  regex = ".+: "
}


resource "gcp" "billing" "accounts" {
  path = "github.com/cloudquery/cq-provider-gcp/resources/services/cloudbilling.BillingAccountWrapper"
  ignoreError "IgnoreError" {
    path = "github.com/cloudquery/cq-provider-gcp/client.IgnoreErrorHandler"
  }

  multiplex "ProjectMultiplex" {
    path = "github.com/cloudquery/cq-provider-gcp/client.ProjectMultiplex"
  }

  deleteFilter "ProjectDeleteFilter" {
    path = "github.com/cloudquery/cq-provider-gcp/client.DeleteProjectFilter"
  }

  column "billing_account" {
    skip_prefix = true
  }

  column "project_billing_info_billing_account_name" {
    skip = true
  }

  column "project_billing_info_billing_enabled" {
    rename = "project_billing_enabled"
  }

  column "project_billing_info_name" {
    rename = "project_name"
  }

  column "project_billing_info_project_id" {
    rename = "project_id"
  }
}