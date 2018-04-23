package main

import (
	"os"

	"github.com/urfave/cli"
)

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
	}

	app.Before = setregion
	app.Run(os.Args)

}
