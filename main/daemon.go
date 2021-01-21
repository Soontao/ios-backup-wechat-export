package main

import (
	"github.com/takama/daemon"
	"github.com/urfave/cli"
)

func createDaemonCommands(name, description string) ([]cli.Command, error) {
	m, err := daemon.New(name, description, daemon.SystemDaemon)
	if err != nil {
		return nil, err
	}
	rt := []cli.Command{
		{
			Name:  "install",
			Usage: "install service",
			Action: func(c *cli.Context) error {
				opt := c.StringSlice("args")
				_, err := m.Install(opt...)
				return err
			},
			Flags: []cli.Flag{
				cli.StringSliceFlag{
					Name:   "args, a",
					Usage:  "service running args",
					Value:  &cli.StringSlice{"entry"},
					EnvVar: "SERVICE_RUNNING_ARGS",
				},
			},
		},
		{
			Name:  "uninstall",
			Usage: "uninstall service",
			Action: func(c *cli.Context) error {
				_, err := m.Remove()
				return err
			},
		},
		{
			Name:  "start",
			Usage: "start service",
			Action: func(c *cli.Context) error {
				_, err := m.Start()
				return err
			},
		},
		{
			Name:  "stop",
			Usage: "stop service",
			Action: func(c *cli.Context) error {
				_, err := m.Stop()
				return err
			},
		},
		{
			Name:  "status",
			Usage: "output status of service",
			Action: func(c *cli.Context) error {
				_, err := m.Status()
				return err
			},
		},
	}

	return rt, nil
}
