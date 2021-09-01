resource "google_compute_firewall" "google_compute_firewalls_firewall" {
  name = "google-compute-firewalls-firewall-${var.test_suffix}"
  network = google_compute_network.network.name


  allow {
    protocol = "tcp"
    ports = [
      "80",
      "22",
      "8080",
      "1000-2000"]
  }


  source_tags = [
    "web"]
}
