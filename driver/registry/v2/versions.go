package registry_v2

import (
	"dam/driver/structures"
	"encoding/json"
	"net/http"
	"time"

	"dam/config"
	"dam/driver/logger"
)

func GetAppVersions(repo *structures.Repo, appName string) *[]string {
	tr := &http.Transport{
		MaxIdleConns:    config.SEARCH_MAX_CONNECTIONS,
		IdleConnTimeout: time.Duration(config.SEARCH_TIMEOUT_MS) * time.Millisecond,
	}
	client := &http.Client{Transport: tr}

	url := SessionURL + "v2/" + appName + "/tags/list"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil || req == nil {
		logger.Fatal("Cannot create new request for get URL '%s' with error: %s", url, err)
	}
	if repo.Username != "" {
		req.SetBasicAuth(repo.Username, repo.Password)
	}

	resp, err := client.Do(req)
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot get response from URL '%s' with error: %s", url, err)
	}

	type AppVersionsResponse struct {
		Tags []string `json:"tags"`
	}
	var body AppVersionsResponse

	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		logger.Fatal("Cannot parse app versions in the body from URL '%s' with error: %s", url, err)
	}
	vers := body.Tags
	return &vers
}