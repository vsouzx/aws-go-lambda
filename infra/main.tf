provider "aws" {
  region = "us-east-1"
}

resource "aws_iam_role" "lambda_role" {
  name = "lambda-execution-role"

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

resource "aws_iam_role_policy_attachment" "lambda_basic_execution" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_lambda_function" "go_lambda" {
  filename         = "deployment.zip"
  function_name    = "go-lambda-example"
  role             = aws_iam_role.lambda_role.arn
  handler          = "main"
  source_code_hash = filebase64sha256("lambda.zip")
  runtime          = "go1.x"

  memory_size      = 128
  timeout          = 5
}

resource "aws_lambda_function_url" "lambda_url" {
  function_name      = aws_lambda_function.go_lambda.function_name
  authorization_type = "NONE"
}
