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

# Deploy the Cloud Function
resource "google_cloudfunctions_function" "req_ping" {
  name        = "${var.project_id}-req-ping"
  description = "A simple HTTP function that returns a pong response"
  runtime     = "go122"

  available_memory_mb   = 128
  source_archive_bucket = google_storage_bucket.source_bucket.name
  source_archive_object = google_storage_bucket_object.source_archive_req_ping.name
  trigger_http         = true
  entry_point         = "Ping"
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