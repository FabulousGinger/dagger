resource "aws_flow_log" "flow_log" {
  iam_role_arn    = aws_iam_role.flow_log_iam_role.arn
  log_destination = aws_cloudwatch_log_group.flow_log_group.arn
  traffic_type    = "REJECT"
  vpc_id          = concat(aws_vpc.this.*.id, [""])[0]
}

resource "aws_cloudwatch_log_group" "flow_log_group" {
  name = "${var.name}-flow-log"
}

resource "aws_iam_role" "flow_log_iam_role" {
  name = "${var.name}-role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "vpc-flow-logs.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "flow_log_iam_role_policy" {
  name = "${var.name}-role-policy"
  role = aws_iam_role.flow_log_iam_role.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents",
        "logs:DescribeLogGroups",
        "logs:DescribeLogStreams"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF
}
