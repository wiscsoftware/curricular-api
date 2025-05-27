resource "aws_api_gateway_rest_api" "curricular_api" {
  name        = "Curricular API"
  description = "An API for exposing teacher, class, course data"
}

resource "aws_api_gateway_method_settings" "curricular_api" {
  rest_api_id = aws_api_gateway_rest_api.curricular_api.id
  stage_name  = aws_api_gateway_stage.curricular_api.stage_name
  method_path = "*/*"

  settings {
    logging_level = "INFO"
  }
}

resource "aws_api_gateway_resource" "proxy" {
  rest_api_id = aws_api_gateway_rest_api.curricular_api.id
  parent_id   = aws_api_gateway_rest_api.curricular_api.root_resource_id
  path_part   = "{proxy+}"
}

resource "aws_api_gateway_method" "method" {
  rest_api_id      = aws_api_gateway_rest_api.curricular_api.id
  resource_id      = aws_api_gateway_resource.proxy.id
  http_method      = "ANY" // both GET and POST
  authorization    = "NONE"
  api_key_required = true
}

resource "aws_api_gateway_integration" "get_student_integration" {
  rest_api_id = aws_api_gateway_rest_api.curricular_api.id
  resource_id = aws_api_gateway_resource.proxy.id
  http_method = aws_api_gateway_method.method.http_method

  // https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-lambda-integrations.html
  integration_http_method = "POST" // Lambda integration is through POST
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.curricular_api.invoke_arn
}

resource "aws_api_gateway_deployment" "curricular_api" {
  rest_api_id = aws_api_gateway_rest_api.curricular_api.id

  lifecycle {
    create_before_destroy = true
  }

  triggers = {
    redeployment = sha1(jsonencode([
      aws_api_gateway_resource.proxy,
      aws_api_gateway_method.method,
      aws_api_gateway_integration.get_student_integration,
    ]))
  }
}

resource "aws_api_gateway_stage" "curricular_api" {
  deployment_id = aws_api_gateway_deployment.curricular_api.id
  rest_api_id   = aws_api_gateway_rest_api.curricular_api.id
  stage_name    = "curricular-and-academic-api"
}

resource "aws_lambda_permission" "apigw" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.curricular_api.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.curricular_api.execution_arn}/*/*"
}

resource "aws_api_gateway_usage_plan" "curricular_api" {
  name = "curricular_api"

  depends_on = [
    aws_api_gateway_method_settings.curricular_api,
    aws_api_gateway_stage.curricular_api
  ]

  api_stages {
    api_id = aws_api_gateway_rest_api.curricular_api.id
    stage  = aws_api_gateway_stage.curricular_api.stage_name
  }
}

resource "aws_api_gateway_api_key" "curricular_and_academic_api" {
  name    = "CurricularApiKey"
  enabled = true
}

resource "aws_api_gateway_usage_plan_key" "curricular_and_academic_api" {
  key_id        = aws_api_gateway_api_key.curricular_and_academic_api.id
  key_type      = "API_KEY"
  usage_plan_id = aws_api_gateway_usage_plan.curricular_api.id
}
