# ============== LAMBDA ROLE
resource "aws_iam_role" "go_lambda_validate_tls_role" {
  name = "GoLambdaValidateTLS-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Principal = {
        Service = "lambda.amazonaws.com"
      }
    }]
  })
}

resource "aws_iam_role_policy_attachment" "go_lambda_validate_tls_basic" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  role       = aws_iam_role.go_lambda_validate_tls_role.name
}

# ============== LAMBDA FUNCTION
resource "aws_lambda_function" "go_lambda_validate_tls_api" {
  function_name = "GoValidateTLSAPI"
  handler       = "bootstrap"
  runtime       = "provided.al2023"
  role          = aws_iam_role.go_lambda_validate_tls_role.arn
  filename      = "../api/build.zip"
}

# ============== API GATEWAY
resource "aws_apigatewayv2_api" "api_gw" {
  name          = "ValidateTLSGW"
  protocol_type = "HTTP"
  cors_configuration {
    allow_origins = ["*"]
    allow_methods = ["GET", "OPTIONS"]
  }
}

resource "aws_apigatewayv2_integration" "api_gw_integration" {
  api_id                 = aws_apigatewayv2_api.api_gw.id
  integration_type       = "AWS_PROXY"
  integration_method     = "GET"
  integration_uri        = aws_lambda_function.go_lambda_validate_tls_api.arn
  payload_format_version = "2.0"
}

resource "aws_apigatewayv2_route" "api_gw_route_proxy" {
  api_id    = aws_apigatewayv2_api.api_gw.id
  route_key = "GET /{proxy+}"
  target    = "integrations/${aws_apigatewayv2_integration.api_gw_integration.id}"
}

resource "aws_apigatewayv2_stage" "api_gw_prod_stage" {
  api_id      = aws_apigatewayv2_api.api_gw.id
  name        = "prod"
  auto_deploy = true
}

resource "aws_lambda_permission" "apigw_lambda_permission" {
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.go_lambda_validate_tls_api.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.api_gw.execution_arn}/*/*"
}
