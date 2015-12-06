package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/lgtmco/lgtm-go/lgtm"
)

var ListCmd = cli.Command{
	Name:  "ls",
	Usage: "list repositories",
	Action: func(c *cli.Context) {
		handle(c, listCmd)
	},
	Flags: []cli.Flag{
		cli.BoolTFlag{
			Name:  "active",
			Usage: "list active repositories",
		},
		cli.BoolFlag{
			Name:  "inactive",
			Usage: "list inactive repositories",
		},
	},
}

func listCmd(c *cli.Context, client lgtm.Client) error {
	var (
		active   = c.BoolT("active")
		inactive = c.Bool("inactive")
	)
	repos, err := client.Repos()
	if err != nil {
		return err
	}
	for _, repo := range repos {
		switch {
		case active && repo.ID != 0:
			fmt.Println(repo.Slug)
		case inactive && repo.ID == 0:
			fmt.Println(repo.Slug)
		}
	}
	return nil
}
