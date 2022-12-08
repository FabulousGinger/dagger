module "vpc" {
  source = "../modules/vpc"

  name            = local.name
  cidr            = "10.0.0.0/16"
  azs             = ["us-east-2a", "us-east-2b", "us-east-2c"]
  private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  public_subnets  = ["10.0.101.0/24", "10.0.102.0/24", "10.0.103.0/24"]

  enable_dns_hostnames = true
  enable_dns_support   = true

  enable_nat_gateway = true
  single_nat_gateway = true

  public_subnet_tags = {
    Name = "${local.name}-public"
  }

  private_subnet_tags = {
    Name = "${local.name}-private"
  }

  tags = {
    Environment = "demo"
  }

  vpc_tags = {
    Name = local.name
  }
}

module "ecs_fargate" {
  source = "../modules/ecs-fargate"

  ecs_cluster_name = local.name

  alb_s3_bucket                  = "${local.name}-alb"
  alb_name                       = local.name
  alb_security_group_description = "Security Group for dagger demo ALB"
  alb_security_group_name        = "${local.name}-alb"

  target_group_name    = local.name
  tg_health_check_path = "/"
  target_group_port    = 8080

  fargate_security_group_name        = "${local.name}-fargate"
  fargate_security_group_description = "Security Group for dagger demo ECS"

  create_route53_record    = true
  alb_dns_record           = local.name
  zone_id                  = "Z05898413EB6JMZYB6JQ5"
  listener_certificate_arn = "arn:aws:acm:us-east-2:650405159858:certificate/4693c334-2788-450c-8ca2-11aad881a988"

  ecr_repository_name = local.name

  log_group_name = local.name

  vpc_id           = module.vpc.vpc_id
  subnets          = module.vpc.public_subnets
  network_vpc_cidr = module.vpc.vpc_cidr_block

  depends_on = [
    module.vpc
  ]
}
