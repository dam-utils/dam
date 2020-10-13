package registry_official

import (
	"encoding/json"
	"net/http"

	"dam/config"
	"dam/driver/logger"
)

func GetBearerToken(app string) string {
	url := config.OFFICIAL_REGISTRY_AUTH_URL+ "&scope=repository:"+app+":pull"
	resp, err := http.Get(url)
	if err != nil {
		logger.Fatal("Cannot get token from URL '%s' with error: %s", url)
	}
	defer resp.Body.Close()

	type TokenResponse struct {
		AccessToken string `json:"access_token"`
	}
	var body TokenResponse

	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		logger.Fatal("Cannot parse token in the body from URL '%s' with error: %s", url, err)
	}
	return body.AccessToken
}

