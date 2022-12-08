
#ALB
output "dns_name" {
  description = "The DNS name of the load balancer."
  value       = aws_alb.alb.dns_name
}

output "zone_id" {
  description = "The canonical hosted zone ID of the load balancer."
  value       = aws_alb.alb.zone_id
}

output "alb_arn" {
  description = "The ARN of the load balancer (matches id)."
  value       = aws_alb.alb.arn
}

output "listener_arn" {
  description = "The ARN of the listener (matches id)"
  value       = aws_alb_listener.alb_listener.arn
}

output "tg_arn" {
  description = "The ARN of the Target Group (matches id)"
  value       = aws_alb_target_group.alb_tg.arn
}

#ECR
output "ecr_repository_url" {
  description = "The URL of the ECR"
  value       = aws_ecr_repository.this.repository_url
}

output "proxy_repository_url" {
  description = "The URL of the proxy repository"
  value       = one(aws_ecr_repository.proxy[*].repository_url)
}

#SG
output "fargate_security_group" {
  description = "The security group of Fargate."
  value       = aws_security_group.fargate.id
}

output "alb_security_group" {
  description = "The security group of the ALB."
  value       = aws_security_group.alb_sg.id
}

#IAM
output "ecs_role_arn" {
  description = "The Amazon Resource Name (ARN) specifying the role."
  value       = aws_iam_role.ecs_role.arn
}
