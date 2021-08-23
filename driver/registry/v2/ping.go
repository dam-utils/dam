package registry_v2

import (
	"net/http"

	"dam/driver/conf/option"
	"dam/driver/structures"
)

var SessionURL string

func CheckRepo(repo *structures.Repo, protocol string) error {
	tr := &http.Transport{
		MaxIdleConns:    option.Config.Search.GetMaxConnections(),
		IdleConnTimeout: option.Config.Search.GetTimeoutMs(),
	}
	client := &http.Client{Transport: tr}
	SessionURL = protocol + "://" + repo.Server + "/"

	req, err := http.NewRequest("GET", SessionURL, nil)
	if err != nil {
		return err
	}
	if repo.Username != "" {
		req.SetBasicAuth(repo.Username, repo.Password)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}
