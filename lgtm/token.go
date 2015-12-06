package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/lgtmco/lgtm-go/lgtm"
)

var TokenCmd = cli.Command{
	Name:   "token",
	Usage:  "get an api token",
	Action: tokenCmd,
}

func tokenCmd(c *cli.Context) {
	server := c.GlobalString("server")
	github := c.GlobalString("github-token")
	if len(github) == 0 {
		github = c.Args().First()
	}
	client := lgtm.NewClient(server)
	token, err := client.Token(github)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(token.Access)
}
