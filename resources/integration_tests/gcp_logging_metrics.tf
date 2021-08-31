resource "google_logging_metric" "logging_metric" {
  name   = "my-test-metric"
  filter = "protoPayload.methodName=\"cloudsql.instances.update\""

  metric_descriptor {
    metric_kind = "DELTA"
    value_type  = "INT64"
  }
}