terraform {
  required_version = ">= 0.11.0"

  backend "s3" {
    bucket = "delete-branch-bot"
    key    = "delete-branch-bot.tfstate"
    region = "ap-northeast-1"
  }
}

variable "aws_access_key" {}
variable "aws_secret_key" {}
variable "region" {
    default = "ap-northeast-1"
}

provider "aws" {
  access_key = "${var.aws_access_key}"
  secret_key = "${var.aws_secret_key}"
  region     = "${var.region}"
}
