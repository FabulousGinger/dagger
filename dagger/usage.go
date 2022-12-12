package main

const (
	usage = `
Dagger CI Tool

Provide a task to run

options:
test					Run go test
sonar					Scan with SonarQube and upload to project
ecrlogin				Login to AWS ECR
deploy					Will run everything needed to Deploy to AWS ECS
pipeline				Will run test, terraform apply, build, and deploy 
git [ hash ] 				Git short hash, git rev-parse HEAD
task [ name ]				Create/update task definition on AWS ECS
service [ name ]			Create service on AWS ECS
terraform [ plan | apply | destroy ]	Run Terraform commands
`
	terraformUsage = `
Dagger CI Tool

Provide a subtask for Terraform [ plan | apply | destroy ]

dagger [terraform plan | terraform apply | terraform destroy ]
`
	gitUsage = `
Dagger CI Tool

Provide a subtask for Git [ hash ]

dagger [ git hash ]
`
	ECSTaskUsage = `
Dagger CI Tool

Provide a name for the ECS task to create

dagger [ task name ]
`
	ECSServiceUsage = `
Dagger CI Tool

Provide a name for the ECS service to create

dagger [ service name ]
`
)
