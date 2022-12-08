data "aws_elb_service_account" "main" {}

resource "aws_s3_bucket" "alb" {
  bucket = var.alb_s3_bucket
  tags = var.alb_s3_tags
}

resource "aws_s3_bucket_policy" "this" {
  bucket = aws_s3_bucket.alb.id
  policy = data.aws_iam_policy_document.this.json
}

data "aws_iam_policy_document" "this" {
  statement {
    principals {
      type        = "AWS"
      identifiers = [data.aws_elb_service_account.main.arn]
    }

    actions = [
      "s3:GetObject",
      "s3:ListBucket",
      "s3:PutObject"
    ]

    resources = [
      aws_s3_bucket.alb.arn,
      "${aws_s3_bucket.alb.arn}/*",
    ]
  }
}

resource "aws_s3_bucket_acl" "this" {
  bucket = aws_s3_bucket.alb.id
  acl    = "private"
}

resource "aws_alb" "alb" {
  name            = var.alb_name
  internal        = var.alb_internal
  security_groups = [aws_security_group.alb_sg.id]
  subnets         = var.subnets

  enable_deletion_protection = true

  access_logs {
    enabled = true
    bucket  = aws_s3_bucket.alb.bucket
  }

  tags = var.tags
}

resource "aws_alb_target_group" "alb_tg" {
  name                 = var.target_group_name
  port                 = var.target_group_port
  protocol             = var.target_group_protocol
  vpc_id               = var.vpc_id
  deregistration_delay = var.deregistration_delay
  target_type          = "ip"

  health_check {
    path    = var.tg_health_check_path
    matcher = var.tg_health_code
  }
}

resource "aws_alb_listener" "redirect_listener" {
  load_balancer_arn = aws_alb.alb.arn
  port              = "80"
  protocol          = "HTTP"

  default_action {
    type = "redirect"

    redirect {
      port        = var.listener_port
      protocol    = var.listener_protocol
      status_code = "HTTP_301"
    }
  }
}

resource "aws_alb_listener" "alb_listener" {
  load_balancer_arn = aws_alb.alb.arn
  port              = var.listener_port
  protocol          = var.listener_protocol
  ssl_policy        = "ELBSecurityPolicy-2016-08"
  certificate_arn   = var.listener_certificate_arn

  default_action {
    target_group_arn = aws_alb_target_group.alb_tg.arn
    type             = "forward"
  }
}

resource "aws_lb_listener_certificate" "additional-certs" {
  count = length(var.additional_certificates)

  listener_arn    = aws_alb_listener.alb_listener.arn
  certificate_arn = element(var.additional_certificates, count.index)
}

resource "aws_route53_record" "alb-route53" {
  count = var.create_route53_record ? 1 : 0

  zone_id = var.zone_id
  name    = var.alb_dns_record
  type    = "A"

  alias {
    name                   = aws_alb.alb.dns_name
    zone_id                = aws_alb.alb.zone_id
    evaluate_target_health = true
  }

}