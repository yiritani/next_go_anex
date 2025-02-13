resource "google_storage_bucket" "source_bucket" {
  project = var.project_id
  name    = "${var.project_id}-source-bucket"
  location = "US"
}