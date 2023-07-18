package spotifyUsers

import (
	"encoding/json"
	spotifyRequest "spotifyApp/httpRequest"
	"spotifyApp/schemas"
)

func LoginDeveloper() (*schemas.AccountAPIData, error) {
	devAccount := &schemas.AccountAPIData{}
	postBody := make(map[string]string, 0)
	postBody["grant_type"] = "client_credentials"
	postBody["client_id"] = "c3e3f0a2435b485f87d962bf812228f1"
	postBody["client_secret"] = "54f1386b974744c38668c4c1c0a3fda0"
	headers := make(map[string]string, 0)
	headers["Content-Type"] = "application/x-www-form-urlencoded"

	body, err := spotifyRequest.MakeHttpRequest("POST", "https://accounts.spotify.com/api/token", postBody, headers)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &devAccount)
	if err != nil {
		return nil, err
	}
	return devAccount, nil
}
