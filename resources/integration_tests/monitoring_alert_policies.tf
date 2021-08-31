resource "google_monitoring_alert_policy" "alert_policy" {
  display_name = "My Alert Policy"
  combiner = "OR"
  conditions {
    display_name = "test condition"
    condition_threshold {
      filter = "metric.type=\"logging.googleapis.com/user/my-test-metric-test1\" AND resource.type=\"metric\""
      duration = "60s"
      comparison = "COMPARISON_GT"
      aggregations {
        alignment_period = "60s"
        per_series_aligner = "ALIGN_RATE"
      }
    }
  }

  user_labels = {
    foo = "bar"
  }
}

resource "google_logging_metric" "logging_metric_test" {
  name   = "my-test-metric-test1"
  filter = "resource.type=gcs_bucket AND protoPayload.methodName=\"storage.setIamPermissions\""

  metric_descriptor {
    metric_kind = "DELTA"
    value_type  = "INT64"
  }
}

resource "google_logging_metric" "logging_metric_test1" {
  name   = "my-test-metric-test2"
  filter = "protoPayload.methodName= \"storage.setIamPermissions\" AND         resource.type=gcs_bucket"

  metric_descriptor {
    metric_kind = "DELTA"
    value_type  = "INT64"
  }
}
