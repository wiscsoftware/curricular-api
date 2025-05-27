terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.40.0"
    }
  }
}

provider "aws" {
  default_tags {
    tags = {
      "user:CostCenter" = "Curricular API"
      CreatedBy            = "Terraform"
    }
  }

  allowed_account_ids = var.aws_account_ids
}
