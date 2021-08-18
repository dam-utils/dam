package registry_official

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"dam/config"
	"dam/driver/logger"
)

func GetAppVersions(app string) *[]string {
	url := config.OFFICIAL_REGISTRY_URL+"/v2/"+app+"/tags/list"

	tr := &http.Transport{
		MaxIdleConns:    config.SEARCH_MAX_CONNECTIONS,
		IdleConnTimeout: time.Duration(config.SEARCH_TIMEOUT_MS) * time.Millisecond,
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println()
		logger.Fatal(err.Error())
	}
	token := GetBearerToken(app)
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		logger.Fatal("Cannot send request to URL '%s'", url)
	}
	defer resp.Body.Close()

	type AppVersionsResponse struct {
		Tags []string `json:"tags"`
	}
	var body AppVersionsResponse
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		logger.Fatal("Cannot parse app versions in the body from URL '%s' with error: %s", url)
	}
	vers := body.Tags
	return &vers
}

