
terraform {
  backend "gcs" {
    # Bucket is passed in via cli arg. Eg, terraform init -reconfigure -backend-configuration=dev.tfbackend
  }
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4.0"
    }
  }
}



provider "google" {
  region  = var.gcp_region
  project = var.gcp_project
}

data "google_compute_zones" "available_zones" {}



output "web_scrap_function_uri" {
  value = google_cloudfunctions2_function.web_scrap.service_config[0].uri
}
