data "aws_caller_identity" "self" { }

resource "aws_lambda_function" "delete_branch_bot" {
    filename = "handler.zip"
    function_name = "delete_branch_bot"
    role = "arn:aws:iam::${data.aws_caller_identity.self.account_id}:role/service-role/lambda"
    handler = "delete-branch"
    runtime = "go1.x"
    timeout = 150
}

