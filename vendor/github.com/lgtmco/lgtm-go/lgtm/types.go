package lgtm

// Person represets an individual in the MAINTAINERS file.
type Person struct {
	Name  string `json:"name"  toml:"name"`
	Email string `json:"email" toml:"email"`
	Login string `json:"login" toml:"login"`
}

// Org represents a group, team or subset of users.
type Org struct {
	People []string `json:"people" toml:"people"`
}

// Maintainer represents a MAINTAINERS file.
type Maintainer struct {
	People    map[string]*Person `json:"people"    toml:"people"`
	Org       map[string]*Org    `json:"org"       toml:"org"`
	Approvals int                `json:"approvals" toml:"approvals"`
}

// User represents a github user.
type User struct {
	ID     int64  `json:"id"`
	Login  string `json:"login"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}

// Repo represents a github repository.
type Repo struct {
	ID      int64  `json:"id,omitempty"`
	Owner   string `json:"owner"`
	Name    string `json:"name"`
	Slug    string `json:"slug"`
	Private bool   `json:"private"`
	Link    string `json:"link_url"`
}

// Token represents an Access Token used to authenticate
// to the remote server.
type Token struct {
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
	Expires int64  `json:"expires_in"`
}
