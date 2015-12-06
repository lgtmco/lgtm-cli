package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
	"github.com/lgtmco/lgtm-go/lgtm"
	"golang.org/x/oauth2"
)

var PushCmd = cli.Command{
	Name:  "push",
	Usage: "push the maintainer file to GitHub",
	Action: func(c *cli.Context) {
		handle(c, pushCmd)
	},
}

func pushCmd(c *cli.Context, _ lgtm.Client) error {
	var (
		src  io.ReadCloser
		repo string = c.Args().Get(0)
		path string = c.Args().Get(1)
		err  error
	)

	if len(path) != 0 {
		src, err = os.Open(path)
		if err != nil {
			return err
		}
		defer src.Close()
	} else {
		src = os.Stdin
	}

	var dest bytes.Buffer
	io.Copy(&dest, src)
	if err != nil {
		return err
	}
	v := new(lgtm.Maintainer)
	_, err = toml.Decode(dest.String(), v)
	if err != nil {
		return err
	}

	//server := c.GlobalString("github-server")
	token := c.GlobalString("github-token")
	client := setupClient(token)
	sha, err := createUpdateFile(
		client,
		strings.Split(repo, "/")[0],
		strings.Split(repo, "/")[1],
		dest.Bytes(),
	)
	if err == nil {
		fmt.Printf("pushed MAINTAINERS file to %s@%s\n", repo, sha[:7])
	}
	return err
}

func setupClient(accessToken string) *github.Client {
	token := oauth2.Token{AccessToken: accessToken}
	source := oauth2.StaticTokenSource(&token)
	client := oauth2.NewClient(oauth2.NoContext, source)
	return github.NewClient(client)
}

func createUpdateFile(client *github.Client, owner, name string, data []byte) (string, error) {
	sha, err := getSha(client, owner, name)
	if err != nil {
		return createFile(client, owner, name, data)
	} else {
		return updateFile(client, owner, name, sha, data)
	}
}

// helper function to get the MAINTAINER files SHA
func getSha(client *github.Client, owner, name string) (string, error) {
	opt := new(github.RepositoryContentGetOptions)
	res, _, _, err := client.Repositories.GetContents(owner, name, "MAINTAINERS", opt)
	if err != nil {
		return "", err
	}
	return *res.SHA, nil
}

// helper function to update the MAINTAINER file
func updateFile(client *github.Client, owner, name, sha string, data []byte) (string, error) {
	res, _, err := client.Repositories.UpdateFile(
		owner,
		name,
		"MAINTAINERS",
		&github.RepositoryContentFileOptions{
			Content: data,
			Message: github.String("Updated MAINTAINERS file [CI SKIP]"),
			SHA:     github.String(sha),
		},
	)
	if err != nil {
		return "", fmt.Errorf("Error updating MAINTAINERS file at %s. %s", sha, err)
	}
	return *res.SHA, nil
}

// helper function to update the MAINTAINER file
func createFile(client *github.Client, owner, name string, data []byte) (string, error) {
	res, _, err := client.Repositories.CreateFile(
		owner,
		name,
		"MAINTAINERS",
		&github.RepositoryContentFileOptions{
			Content: data,
			Message: github.String("Created MAINTAINERS file [CI SKIP]"),
		},
	)
	if err != nil {
		return "", fmt.Errorf("Error creating MAINTAINERS file. %s", err)
	}
	return *res.SHA, nil
}
