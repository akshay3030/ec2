package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/urfave/cli"
)

var ec2Svc *ec2.EC2

func createSession(c *cli.Context) error {
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

func main() {

	app := cli.NewApp()
	app.Name = "luizc2"
	app.Usage = "Easy way to stop/start/backup AWS instances from cli"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "region", Usage: "set region or use all"},
	}
	app.Commands = []cli.Command{
		{
			Name:    "list",
			Usage:   "list instances",
			Aliases: []string{"ls"},
			Action: func(c *cli.Context) error {
				list()
				return nil
			},
		},
		{
			Name:  "start",
			Usage: "start instance",
			Action: func(c *cli.Context) error {
				start(os.Args)

				return nil
			},
		},
		{
			Name:  "stop",
			Usage: "stopt instance",
			Action: func(c *cli.Context) error {
				stop(os.Args)

				return nil
			},
		},
		{
			Name:    "IP",
			Usage:   "IP instance",
			Aliases: []string{"ip", "hostname"},
			Action: func(c *cli.Context) error {
				getIP(os.Args)

				return nil
			},
		},
		{
			Name:    "backup",
			Usage:   "backup instance",
			Aliases: []string{"bck"},
			Action: func(c *cli.Context) error {
				backup(os.Args)

				return nil
			},
		},
	}

	app.Before = createSession
	app.Run(os.Args)

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
