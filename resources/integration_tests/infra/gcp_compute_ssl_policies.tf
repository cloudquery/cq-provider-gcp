resource "google_compute_ssl_policy" "gcp_compute_ssl_policies_policy" {
  name = "ssl-policies-policy-${var.test_prefix}${var.test_suffix}"
  min_tls_version = "TLS_1_2"
  profile = "CUSTOM"
  custom_features = [
    "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
    "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384"]
}