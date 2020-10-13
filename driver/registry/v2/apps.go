package registry_v2

import (
	"dam/driver/structures"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"dam/config"
	"dam/driver/logger"
)

type ResponseGetAppNames struct {
	Repositories []string `json:"repositories"`
}

func GetAppNames(repo *structures.Repo) *[]string {
	tr := &http.Transport{
		MaxIdleConns:    config.SEARCH_MAX_CONNECTS,
		IdleConnTimeout: time.Duration(config.SEARCH_TIMEOUT_MS) * time.Millisecond,
	}
	client := &http.Client{Transport: tr}

	url := SessionURL + "v2/_catalog?n="+strconv.Itoa(config.INTERNAL_REPO_SEARCH_APPS_LIMIT)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
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

	var body ResponseGetAppNames
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		logger.Fatal("Cannot parse response from URL '%s' with error: %s", url, err)
	}
	return &body.Repositories
}



