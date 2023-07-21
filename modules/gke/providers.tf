terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "4.7.0"
    }
  }
}

provider "google" {
  project = "norse-rampart-393313"
  region  = "eu-west3"
  zone    = "eu-west3-c"
}
