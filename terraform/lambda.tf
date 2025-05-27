resource "aws_lambda_function" "curricular_api" {
  function_name = "curricular-api"

  filename         = data.archive_file.lambda.output_path
  source_code_hash = data.archive_file.lambda.output_base64sha256

  # "bootstrap" is the executable filename within the zip file
  handler     = "bootstrap"
  runtime     = "provided.al2"
  timeout     = "30"
  memory_size = "128"

  role = aws_iam_role.lambda_execution_role.arn

  vpc_config {
    subnet_ids = data.aws_subnets.subnet.ids
    security_group_ids = [
      data.aws_security_group.security_group.id
    ]
  }
}
