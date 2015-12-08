package main

import (
	"fmt"
	"path"

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
		cli.StringFlag{
			Name:  "exclude",
			Usage: "exclude repositories matching the pattern",
		},
	},
}

func listCmd(c *cli.Context, client lgtm.Client) error {
	var (
		active   = c.BoolT("active")
		inactive = c.Bool("inactive")
		exclude  = c.String("exclude")
	)
	if inactive {
		active = false
	}
	repos, err := client.Repos()
	if err != nil {
		return err
	}
	for _, repo := range repos {
		match, _ := path.Match(exclude, repo.Slug)
		if len(exclude) != 0 && !match {
			continue
		}

		switch {
		case active && repo.ID != 0:
			fmt.Println(repo.Slug)
		case inactive && repo.ID == 0:
			fmt.Println(repo.Slug)
		}
	}
	return nil
}
