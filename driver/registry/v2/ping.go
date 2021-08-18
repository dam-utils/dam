package registry_v2

import (
	"dam/driver/structures"
	"net/http"
	"time"

	"dam/config"
)

var SessionURL string

func CheckRepo(repo *structures.Repo, protocol string) error {
	tr := &http.Transport{
		MaxIdleConns:    config.SEARCH_MAX_CONNECTIONS,
		IdleConnTimeout: time.Duration(config.SEARCH_TIMEOUT_MS) * time.Millisecond,
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
