package main

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/codegangsta/cli"
	"github.com/lgtmco/lgtm-go/lgtm"
)

var GetCmd = cli.Command{
	Name:  "get",
	Usage: "get maintainer detail",
	Action: func(c *cli.Context) {
		handle(c, getCmd)
	},
}

func getCmd(c *cli.Context, client lgtm.Client) error {
	var (
		repo = c.Args().Get(0)
		org  = c.Args().Get(1)
		err  error

		maintainer *lgtm.Maintainer
	)

	if len(org) == 0 {
		maintainer, err = client.Maintainer(repo)
	} else {
		maintainer, err = client.MaintainerOrg(repo, org)
	}
	if err != nil {
		return err
	}
	return toml.NewEncoder(os.Stdout).Encode(maintainer)
}
