package main

import "github.com/urfave/cli"

var options = []cli.Flag{
	&cli.StringFlag{
		Name:   "username, u",
		EnvVar: "USER",
		Usage:  "Username",
	},
}
