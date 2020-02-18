terraform {
  required_version = "> 0.12.0"
  required_providers {
    aws = "~> 2.33.0"
  }
  backend "s3" {
    profile        = "terraform-practice"
    bucket         = "terraform-ractice-tfstate"
    key            = "classes/terraform-101.tfstate"
    region         = "us-east-1"
    encrypt        = true
    kms_key_id     = "arn:aws:kms:us-east-1:AKIAZL74OL73ZSCSISXS:key/VIGQtYstC9BZ65Eg6bk+QkMcGap2Pct9bKQt/N8G"
    acl            = "private"
    dynamodb_table = "tfstate"
  }
}

provider "aws" {
  profile = "terraform-practice"
  region  = "us-east-1"
}

provider "ns1" {
  apikey = var.personal_ns1_apikey
}
