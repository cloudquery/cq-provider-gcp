resource "google_compute_url_map" "urlmaps_urlmap" {
  name = "urlmap${var.test_prefix}${var.test_suffix}"
  description = "a description"

  region = var.region
  default_service = google_compute_backend_bucket.urlmaps_compute_backend_bucket.id

  host_rule {
    hosts = [
      "mysite.com"]
    path_matcher = "mysite"
  }

  host_rule {
    hosts = [
      "myothersite.com"]
    path_matcher = "otherpaths"
  }

  path_matcher {
    name = "mysite"
    default_service = google_compute_backend_bucket.urlmaps_compute_backend_bucket.id

    path_rule {
      paths = [
        "/home"]
      service = google_compute_backend_bucket.urlmaps_compute_backend_bucket.id
    }

    path_rule {
      paths = [
        "/login"]
      service = google_compute_region_backend_service.gcp_forwarding_rules_backend_svc.id
    }

    path_rule {
      paths = [
        "/static"]
      service = google_compute_backend_bucket.urlmaps_compute_backend_bucket.id
    }
  }

  path_matcher {
    name = "otherpaths"
    default_service = google_compute_backend_bucket.urlmaps_compute_backend_bucket.id
  }

  test {
    service = google_compute_backend_bucket.urlmaps_compute_backend_bucket.id
    host = "hi.com"
    path = "/home"
  }

  depends_on = [google_compute_backend_bucket.urlmaps_compute_backend_bucket]
}


resource "google_compute_backend_bucket" "urlmaps_compute_backend_bucket" {
  name = "static-asset-backend-${var.test_prefix}${var.test_suffix}"
  bucket_name = google_storage_bucket.urlmaps_storage_bucket.name
  enable_cdn = true
}

resource "google_storage_bucket" "urlmaps_storage_bucket" {
  name = "static-asset-bucket-${var.test_prefix}${var.test_suffix}"
  location = "US"
}