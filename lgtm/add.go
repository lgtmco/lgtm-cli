package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/lgtmco/lgtm-go/lgtm"
)

var AddCmd = cli.Command{
	Name:  "add",
	Usage: "activate repository approvals",
	Action: func(c *cli.Context) {
		handle(c, addCmd)
	},
}

func addCmd(c *cli.Context, client lgtm.Client) error {
	repo, err := client.Activate(c.Args().Get(0))
	if err != nil {
		return err
	}
	fmt.Println("successfully activated %s", repo.Slug)
	return nil
}
