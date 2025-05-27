data "aws_region" "current" {}

data "aws_vpc" "vpc" {
  tags = {
    Name        = "default"
  }
}

data "aws_subnets" "subnet" {
  filter {
    name   = "vpc-id"
    values = [data.aws_vpc.vpc.id]
  }
}

data "aws_security_group" "security_group" {
  vpc_id = data.aws_vpc.vpc.id

  tags = {
    Name = var.sec_group
  }
}

data "archive_file" "lambda" {
  type             = "zip"
  source_file      = "../src/bootstrap"
  output_file_mode = "0666"
  output_path      = "../lambda_function.zip"
}
