resource "google_cloudfunctions_function_iam_member" "invoker" {
  project        = google_cloudfunctions_function.req_ping.project
  region         = google_cloudfunctions_function.req_ping.region
  cloud_function = google_cloudfunctions_function.req_ping.name

  role   = "roles/cloudfunctions.invoker"
  member = "allUsers"
}

resource "google_cloudfunctions_function" "req_ping" {
  project      = var.project_id
  name         = "${var.project_id}-req-ping"
  description  = "A simple HTTP function that returns a pong response."
  entry_point  = "ping"
  runtime      = "go122"
  available_memory_mb = 128
  timeout = 60

  source_archive_bucket = google_storage_bucket.source_bucket.name
  source_archive_object = google_storage_bucket_object.source_archive_req_ping.name

  trigger_http = true

}

data "archive_file" "function_got_archive" {
  type        = "zip"
  source_dir  = "${path.module}/../../../apps/backend/src/functions/req_ping"
  output_path = "${path.module}/../../../apps/backend/src/functions/req_ping.zip"
}

resource "google_storage_bucket_object" "source_archive_req_ping" {
  name   = "packages/go/function_got.${data.archive_file.function_got_archive.output_md5}.zip"
  bucket  = google_storage_bucket.source_bucket.name
  source = data.archive_file.function_got_archive.output_path
}