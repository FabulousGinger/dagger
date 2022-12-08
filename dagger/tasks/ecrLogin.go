package tasks

import (
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

func ECRLogin() (err error) {
	awsRegion := os.Getenv("AWS_DEFAULT_REGION")

	sess := session.Must(session.NewSession())

	svc := ecr.New(sess, aws.NewConfig().WithRegion(awsRegion))

	input := &ecr.GetAuthorizationTokenInput{}
	resp, err := svc.GetAuthorizationToken(input)
	auth := resp.AuthorizationData

	decode, err := base64.StdEncoding.DecodeString(*auth[0].AuthorizationToken)
	token := strings.SplitN(string(decode), ":", 2)
	user := token[0]
	password := token[1]
	endpoint := *auth[0].ProxyEndpoint

	cmd := fmt.Sprintf("docker login -u %s -p %s %s", user, password, endpoint)
	login := exec.Command("bash", "-c", cmd)
	err = login.Run()

	Info("Login Succeeded.\n\nThe registry endpoint is:\n%s",
		endpoint)

	return
}
