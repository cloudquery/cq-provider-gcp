resource "google_compute_target_ssl_proxy" "ssl_proxy_test" {
  name             = "test-proxy"
  backend_service  = google_compute_backend_service.ssl_proxy_test.id
  ssl_certificates = [google_compute_ssl_certificate.ssl_proxy_test.id]
}

resource "google_compute_ssl_certificate" "ssl_proxy_test" {
  name        = "ssl-proxy-test-default-cert"

  private_key = file("certs/example.key")
  certificate = file("certs/example.crt")
}

resource "google_compute_backend_service" "ssl_proxy_test" {
  name          = "ssl-proxy-test-backend-service"
  protocol      = "SSL"
  health_checks = [google_compute_health_check.ssl_proxy_test.id]
}

resource "google_compute_health_check" "ssl_proxy_test" {
  name               = "ssl-proxy-test-health-check"
  check_interval_sec = 1
  timeout_sec        = 1
  tcp_health_check {
    port = "443"
  }
}