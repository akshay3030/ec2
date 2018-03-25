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
				list()
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
					if len(inst) > 1 {
						fmt.Println("Start all these instances?")
						for _, v := range inst {
							fmt.Println(trimQuotes(v))
						}
						fmt.Println("Type yes to confirm")
						if confirmation() {
							for _, v := range inst {
								start(trimQuotes(v))
							}
						}
					} else {
						start(trimQuotes(inst[0]))
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
					if len(inst) > 1 {
						fmt.Println("Stop all these instances?")
						for _, v := range inst {
							fmt.Println(trimQuotes(v))
						}
						fmt.Println("Type yes to confirm")
						if confirmation() {
							for _, v := range inst {
								stop(trimQuotes(v))
							}
						}
					} else {
						stop(trimQuotes(inst[0]))
					}
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}
