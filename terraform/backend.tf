terraform {
  backend "s3" {
    bucket = "<s3-bucket-name>" // s3 bucket to store state file
    key    = "curricular-api.tfstate"
  }
}
