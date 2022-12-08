
#ALB
output "dns_name" {
  description = "The DNS name of the load balancer."
  value       = module.ecs_fargate.dns_name
}

output "zone_id" {
  description = "The canonical hosted zone ID of the load balancer."
  value       = module.ecs_fargate.zone_id
}

output "alb_arn" {
  description = "The ARN of the load balancer (matches id)."
  value       = module.ecs_fargate.alb_arn
}

output "listener_arn" {
  description = "The ARN of the listener (matches id)"
  value       = module.ecs_fargate.listener_arn
}

output "tg_arn" {
  description = "The ARN of the Target Group (matches id)"
  value       = module.ecs_fargate.tg_arn
}

#ECR
output "ecs_repository_url" {
  description = "The URL of the feature repository"
  value       = module.ecs_fargate.ecr_repository_url
}

#SG
output "fargate_security_group" {
  description = "The security group of Fargate."
  value       = module.ecs_fargate.fargate_security_group
}

output "alb_security_group" {
  description = "The security group of the ALB."
  value       = module.ecs_fargate.alb_security_group
}
