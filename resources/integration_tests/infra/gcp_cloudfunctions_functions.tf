resource "google_storage_bucket" "bucket_func" {
  name = "bucket-func-${var.test_prefix}${var.test_suffix}"
}

resource "google_storage_bucket_object" "bucket_object_function" {
  name   = "helloworld-${var.test_prefix}${var.test_suffix}"
  bucket = google_storage_bucket.bucket_func.name
  source = "./helloworld.zip"
}

resource "google_cloudfunctions_function" "helloworld_function" {
  name        = "helloworld-${var.test_prefix}${var.test_suffix}"
  description = "My function ${var.test_prefix}${var.test_suffix}"
  runtime     = "go113"

  available_memory_mb   = 128
  source_archive_bucket = google_storage_bucket.bucket_func.name
  source_archive_object = google_storage_bucket_object.bucket_object_function.name
  trigger_http          = true
  entry_point           = "HelloHTTP"
}
