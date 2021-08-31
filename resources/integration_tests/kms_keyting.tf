resource "google_kms_key_ring" "example-keyring" {
  name     = "keyring-example"
  location = "global"
}

resource "google_kms_crypto_key" "example-key" {
  name            = "crypto-key-example"
  key_ring        = google_kms_key_ring.example-keyring.id
  rotation_period = "100000s"

  lifecycle {
    prevent_destroy = true
  }
}

data "google_iam_policy" "admin" {
  binding {
    role = "roles/cloudkms.cryptoKeyEncrypter"

    members = [
      "user:andriir@cloudquery.io",
    ]
  }
}

resource "google_kms_crypto_key_iam_policy" "crypto_key" {
  crypto_key_id = google_kms_crypto_key.example-key.id
  policy_data = data.google_iam_policy.admin.policy_data
}