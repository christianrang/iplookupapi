provider "aws" {
    region = "us-east-1"
}

data "archive_file" "lambda_zip" {
    type        = "zip"
    source_file = "bin/lambda"
    output_path = "bin/lambda.zip"
}

resource "random_id" "unique_suffix" {
    byte_length = 2
}

# Buld Lambda Function

resource "aws_lambda_function" "lambda_func" {
    filename         = data.archive_file.lambda_zip.output_path
    function_name    = local.app_id
    handler          = "lambda"
    source_code_hash = base64sha256(data.archive_file.lambda_zip.output_path)
    runtime          = "go1.x"
    role             = aws_iam_role.lambda_exec.arn

    environment {
      variables = {
          VIRUSTOTAL_API_KEY = var.virustotal_api_key
      }
    }
}

# Build IAM role

resource "aws_iam_role" "lambda_exec" {
    name_prefix = local.app_id
 
    assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_policy_attachment" "role_attach" {
    name       = "policy-${local.app_id}"
    roles      = [aws_iam_role.lambda_exec.id]
    count      = length(var.iam_policy_arn)
    policy_arn = element(var.iam_policy_arn, count.index)
}

# Build API Gateway

resource "aws_api_gateway_rest_api" "api" {
  name = local.app_id
}

resource "aws_api_gateway_resource" "proxy" {
    path_part   = "{proxy+}"
    parent_id   = aws_api_gateway_rest_api.api.root_resource_id
    rest_api_id = aws_api_gateway_rest_api.api.id
}

resource "aws_api_gateway_method" "method" {
    rest_api_id   = aws_api_gateway_rest_api.api.id
    resource_id   = aws_api_gateway_resource.proxy.id
    http_method   = "ANY"
    authorization = "NONE"
}

resource "aws_api_gateway_method" "root" {
    rest_api_id   = aws_api_gateway_rest_api.api.id
    resource_id   = aws_api_gateway_rest_api.api.root_resource_id
    http_method   = "ANY"
    authorization = "NONE"
}

resource "aws_api_gateway_integration" "integration" {
    rest_api_id             = aws_api_gateway_rest_api.api.id
    resource_id             = aws_api_gateway_method.method.resource_id
    http_method             = aws_api_gateway_method.method.http_method
    integration_http_method = "POST"
    type                    = "AWS_PROXY"
    uri                     = aws_lambda_function.lambda_func.invoke_arn
}

resource "aws_api_gateway_integration" "integration_root" {
    rest_api_id             = aws_api_gateway_rest_api.api.id
    resource_id             = aws_api_gateway_method.root.resource_id
    http_method             = aws_api_gateway_method.root.http_method
    integration_http_method = "POST"
    type                    = "AWS_PROXY"
    uri                     = aws_lambda_function.lambda_func.invoke_arn
}

resource "aws_api_gateway_deployment" "api_deployment" {
    depends_on = [
      aws_api_gateway_integration.integration,
      aws_api_gateway_integration.integration_root,
    ]

    rest_api_id = aws_api_gateway_rest_api.api.id
    stage_name  = "api"
}

resource "aws_lambda_permission" "lambda_permission" {
    action        = "lambda:InvokeFunction"
    function_name = aws_lambda_function.lambda_func.arn
    principal     = "apigateway.amazonaws.com"
    source_arn    = "${aws_api_gateway_deployment.api_deployment.execution_arn}/*/*"
}
