resource "google_compute_instance_group" "test" {
  name        = "integration-test"
  description = "Integration test instance group"
  zone        = "us-central1-a"
  network     = module.vpc.network_id
}