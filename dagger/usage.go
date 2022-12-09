package main

const (
	usage = `
Dagger CI Tool

Provide a task to run

options:
test
ecrlogin
deploy
git [ hash ]
task [ name ]
service [ name ]
tf [ plan | apply | destroy ]
`
	tfUsage = `
Dagger CI Tool

Provide a subtask for Terraform [ plan | apply | destroy ]

dagger [tf plan | tf apply | tf destroy ]
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
