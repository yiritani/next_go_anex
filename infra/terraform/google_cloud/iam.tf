resource "google_project_service" "enable_iam_credentials_api" {
  project = var.project_id
  service = "iamcredentials.googleapis.com"

  disable_dependent_services = true
  disable_on_destroy          = true
}

resource "google_service_account" "cloud_functions_service_account" {
  account_id   = "cloud-functions-sa"
  display_name = "Cloud Functions Service Account"
  project      = var.project_id
}
