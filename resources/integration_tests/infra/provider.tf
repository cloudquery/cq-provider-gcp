provider "google" {
  credentials = file("credentials.json")
  project     = "cq-e2e"
  region      = "us-west1"
}