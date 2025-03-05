# Archive the function source code
data "archive_file" "function_source" {
  type        = "zip"
  source_dir  = "${path.module}/../../../apps/backend/src/functions/req_ping"
  output_path = "${path.module}/../../../apps/backend/src/functions/req_ping.zip"
}

# Upload the source code to Cloud Storage
resource "google_storage_bucket_object" "source_archive_req_ping" {
  name   = "functions/req_ping/${data.archive_file.function_source.output_md5}.zip"
  bucket = google_storage_bucket.source_bucket.name
  source = data.archive_file.function_source.output_path
}

# Cloud Functions API を有効化
resource "google_project_service" "cloudfunctions" {
  project = var.project_id
  service = "cloudfunctions.googleapis.com"
  disable_dependent_services = true
  disable_on_destroy = false
}

# Cloud Build API も有効化
resource "google_project_service" "cloudbuild" {
  project = var.project_id
  service = "cloudbuild.googleapis.com"
  disable_dependent_services = true
  disable_on_destroy = false
}

# サービスアカウントに必要な権限を付与
resource "google_project_iam_member" "cloudbuild_service_account" {
  project = var.project_id
  role    = "roles/cloudbuild.builds.builder"
  member  = "serviceAccount:${var.project_id}@appspot.gserviceaccount.com"
}

resource "google_project_iam_member" "function_service_account" {
  project = var.project_id
  role    = "roles/cloudfunctions.developer"
  member  = "serviceAccount:${var.project_id}@appspot.gserviceaccount.com"
}

# Deploy the Cloud Function
resource "google_cloudfunctions_function" "req_ping" {
  depends_on = [
    google_project_service.cloudfunctions,
    google_project_service.cloudbuild,
    google_project_iam_member.cloudbuild_service_account,
    google_project_iam_member.function_service_account
  ]
  name        = "${var.project_id}-req-ping"
  description = "A simple HTTP function that returns a pong response"
  runtime     = "go122"

  available_memory_mb   = 128
  source_archive_bucket = google_storage_bucket.source_bucket.name
  source_archive_object = google_storage_bucket_object.source_archive_req_ping.name
  trigger_http         = true
  entry_point         = "Handler"
  timeout             = 60
}

# IAM entry for public access
resource "google_cloudfunctions_function_iam_member" "invoker" {
  project        = google_cloudfunctions_function.req_ping.project
  region         = google_cloudfunctions_function.req_ping.region
  cloud_function = google_cloudfunctions_function.req_ping.name

  role   = "roles/cloudfunctions.invoker"
  member = "allUsers"
}