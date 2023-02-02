variable "gcp_region" {
  type        = string
  description = "Region to use for GCP provider"
  default     = "europe-west2"
}

variable "gcp_project" {
  type        = string
  description = "Project to use for this config"
  default     = "farma-compara"
}

variable "function_name" {
  description = "The name of the function desployed alongside linked resources"
  default     = "myFunction"
}

variable "entry_point" {
  description = "The entrypoint where the function is called"
  default     = "HelloWorld"
}
