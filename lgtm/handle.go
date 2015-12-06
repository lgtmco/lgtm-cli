package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/lgtmco/lgtm-go/lgtm"
)

type handlerFunc func(*cli.Context, lgtm.Client) error

// handle wraps the command function handlers and
// sets up the environment.
func handle(c *cli.Context, fn handlerFunc) {
	token := c.GlobalString("token")
	server := c.GlobalString("server")

	// the token is required for any API calls.
	if len(token) == 0 {
		fmt.Println("Error: you must provide your lgtm access token.")
		os.Exit(1)
	}

	// create the drone client
	client := lgtm.NewClientToken(server, token)

	// handle the function
	if err := fn(c, client); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
