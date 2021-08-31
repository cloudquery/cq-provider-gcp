resource "google_compute_target_https_proxy" "https_proxy_test" {
  name             = "test-proxy"
  url_map          = google_compute_url_map.https_proxy_test.id
  ssl_certificates = [google_compute_ssl_certificate.https_proxy_test.id]
  ssl_policy = google_compute_ssl_policy.custom-ssl-policy.id
}

resource "google_compute_ssl_certificate" "https_proxy_test" {
  name        = "https-proxy-test-cert"

  private_key = file("certs/example.key")
  certificate = file("certs/example.crt")
}

resource "google_compute_url_map" "https_proxy_test" {
  name        = "url-map"
  description = "a description"

  default_service  = google_compute_backend_service.https_proxy_test.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service  = google_compute_backend_service.https_proxy_test.id

    path_rule {
      paths   = ["/*"]
      service = google_compute_backend_service.https_proxy_test.id
    }
  }
}

resource "google_compute_backend_service" "https_proxy_test" {
  name        = "backend-service"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_http_health_check.https_proxy_test.id]
}

resource "google_compute_http_health_check" "https_proxy_test" {
  name               = "https-health-check"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_compute_ssl_policy" "custom-ssl-policy" {
  name            = "custom-ssl-policy"
  min_tls_version = "TLS_1_2"
  profile         = "CUSTOM"
  custom_features = ["TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384", "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384"]
}