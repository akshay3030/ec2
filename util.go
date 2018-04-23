package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/urfave/cli"
)

var ec2Svc *ec2.EC2

func setregion(c *cli.Context) error {
	if c.String("region") != "" {
		var sess = session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
			Config:            aws.Config{Region: aws.String(c.String("region"))},
		}))
		ec2Svc = ec2.New(sess)
	} else {
		var sess = session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
		ec2Svc = ec2.New(sess)
	}
	return nil
}

func trimQuotes(s string) string {
	if len(s) >= 2 {
		if s[0] == '"' && s[len(s)-1] == '"' {
			return s[1 : len(s)-1]
		}
	}
	return s
}

func confirmation() bool {
	var answer string
	fmt.Scanf("%s", &answer)
	if answer != "yes" {
		return false
	}
	return true
}
