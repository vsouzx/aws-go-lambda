resource "aws_iam_role" "iam_for_lambda"{
    name                    = "iam_for_lambda"
    assume_role_policy      = data.aws_iam_policy_document.assume_role.json
}

resource "aws_iam_policy" "lambda_logging" {
    name = "go_lambda_logging"
    path = "/"
    description = "IAM policy for logging from a lambda"

    policy = <<EOF
    {
        "Version":"2012-10-17",
        "Statement": [
            {
                "Action": [
                    "logs:CreateLogGroup",
                    "logs:CreateLogStream",
                    "logs:PutLogEvents"
                ],
                "Resource": "arn:aws:logs:*:*:*",
                "Effect": "Allow"
            }
        ]
    }
    EOF
}

resource "aws_iam_role_policy_attachment" "lambda_logs" {
    role = aws_iam_role.iam_for_lambda.name
    policy_arn = aws_iam_policy.lambda_logging.arn
}

resource "aws_lambda_function" "lambda" {
  function_name    = "go_hello_lambda_function"
  role             = aws_iam_role.iam_for_lambda.arn
  runtime          = "provided.al2"
  handler          = "bootstrap"
  filename         = "lambda.zip"
  source_code_hash = filebase64sha256("${path.module}/lambda.zip")
  
  environment {
    variables = {
      DB_USERNAME = local.db_secret.db_username
      DB_PASSWORD = local.db_secret.db_password
      DB_URL      = aws_db_instance.default.endpoint
      DB_NAME     = aws_db_instance.default.db_name
    }
  }

  vpc_config {
    subnet_ids         = ["subnet-0417119060c4d0943", "subnet-0ce0261f2c16e7733", "subnet-0263f30e9f8071b6a",
                         "subnet-0ab0b479fc261d026", "subnet-06016dabbff1e0e0a", "subnet-0d746cbbd27803abf"]
    security_group_ids = [aws_security_group.lambda_sg.id]
  }
}

resource "aws_cloudwatch_log_group" "example"{
    name = "/aws/lambda/${aws_lambda_function.lambda.function_name}"
    retention_in_days = var.log_retention_days
}