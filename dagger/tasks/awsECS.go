package tasks

import (
	"context"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
)

func ECSDeploy(ctx context.Context) (err error) {

	awsRegion := os.Getenv("AWS_DEFAULT_REGION")
	taskDefinitionName := os.Getenv("TASK_DEFINITION_NAME")
	clusterName := os.Getenv("CLUSTER_NAME")
	serviceName := os.Getenv("ECS_SERVICE_NAME")

	_, err = Push(ctx)
	CheckIfError(err)

	sess := session.Must(session.NewSession())

	svc := ecs.New(sess, aws.NewConfig().WithRegion(awsRegion))

	taskDefinition, err := ECSFargateTask(taskDefinitionName)
	CheckIfError(err)

	input := &ecs.UpdateServiceInput{
		Service:            aws.String(serviceName),
		Cluster:            aws.String(clusterName),
		TaskDefinition:     aws.String(taskDefinition),
		ForceNewDeployment: aws.Bool(true),
	}

	Info("Updating Service")
	result, err := svc.UpdateService(input)
	CheckIfError(err)

	Info("Waiting for Service to be stable")
	err = svc.WaitUntilServicesStable(&ecs.DescribeServicesInput{
		Cluster: aws.String(clusterName),
		Services: []*string{
			aws.String(serviceName),
		},
	})
	CheckIfError(err)

	Info("Result: %s", result)
	Info("Deployment Successful!")

	return nil
}

func ECSFargateService(name string) (err error) {

	awsRegion := os.Getenv("AWS_DEFAULT_REGION")
	clusterName := os.Getenv("CLUSTER_NAME")
	containerName := os.Getenv("CONTAINER_NAME")
	targetGroupARN := os.Getenv("TARGET_GROUP_ARN")
	taskDefinitionName := os.Getenv("TASK_DEFINITION_NAME")

	subnets := os.Getenv("SUBNETS")
	sliceSubnets := strings.Split(subnets, ",")

	securityGroups := os.Getenv("SECURITY_GROUPS")
	sliceSecurityGroups := strings.Split(securityGroups, ",")

	containerPort, err := strconv.Atoi(os.Getenv("CONTAINER_PORT"))
	CheckIfError(err)

	count, err := strconv.Atoi(os.Getenv("ECS_SERVICE_COUNT"))
	CheckIfError(err)

	sess := session.Must(session.NewSession())

	svc := ecs.New(sess, aws.NewConfig().WithRegion(awsRegion))

	input := &ecs.CreateServiceInput{
		Cluster:      aws.String(clusterName),
		DesiredCount: aws.Int64(int64(count)),
		LoadBalancers: []*ecs.LoadBalancer{
			{
				ContainerName:  aws.String(containerName),
				ContainerPort:  aws.Int64(int64(containerPort)),
				TargetGroupArn: aws.String(targetGroupARN),
			},
		},
		ServiceName:    aws.String(name),
		TaskDefinition: aws.String(taskDefinitionName),
		LaunchType:     aws.String("FARGATE"),
		NetworkConfiguration: &ecs.NetworkConfiguration{
			AwsvpcConfiguration: &ecs.AwsVpcConfiguration{
				Subnets:        aws.StringSlice(sliceSubnets),
				SecurityGroups: aws.StringSlice(sliceSecurityGroups),
				AssignPublicIp: aws.String("ENABLED"),
			},
		},
	}
	result, err := svc.CreateService(input)
	CheckIfError(err)

	Info("Result: %s", result)

	return
}

func ECSFargateTask(name string) (taskDefinition string, err error) {
	awsRegion := os.Getenv("AWS_DEFAULT_REGION")
	awsLogGroup := os.Getenv("AWS_LOG_GROUP")
	ecsRole := os.Getenv("ECS_ROLE")
	containerName := os.Getenv("CONTAINER_NAME")
	cpu := os.Getenv("ECS_FARGATE_CPU")
	memory := os.Getenv("ECS_FARGATE_MEMORY")
	containerPort, err := strconv.Atoi(os.Getenv("CONTAINER_PORT"))
	CheckIfError(err)

	sess := session.Must(session.NewSession())

	svc := ecs.New(sess, aws.NewConfig().WithRegion(awsRegion))

	ctx := context.Background()

	address, err := Push(ctx)
	CheckIfError(err)

	input := &ecs.RegisterTaskDefinitionInput{
		ContainerDefinitions: []*ecs.ContainerDefinition{
			{
				Essential: aws.Bool(true),
				Image:     aws.String(address),
				Name:      aws.String(containerName),
				PortMappings: []*ecs.PortMapping{
					{
						ContainerPort: aws.Int64(int64(containerPort)),
						HostPort:      aws.Int64(int64(containerPort)),
						Protocol:      aws.String("tcp"),
					},
				},
				LogConfiguration: &ecs.LogConfiguration{
					LogDriver: aws.String("awslogs"),
					Options: map[string]*string{
						"awslogs-group":         aws.String(awsLogGroup),
						"awslogs-region":        aws.String(awsRegion),
						"awslogs-stream-prefix": aws.String("ecs"),
					},
				},
			},
		},
		Family:           aws.String(name),
		ExecutionRoleArn: aws.String(ecsRole),
		RequiresCompatibilities: aws.StringSlice([]string{
			"FARGATE",
		}),
		NetworkMode: aws.String("awsvpc"),
		Cpu:         aws.String(cpu),
		Memory:      aws.String(memory),
	}

	Info("Registering Task Definition")
	result, err := svc.RegisterTaskDefinition(input)
	CheckIfError(err)

	taskDefinition = *result.TaskDefinition.TaskDefinitionArn

	Info("Result: %s", result)
	return
}
