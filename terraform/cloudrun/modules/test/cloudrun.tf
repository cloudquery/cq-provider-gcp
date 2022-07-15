################################################################################
# Cloud Run Module
################################################################################

resource "google_cloud_run_service" "default" {
  name     = "${lower(var.prefix)}-cloudrun-srv"
  location = "us-central1"

  template {
    spec {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}