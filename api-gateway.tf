resource "aws_api_gateway_rest_api" "delete-branch2" {
  name        = "delete-branch-bot2"
  description = "Delete merged branch when Pull Request merged"

  endpoint_configuration {
    types = ["REGIONAL"]
  }
}
