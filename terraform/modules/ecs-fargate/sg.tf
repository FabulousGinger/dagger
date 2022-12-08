# alb security group
resource "aws_security_group" "alb_sg" {
  name        = var.alb_security_group_name
  description = var.alb_security_group_description
  vpc_id      = var.vpc_id

  # inbound HTTP from anywhere
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # inbound HTTPS from anywhere
  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = var.alb_security_group_tags
}

# ecs fargate security group
resource "aws_security_group" "fargate" {
  name        = var.fargate_security_group_name
  description = var.fargate_security_group_description
  vpc_id      = var.vpc_id

  # outbound all
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = var.fargate_security_group_tags
}

# inbound all to ecs fargate from VPC
resource "aws_security_group_rule" "ecs_ingress_from_vpc" {
  type        = "ingress"
  from_port   = 0
  to_port     = 65535
  protocol    = "tcp"
  cidr_blocks = [var.network_vpc_cidr]

  security_group_id = aws_security_group.fargate.id
}

# inbound all to ecs fargate from ALB
resource "aws_security_group_rule" "ecs_ingress_all" {
  type      = "ingress"
  from_port = 0
  to_port   = 65535
  protocol  = "tcp"

  security_group_id        = aws_security_group.fargate.id
  source_security_group_id = aws_security_group.alb_sg.id
}

# outboud traffic from ALB to ecs fargate
resource "aws_security_group_rule" "alb_egress_all" {
  type      = "egress"
  from_port = 0
  to_port   = 65535
  protocol  = "TCP"

  security_group_id        = aws_security_group.alb_sg.id
  source_security_group_id = aws_security_group.fargate.id
}
