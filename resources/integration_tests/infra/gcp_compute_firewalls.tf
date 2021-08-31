resource "google_compute_firewall" "google_compute_firewalls_firewall" {
  name = "f-${var.test_prefix}${var.test_suffix}"
  network = google_compute_network.google_compute_firewalls_network.name


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

resource "google_compute_network" "google_compute_firewalls_network" {
  name = "n-${var.test_prefix}${var.test_suffix}"
}