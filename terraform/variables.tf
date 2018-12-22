variable "project_name" {
  type = "string"
  default = "k99-project"
}

variable "region" {
  type = "string"
  default = "east1-b"
}

variable "zone" {
  type = "string"
  default = "us-east1-b"
}

variable "min_node_count" {
  type = "string"
  default = 3
}

variable "max_node_count" {
  type = "string"
  default = 3
}

variable "billing_account" {
  description = "Billing account STRING."
}

variable "org_id" {
  description = "Organisation account NR."
}
