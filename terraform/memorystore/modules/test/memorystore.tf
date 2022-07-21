module "memorystore" {
  source  = "terraform-google-modules/memorystore/google"
  version = "~> 4.4.0"

  name    = "${var.prefix}-memorystore"
  project = var.project_id

  region       = var.region
  tier         = "BASIC"
  auth_enabled = true

  labels = var.labels
}