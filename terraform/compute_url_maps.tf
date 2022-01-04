################################################################################
# Compute Module - Url Maps
################################################################################

resource "google_compute_url_map" "url_map" {
  name        = "${local.prefix}-urlmap"

  default_service = google_compute_backend_service.internal-backend-service.id

  host_rule {
    hosts = [
      "cq.example.com"
    ]
    path_matcher = "root"
  }

  host_rule {
    hosts = [
      "beta.cq.example.com"
    ]
    path_matcher = "secondary"
  }

  path_matcher {
    name            = "root"
    default_service = google_compute_backend_service.internal-backend-service.id

    path_rule {
      paths = [
        "/home"
      ]
      route_action {
        weighted_backend_services {
          backend_service = google_compute_backend_service.internal-backend-service.id
          weight = 400
          header_action {
            request_headers_to_remove = ["RequestHeaderToRemove"]
            request_headers_to_add {
              header_name = "RequestHeaderToAdd"
              header_value = "RequestHeaderToAddValue"
              replace = true
            }
            response_headers_to_remove = ["ResponseHeaderToRemove"]
            response_headers_to_add {
              header_name = "ResponseHeaderToAdd"
              header_value = "ResponseHeaderToAddValue"
              replace = false
            }
          }
        }
      }
    }

    path_rule {
      paths = [
        "/login"
      ]
      service = google_compute_backend_service.internal-backend-service.id
    }

    path_rule {
      paths = [
        "/static"
      ]
      service = google_compute_backend_service.internal-backend-service.id
    }
  }

  path_matcher {
    name            = "secondary"
    default_service = google_compute_backend_service.internal-backend-service.id
  }

  test {
    service = google_compute_backend_service.internal-backend-service.id
    host    = "beta.cq.example.com"
    path    = "/healthcheck"
  }

  depends_on = [
    google_compute_backend_service.internal-backend-service
  ]
}