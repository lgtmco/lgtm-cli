package main

// lgtm ls                   DONE
// lgtm get <repo> <team>    DONE
// lgtm add <repo>           DONE
// lgtm rm  <repo>           DONE

// lgtm set <repo> <source>

// lgtm token <token>        DONE

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "lgtm"
	app.Version = "1.0.0"
	app.Usage = "lgtm command line utility"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "server",
			Value:  "https://lgtm.co",
			Usage:  "lgtm server address",
			EnvVar: "LGTM_SERVER",
		},
		cli.StringFlag{
			Name:   "token",
			Value:  "",
			Usage:  "lgtm server api token",
			EnvVar: "LGTM_TOKEN",
		},
		cli.StringFlag{
			Name:   "github-server",
			Value:  "http://api.github.com",
			Usage:  "github server address",
			EnvVar: "GITHUB_SERVER",
		},
		cli.StringFlag{
			Name:   "github-token",
			Value:  "",
			Usage:  "github server api token",
			EnvVar: "GITHUB_TOKEN",
		},
	}
	app.Commands = []cli.Command{
		TokenCmd,
		ListCmd,
		GetCmd,
		AddCmd,
		DelCmd,
		PushCmd,
	}

	app.Run(os.Args)
}
