resource "google_compute_instance_group" "test" {
  name        = "${var.region}-instance-group"
  description = "Integration test instance group"
  zone        = "us-central1-a"
  network     = module.vpc.network_id
}