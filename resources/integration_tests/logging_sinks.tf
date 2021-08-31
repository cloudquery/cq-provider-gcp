# Our logged compute instance
resource "google_compute_instance" "my-logged-instance" {
  name = "my-instance"
  machine_type = "e2-medium"
  zone = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"

    access_config {
    }
  }
  deletion_protection = false
}

# A bucket to store logs in
resource "google_storage_bucket" "log-bucket" {
  name = "log_bucket_sink_test"
  force_destroy = true
  retention_policy {
    retention_period = 123
    is_locked = true
  }
}

# Our sink; this logs all activity related to our "my-logged-instance" instance
resource "google_logging_project_sink" "instance-sink" {
  name = "my-instance-sink"
  description = "some explaination on what this is"
  destination = "storage.googleapis.com/${google_storage_bucket.log-bucket.name}"
  filter = "resource.type = gce_instance AND resource.labels.instance_id = \"${google_compute_instance.my-logged-instance.instance_id}\""

  unique_writer_identity = true
}

# Our sink; this logs all activity related to our "my-logged-instance" instance
resource "google_logging_project_sink" "instance-sink-1" {
  name = "my-instance-sink-1"
  description = "some explaination on what this is"
  destination = "storage.googleapis.com/${google_storage_bucket.log-bucket.name}"

  unique_writer_identity = true
}