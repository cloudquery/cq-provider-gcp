################################################################################
# Compute SSL Module - Policies
################################################################################

resource "google_compute_ssl_policy" "ssl-policy" {
  name    = "${local.prefix}-ssl-policy"
  profile = "MODERN"
}

################################################################################
# Compute SSL Module - Certificate
################################################################################

data "template_file" "ssl_private_key" {
  template = file("fixtures/ssl/private.key")
}


data "template_file" "ssl_public_key" {
  template = file("fixtures/ssl/public.key")
}

resource "google_compute_ssl_certificate" "gcp_compute_ssl_certificates_cert" {
  name_prefix = "${local.prefix}-ssl-cert"
  private_key = data.template_file.ssl_private_key.rendered
  certificate = data.template_file.ssl_public_key.rendered

  lifecycle {
    create_before_destroy = true
  }
}
