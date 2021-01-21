package main

import (
	"io"
	"net/http"

	"github.com/urfave/cli"
)

var commandEntry = cli.Command{
	Name:   "entry",
	Usage:  "program entry",
	Action: entry,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "port",
			Value: ":8888",
		},
	},
}

func entry(c *cli.Context) error {

	port := c.String("port")

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	})

	return http.ListenAndServe(port, nil)

}
