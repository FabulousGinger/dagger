resource "aws_ecr_repository" "proxy" {
  count = var.create_ecr_for_proxy ? 1 : 0

  name = var.proxy_name
}

resource "aws_ecr_repository" "this" {
  name = var.ecr_repository_name
}
