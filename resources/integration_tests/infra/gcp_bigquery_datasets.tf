resource "google_bigquery_dataset" "gcp_bigquery_datasets_ds" {
  dataset_id = "ds_${var.test_prefix}${var.test_suffix}"
  friendly_name = "ds_${var.test_prefix}${var.test_suffix}"
  description = "This is a test description"
  location = "EU"
  default_table_expiration_ms = 3600000

  labels = {
    env = "default"
  }


  access {
    role = "OWNER"
    user_by_email = google_service_account.bqowner.email
  }

  access {
    role = "READER"
    domain = "hashicorp.com"
  }


}

resource "google_bigquery_table" "gcp_bigquery_datasets_tb1" {
  dataset_id = time_sleep.aws_directconnect_virtual_interfaces_wait_for_id.triggers["id"]
  table_id = "test"

  time_partitioning {
    type = "DAY"
  }

  labels = {
    env = "default"
  }

  schema = <<EOF
[
  {
    "name": "permalink",
    "type": "STRING",
    "mode": "NULLABLE",
    "description": "The Permalink"
  },
  {
    "name": "state",
    "type": "STRING",
    "mode": "NULLABLE",
    "description": "State where the head office is located"
  }
]
EOF

  depends_on = [
    google_bigquery_dataset.gcp_bigquery_datasets_ds]

}


resource "time_sleep" "aws_directconnect_virtual_interfaces_wait_for_id" {
  depends_on = [
    google_bigquery_dataset.gcp_bigquery_datasets_ds]

  create_duration = "2m"

  triggers = {
    # This sets up a proper dependency on the RAM association
    id = google_bigquery_dataset.gcp_bigquery_datasets_ds.dataset_id
  }
}

resource "google_service_account" "bqowner" {
  account_id = "bqowner${var.test_suffix}"
}