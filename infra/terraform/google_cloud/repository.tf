resource "google_artifact_registry_repository" "frontend" {
  description   = "next-go"
  format        = "DOCKER"
  location      = var.region
  repository_id = var.image_repo
}