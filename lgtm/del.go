package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/lgtmco/lgtm-go/lgtm"
)

var DelCmd = cli.Command{
	Name:  "rm",
	Usage: "de-activates repository approvals",
	Action: func(c *cli.Context) {
		handle(c, delCmd)
	},
}

func delCmd(c *cli.Context, client lgtm.Client) error {
	repo := c.Args().Get(0)
	if len(repo) == 0 {
		return fmt.Errorf("please provide the repository name")
	}
	err := client.Deactivate(repo)
	if err != nil {
		return fmt.Errorf("error de-activating repository. %s", err)
	}
	fmt.Printf("successfully de-activated %s\n", repo)
	return nil
}
