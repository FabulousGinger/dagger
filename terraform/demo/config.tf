provider "aws" {}

terraform {
  backend "s3" {
    bucket  = "fuzzplay-terraform"
    key     = "dagger/terraform.tfstate"
    region  = "us-east-1"
    profile = "fuzzplay"
  }
}
