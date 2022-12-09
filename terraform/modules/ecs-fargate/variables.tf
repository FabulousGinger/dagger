# ECS
variable "ecs_cluster_name" {
  description = "The name of the ECS cluster"
  type = string
}

# ALB
variable "alb_s3_bucket" {
  description = "The bucket name used to store ALB logs."
  type        = string
}

variable "alb_name" {
  description = "The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen"
  type        = string
}

variable "create_route53_record" {
  description = "Set to true to create a DNS record for the ALB"
  type        = bool
  default     = false
}

variable "alb_dns_record" {
  description = "The DNS record of the ALB to be set in Route53"
  type        = string
  default     = ""
}

variable "alb_internal" {
  description = "If true, the LB will be internal."
  type        = bool
  default     = false
}

variable "subnets" {
  description = "A list of subnet IDs to attach to the LB."
  type        = list(any)
}

variable "target_group_name" {
  description = "The name of the target group. If omitted, Terraform will assign a random, unique name."
  type        = string
}

variable "target_group_port" {
  description = "The port on which targets receive traffic."
  type        = number
  default     = 80
}

variable "target_group_protocol" {
  description = "The protocol to use for routing traffic to the targets. Should be one of TCP, TLS, UDP, TCP_UDP, HTTP or HTTPS."
  type        = string
  default     = "HTTP"
}

variable "vpc_id" {
  description = "The identifier of the VPC in which to create the target group."
  type        = string
}

variable "zone_id" {
  description = "The ID of the hosted zone to contain DNS records."
  type        = string
  default     = ""
}

variable "deregistration_delay" {
  description = "he amount time for Elastic Load Balancing to wait before changing the state of a deregistering target from draining to unused."
  type        = number
  default     = 5
}

variable "tg_health_check_path" {
  description = "The destination for the health check request."
  type        = string
  default     = "/" #/haproxy?monitor for haproxy
}

variable "tg_health_code" {
  description = "The HTTP codes to use when checking for a successful response from a target."
  type        = string
  default     = "200"
}

variable "listener_port" {
  description = "The listener port. Specify a value from 1 to 65535."
  type        = string
  default     = "443"
}

variable "listener_protocol" {
  description = "The listener protocol. Valid values are HTTP, HTTPS."
  type        = string
  default     = "HTTPS"
}

variable "listener_certificate_arn" {
  description = "The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS."
  type        = string
}

variable "additional_certificates" {
  description = "A list of ARN of the certificates to attach to the listener."
  type        = list(any)
  default     = []
}

variable "tags" {
  description = "A map of tags to add to all resources"
  type        = map(string)
  default     = {}
}

# ECR
variable "create_ecr_for_proxy" {
  description = "Set to true to create the ecr for proxy"
  type        = bool
  default     = false
}

variable "ecr_repository_name" {
  description = "The name of the ECR."
  type        = string
}

variable "proxy_name" {
  description = "The name of the ECR for the proxy."
  type        = string
  default     = ""
}

# SG
variable "alb_security_group_name" {
  description = "The name of the security group"
  type = string
}

variable "alb_security_group_description" {
  description = "The description of the security group"
  type = string
}

variable "network_vpc_cidr" {
  description = "The CIDR block of the VPC"
  type        = string
}

variable "alb_security_group_tags" {
  description = "A map of tags to add to resources"
  type        = map(string)
  default     = {}
}

variable "fargate_security_group_name" {
  description = "The name of the security group"
  type = string
}

variable "fargate_security_group_description" {
  description = "The description of the security group"
  type = string
  default = ""
}

variable "fargate_security_group_tags" {
  description = "A map of tags to add to resources"
  type        = map(string)
  default     = {}
}

# CloudWatch
variable "log_group_name" {
  description = "The name of the log group that the ECS cluster will be using."
  type        = string
}

# S3
variable "alb_s3_tags" {
  description = "A map of tags to add to all resources"
  type        = map(string)
  default     = {}
}