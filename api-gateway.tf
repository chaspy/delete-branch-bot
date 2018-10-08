resource "aws_api_gateway_rest_api" "delete-branch" {
  name        = "delete-branch-bot"
  description = "Delete merged branch when Pull Request merged"

  endpoint_configuration {
    types = ["REGIONAL"]
  }
}

resource "aws_api_gateway_deployment" "prod" {
  depends_on = [
    "aws_api_gateway_integration.delete-branch",
  ]
  rest_api_id = "${aws_api_gateway_rest_api.delete-branch.id}"
  stage_description = "production"
  stage_name = "prod"
}

# /post
resource "aws_api_gateway_resource" "post" {
  parent_id = "${aws_api_gateway_rest_api.delete-branch.root_resource_id}"
  path_part = "/"
  rest_api_id = "${aws_api_gateway_rest_api.delete-branch.id}"
}

resource "aws_api_gateway_integration" "delete-branch" {
  http_method = "${aws_api_gateway_method.post.http_method}"
  integration_http_method = "POST"
  request_templates = {
    "application/json" = <<EOF
{
  "id": "$input.params('id')",
  "_method": "POST"
}
EOF
  }
  resource_id = "${aws_api_gateway_resource.post.id}"
  rest_api_id = "${aws_api_gateway_rest_api.delete-branch.id}"
  type = "AWS"
  uri = "arn:aws:apigateway:${var.region}:lambda:path/2015-03-31/functions/arn:aws:lambda:${var.region}:${var.aws_account_id}:function:${aws_lambda_function.delete_branch_bot.function_name}/invocations"
}

resource "aws_api_gateway_method" "post" {
  authorization = "NONE"
  http_method = "POST"
  resource_id = "${aws_api_gateway_resource.post.id}"
  rest_api_id = "${aws_api_gateway_rest_api.delete-branch.id}"
}

resource "aws_api_gateway_method_response" "200" {
  http_method = "${aws_api_gateway_method.post.http_method}"
  resource_id = "${aws_api_gateway_resource.post.id}"
  rest_api_id = "${aws_api_gateway_rest_api.delete-branch.id}"
  status_code = "200"
}

resource "aws_api_gateway_integration_response" "post_200" {
  depends_on = ["aws_api_gateway_integration.delete-branch"]
  http_method = "${aws_api_gateway_method.post.http_method}"
  resource_id = "${aws_api_gateway_resource.post.id}"
  rest_api_id = "${aws_api_gateway_rest_api.delete-branch.id}"
  status_code = "${aws_api_gateway_method_response.200.status_code}"
}
