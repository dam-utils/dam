package servers

import (
	"dam/driver/conf/option"
	"sort"
	"strings"

	"dam/driver/validate"
)

type label struct {
	repos    map[string]bool
	official bool
}

func NewLabel(s string) *label {
	storage := new(label)
	storage.repos = make(map[string]bool)

	if s == "" {
		return storage
	}

	str := strings.Split(s, option.Config.DefaultRepo.GetLabelReposSeparator())
	for _, repo := range str {
		storage.AddRepo(repo)
	}

	return storage
}

func (l *label) AddRepo(repo string) {
	switch repo {
	case "":
		l.official = true
	default:
		l.repos[repo] = true
	}
}

func (l *label) ValidateRepos() error {
	for repo := range l.repos {
		err := validate.CheckServer(repo)
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *label) ReposList() ([]string, bool) {
	return map2slice(l.repos), l.official
}

func (l *label) String() string {
	var result strings.Builder

	repos := map2slice(l.repos)

	for i, repo := range repos {
		result.WriteString(repo)
		if i != len(repos) - 1 {
			result.WriteString(option.Config.DefaultRepo.GetLabelReposSeparator())
		}
	}

	if l.official {
		result.WriteString(option.Config.DefaultRepo.GetLabelReposSeparator())
	}

	return result.String()
}

func map2slice(m map[string]bool) []string {
	result := make([]string, 0)

	for key := range m {
		result = append(result, key)
	}

	sort.Strings(result)

	return result
}