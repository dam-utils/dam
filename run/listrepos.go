package run

import (
	"dam/driver/decorate"
)

type ListReposSettings struct {
	Raw    bool
}

var ListReposFlags = new(ListReposSettings)

func ListRepos() {
	if ListReposFlags.Raw {
		decorate.PrintRAWReposList()
	} else {
		decorate.PrintReposList()
	}
}
