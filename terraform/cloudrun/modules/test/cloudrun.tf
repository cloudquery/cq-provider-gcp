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
        env {
          name = "ENV_VAR"
          value = "test"
        }
      }
    }
    metadata {
      name = "${lower(var.prefix)}-cloudrun-srv-green"
    }
  }

  metadata {
    annotations = {
      generated-by = "magic-modules"
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
  autogenerate_revision_name = false

  lifecycle {
    ignore_changes = [
      metadata.0.annotations,
    ]
  }
}