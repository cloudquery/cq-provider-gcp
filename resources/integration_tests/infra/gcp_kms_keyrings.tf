resource "google_kms_key_ring" "gcp_kms_keyrings_keyring" {
  name = "keyring-${var.test_prefix}${var.test_suffix}"
  location = "global"

}

resource "google_kms_crypto_key" "gcp_kms_keyrings_key" {
  name = "crypto-key-${var.test_prefix}${var.test_suffix}"
  key_ring = google_kms_key_ring.gcp_kms_keyrings_keyring.id
  rotation_period = "100000s"

  lifecycle {
    prevent_destroy = false
  }
  labels = {
    test = "test"
  }
}

data "google_iam_policy" "gcp_kms_keyrings_admin_policy" {
  binding {
    role = "roles/cloudkms.cryptoKeyEncrypter"

    members = [
      "serviceAccount:${google_service_account.gcp_kms_keyrings_service_account.email}",
    ]
  }

}

resource "google_kms_crypto_key_iam_policy" "gcp_kms_keyrings_crypto_key" {
  crypto_key_id = google_kms_crypto_key.gcp_kms_keyrings_key.id
  policy_data = data.google_iam_policy.gcp_kms_keyrings_admin_policy.policy_data
}

resource "google_service_account" "gcp_kms_keyrings_service_account" {
  account_id = "kms-account-${var.test_suffix}"
  display_name = "Service Account ${var.test_prefix}${var.test_suffix}"
}