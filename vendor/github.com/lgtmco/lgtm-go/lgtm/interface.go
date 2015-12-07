package lgtm

// Client to access the LGTM remote APIs.
type Client interface {
	User() (*User, error)
	Token(string) (*Token, error)
	Repo(string) (*Repo, error)
	Repos() ([]*Repo, error)
	Activate(string) (*Repo, error)
	Deactivate(string) error
	Maintainer(string) (*Maintainer, error)
	MaintainerOrg(string, string) (*Maintainer, error)
}
