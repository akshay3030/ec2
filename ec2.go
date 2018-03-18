package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "luizc2"
	app.Usage = "Easy way to stop/start/backup AWS instances from cli"
	app.Commands = []cli.Command{
		{
			Name:    "list",
			Usage:   "list instances",
			Aliases: []string{"ls"},
			Action: func(c *cli.Context) error {
				ec2list := list()
				for i, v := range ec2list {
					fmt.Printf("%d - %s\n", i, trimQuotes(v))
				}
				return nil
			},
		},
		{
			Name:  "start",
			Usage: "start instance",
			Action: func(c *cli.Context) error {
				if strings.HasPrefix(os.Args[2], "i-") {
					start(os.Args[2])
				} else {
					inst := listByName(os.Args[2])
					for _, v := range inst {
						start(trimQuotes(v))
					}
				}

				return nil
			},
		},
		{
			Name:  "stop",
			Usage: "stopt instance",
			Action: func(c *cli.Context) error {
				if strings.HasPrefix(os.Args[2], "i-") {
					stop(os.Args[2])
				} else {
					inst := listByName(os.Args[2])
					for _, v := range inst {
						stop(trimQuotes(v))
					}
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}
