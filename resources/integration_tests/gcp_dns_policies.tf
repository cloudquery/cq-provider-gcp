resource "google_dns_policy" "example-policy" {
  name = "example-policy"
  enable_inbound_forwarding = true

  enable_logging = false

  alternative_name_server_config {
    target_name_servers {
      ipv4_address = "172.16.1.10"
      forwarding_path = "private"
    }
    target_name_servers {
      ipv4_address = "172.16.1.20"
    }
  }

  networks {
    network_url = google_compute_network.network-1.id
  }
  networks {
    network_url = google_compute_network.network-2.id
  }
}

resource "google_compute_network" "network-1" {
  name = "network-1"
  auto_create_subnetworks = false

}

resource "google_compute_network" "network-2" {
  name = "network-2"
  auto_create_subnetworks = false
}