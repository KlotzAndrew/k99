terraform {
  backend "gcs" {
    bucket = "k99-project"
    prefix = "prod"
  }
}
