output "lambda_function_url" {
  description = "URL pública da Lambda"
  value       = aws_lambda_function_url.lambda_url.function_url
}
