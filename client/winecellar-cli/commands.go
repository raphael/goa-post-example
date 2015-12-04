package main

import (
	"net/http"

	"github.com/raphael/goa-post-example/client"
	"gopkg.in/alecthomas/kingpin.v2"
)

type (
	// ShowBottleCommand is the command line data structure for the show action of Bottle
	ShowBottleCommand struct {
		// Path is the HTTP request path.
		Path string
	}
)

// Run makes the HTTP request corresponding to the ShowBottleCommand command.
func (cmd *ShowBottleCommand) Run(c *client.Client) (*http.Response, error) {
	return c.ShowBottle(cmd.Path)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowBottleCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, format is /cellar/bottles/:bottleID`).Required().StringVar(&cmd.Path)
}
