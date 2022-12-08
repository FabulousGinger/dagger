# aws-ecs-fargate

This Terraform module will create the following for a Fargate ECS cluster:  
ALB with Target group, ECR for ECS, ECR for proxy, security group for  
Fargate, security group for ALB, and the ECS cluster.

## Requirements

| Name | Version |
|------|---------|
| aws | ~> 3.0 |

## Providers

| Name | Version |
|------|---------|
| aws | ~> 3.0 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| additional\_certificates | A list of ARN of the certificates to attach to the listener. | `list` | `[]` | no |
| alb\_dns\_record | The DNS record of the ALB to be set in Route53 | `string` | n/a | yes |
| alb\_internal | If true, the LB will be internal. | `bool` | `false` | no |
| alb\_name | The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen | `string` | n/a | yes |
| alb\_s3\_bucket | The bucket name used to store ALB logs. | `string` | n/a | yes |
| deregistration\_delay | he amount time for Elastic Load Balancing to wait before changing the state of a deregistering target from draining to unused. | `number` | `300` | no |
| environment | The environment name. EX. dev, staging, production. | `string` | n/a | yes |
| listener\_certificate\_arn | The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. | `string` | n/a | yes |
| listener\_port | The listener port. Specify a value from 1 to 65535. | `string` | `"443"` | no |
| listener\_protocol | The listener protocol. Valid values are HTTP, HTTPS. | `string` | `"HTTPS"` | no |
| log\_group\_name | The name of the log group that the ECS cluster will be using. | `string` | n/a | yes |
| network\_vpc\_cidr | The CIDR block of the VPC | `string` | n/a | yes |
| proxy\_name | The name of the ECR for the proxy. | `string` | n/a | yes |
| subnets | A list of subnet IDs to attach to the LB. | `list` | n/a | yes |
| target\_group\_name | The name of the target group. If omitted, Terraform will assign a random, unique name. | `string` | n/a | yes |
| target\_group\_port | The port on which targets receive traffic. | `number` | `80` | no |
| target\_group\_protocol | The protocol to use for routing traffic to the targets. Should be one of TCP, TLS, UDP, TCP\_UDP, HTTP or HTTPS. | `string` | `"HTTP"` | no |
| tg\_health\_check\_path | The destination for the health check request. | `string` | `"/"` | no |
| tg\_health\_code | The HTTP codes to use when checking for a successful response from a target. | `string` | `"200"` | no |
| vpc\_id | The identifier of the VPC in which to create the target group. | `string` | n/a | yes |
| zone\_id | The ID of the hosted zone to contain DNS records. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| alb\_arn | The ARN of the load balancer (matches id). |
| alb\_security\_group | The security group of the ALB. |
| dns\_name | The DNS name of the load balancer. |
| ecs\_role\_arn | The Amazon Resource Name (ARN) specifying the role. |
| fargate\_security\_group | The security group of Fargate. |
| feature\_repository\_url | The URL of the feature repository |
| listener\_arn | The ARN of the listener (matches id) |
| proxy\_repository\_url | The URL of the proxy repository |
| tg\_arn | The ARN of the Target Group (matches id) |
| zone\_id | The canonical hosted zone ID of the load balancer. |


<!-- BEGIN_TF_DOCS -->
# aws-ecs-fargate

This Terraform module will create the following for a Fargate ECS cluster:
ALB with Target group, ECR for ECS, ECR for proxy, security group for
Fargate, security group for ALB, and the ECS cluster.

## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.3.5 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 4.43.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | ~> 4.43.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_alb.alb](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/alb) | resource |
| [aws_alb_listener.alb_listener](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/alb_listener) | resource |
| [aws_alb_listener.redirect_listener](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/alb_listener) | resource |
| [aws_alb_target_group.alb_tg](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/alb_target_group) | resource |
| [aws_cloudwatch_log_group.log_group](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_log_group) | resource |
| [aws_ecr_repository.proxy](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecr_repository) | resource |
| [aws_ecr_repository.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecr_repository) | resource |
| [aws_ecs_cluster.ecs](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecs_cluster) | resource |
| [aws_iam_policy.ecs_service_logging](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_policy) | resource |
| [aws_iam_role.ecs_role](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role) | resource |
| [aws_iam_role_policy_attachment.ecs_role_attachment_logging](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy_attachment) | resource |
| [aws_lb_listener_certificate.additional-certs](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lb_listener_certificate) | resource |
| [aws_route53_record.alb-route53](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route53_record) | resource |
| [aws_s3_bucket.alb](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket) | resource |
| [aws_security_group.alb_sg](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group) | resource |
| [aws_security_group.fargate](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group) | resource |
| [aws_security_group_rule.alb_egress_all](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group_rule) | resource |
| [aws_security_group_rule.ecs_ingress_all](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group_rule) | resource |
| [aws_security_group_rule.ecs_ingress_from_vpc](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group_rule) | resource |
| [aws_elb_service_account.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/elb_service_account) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_additional_certificates"></a> [additional\_certificates](#input\_additional\_certificates) | A list of ARN of the certificates to attach to the listener. | `list(any)` | `[]` | no |
| <a name="input_alb_dns_record"></a> [alb\_dns\_record](#input\_alb\_dns\_record) | The DNS record of the ALB to be set in Route53 | `string` | `""` | no |
| <a name="input_alb_internal"></a> [alb\_internal](#input\_alb\_internal) | If true, the LB will be internal. | `bool` | `false` | no |
| <a name="input_alb_name"></a> [alb\_name](#input\_alb\_name) | The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen | `string` | n/a | yes |
| <a name="input_alb_s3_bucket"></a> [alb\_s3\_bucket](#input\_alb\_s3\_bucket) | The bucket name used to store ALB logs. | `string` | n/a | yes |
| <a name="input_create_ecr_for_proxy"></a> [create\_ecr\_for\_proxy](#input\_create\_ecr\_for\_proxy) | Set to true to create the ecr for proxy | `bool` | `false` | no |
| <a name="input_create_route53_record"></a> [create\_route53\_record](#input\_create\_route53\_record) | Set to true to create a DNS record for the ALB | `bool` | `false` | no |
| <a name="input_deregistration_delay"></a> [deregistration\_delay](#input\_deregistration\_delay) | he amount time for Elastic Load Balancing to wait before changing the state of a deregistering target from draining to unused. | `number` | `300` | no |
| <a name="input_ecr_repository_name"></a> [ecr\_repository\_name](#input\_ecr\_repository\_name) | The name of the ECR. | `string` | n/a | yes |
| <a name="input_listener_certificate_arn"></a> [listener\_certificate\_arn](#input\_listener\_certificate\_arn) | The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. | `string` | n/a | yes |
| <a name="input_listener_port"></a> [listener\_port](#input\_listener\_port) | The listener port. Specify a value from 1 to 65535. | `string` | `"443"` | no |
| <a name="input_listener_protocol"></a> [listener\_protocol](#input\_listener\_protocol) | The listener protocol. Valid values are HTTP, HTTPS. | `string` | `"HTTPS"` | no |
| <a name="input_log_group_name"></a> [log\_group\_name](#input\_log\_group\_name) | The name of the log group that the ECS cluster will be using. | `string` | n/a | yes |
| <a name="input_network_vpc_cidr"></a> [network\_vpc\_cidr](#input\_network\_vpc\_cidr) | The CIDR block of the VPC | `string` | n/a | yes |
| <a name="input_proxy_name"></a> [proxy\_name](#input\_proxy\_name) | The name of the ECR for the proxy. | `string` | `""` | no |
| <a name="input_subnets"></a> [subnets](#input\_subnets) | A list of subnet IDs to attach to the LB. | `list(any)` | n/a | yes |
| <a name="input_target_group_name"></a> [target\_group\_name](#input\_target\_group\_name) | The name of the target group. If omitted, Terraform will assign a random, unique name. | `string` | n/a | yes |
| <a name="input_target_group_port"></a> [target\_group\_port](#input\_target\_group\_port) | The port on which targets receive traffic. | `number` | `80` | no |
| <a name="input_target_group_protocol"></a> [target\_group\_protocol](#input\_target\_group\_protocol) | The protocol to use for routing traffic to the targets. Should be one of TCP, TLS, UDP, TCP\_UDP, HTTP or HTTPS. | `string` | `"HTTP"` | no |
| <a name="input_tg_health_check_path"></a> [tg\_health\_check\_path](#input\_tg\_health\_check\_path) | The destination for the health check request. | `string` | `"/"` | no |
| <a name="input_tg_health_code"></a> [tg\_health\_code](#input\_tg\_health\_code) | The HTTP codes to use when checking for a successful response from a target. | `string` | `"200"` | no |
| <a name="input_vpc_id"></a> [vpc\_id](#input\_vpc\_id) | The identifier of the VPC in which to create the target group. | `string` | n/a | yes |
| <a name="input_zone_id"></a> [zone\_id](#input\_zone\_id) | The ID of the hosted zone to contain DNS records. | `string` | `""` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_alb_arn"></a> [alb\_arn](#output\_alb\_arn) | The ARN of the load balancer (matches id). |
| <a name="output_alb_security_group"></a> [alb\_security\_group](#output\_alb\_security\_group) | The security group of the ALB. |
| <a name="output_dns_name"></a> [dns\_name](#output\_dns\_name) | The DNS name of the load balancer. |
| <a name="output_ecr_repository_url"></a> [ecr\_repository\_url](#output\_ecr\_repository\_url) | The URL of the ECR |
| <a name="output_ecs_role_arn"></a> [ecs\_role\_arn](#output\_ecs\_role\_arn) | The Amazon Resource Name (ARN) specifying the role. |
| <a name="output_fargate_security_group"></a> [fargate\_security\_group](#output\_fargate\_security\_group) | The security group of Fargate. |
| <a name="output_listener_arn"></a> [listener\_arn](#output\_listener\_arn) | The ARN of the listener (matches id) |
| <a name="output_proxy_repository_url"></a> [proxy\_repository\_url](#output\_proxy\_repository\_url) | The URL of the proxy repository |
| <a name="output_tg_arn"></a> [tg\_arn](#output\_tg\_arn) | The ARN of the Target Group (matches id) |
| <a name="output_zone_id"></a> [zone\_id](#output\_zone\_id) | The canonical hosted zone ID of the load balancer. |
<!-- END_TF_DOCS -->