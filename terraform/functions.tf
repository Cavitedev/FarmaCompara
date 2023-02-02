
# Setup the root directory of where the source code will be stored.
locals {
  root_dir = abspath("../functions")
}


# Zip up our code so that we can store it for deployment.
data "archive_file" "web_scrap" {
  type        = "zip"
  source_dir  = "${local.root_dir}/web_scrap"
  output_path = "/tmp/web_scrap.zip"
}

resource "google_storage_bucket" "bucket" {
  name     = "${var.gcp_project}-${var.function_name}"
  location = var.gcp_region
}

resource "google_storage_bucket_object" "web_scrap_md5" {
  # Use an MD5 here. If there's no changes to the source code, this won't change either.
  # We can avoid unnecessary redeployments by validating the code is unchanged, and forcing
  # a redeployment when it has!
  name   = "${data.archive_file.web_scrap.output_md5}.zip"
  bucket = google_storage_bucket.bucket.name
  source = data.archive_file.web_scrap.output_path
}

resource "google_cloudfunctions2_function" "web_scrap" {

  name     = var.function_name
  location = var.gcp_region
  project  = var.gcp_project

  build_config {
    runtime     = "go119"
    entry_point = var.entry_point
    source {
      storage_source {
        bucket = google_storage_bucket.bucket.name
        object = "${data.archive_file.web_scrap.output_md5}.zip"
      }
    }
  }

  service_config {
    max_instance_count    = 1
    timeout_seconds       = 3600
    available_memory      = "256Mi"
    ingress_settings      = "ALLOW_ALL"
    service_account_email = google_service_account.function-sa.email
  }



}

# IAM Configuration. This allows unauthenticated, public access to the function.
# Change this if you require more control here.
resource "google_cloudfunctions2_function_iam_member" "member" {
  project        = google_cloudfunctions2_function.web_scrap.project
  location       = google_cloudfunctions2_function.web_scrap.location
  cloud_function = google_cloudfunctions2_function.web_scrap.name

  role   = "roles/cloudfunctions.invoker"
  member = "allUsers"
}

# This is the service account in which the function will act as.
resource "google_service_account" "function-sa" {
  account_id   = "function-sa"
  description  = "Controls the workflow for the cloud pipeline"
  display_name = "function-sa"
  project      = var.gcp_project

}

resource "google_project_iam_member" "functions-sa-iam" {
  project = var.gcp_project
  role    = "roles/datastore.user"
  member = "serviceAccount:${google_service_account.function-sa.email}"
}
