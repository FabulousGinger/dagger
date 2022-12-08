resource "aws_ecs_cluster" "ecs" {
  name = var.ecs_cluster_name
}

resource "aws_iam_role" "ecs_role" {
  name = "${aws_ecs_cluster.ecs.name}-ECSRole"
  path = "/"

  assume_role_policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ecs-tasks.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
POLICY
}

resource "aws_iam_policy" "ecs_service_logging" {
  name = "ecs-service-${aws_ecs_cluster.ecs.name}"

  policy = <<POLICY
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": [
                "logs:CreateLogGroup",
                "logs:CreateLogStream",
                "logs:PutLogEvents",
                "logs:DescribeLogGroups",
                "logs:DescribeLogStreams",
                "ecr:GetDownloadUrlForLayer",
                "ecr:BatchGetImage",
                "ecr:BatchCheckLayerAvailability",
                "ecr:PutImage",
                "ecr:InitiateLayerUpload",
                "ecr:UploadLayerPart",
                "ecr:CompleteLayerUpload",
                "ecr:GetAuthorizationToken",
                "secretsmanager:GetSecretValue"
            ],
            "Effect": "Allow",
            "Resource": "*"
        }
    ]
}
POLICY
}

resource "aws_iam_role_policy_attachment" "ecs_role_attachment_logging" {
  role       = aws_iam_role.ecs_role.name
  policy_arn = aws_iam_policy.ecs_service_logging.arn
}

resource "aws_cloudwatch_log_group" "log_group" {
  name = var.log_group_name
}