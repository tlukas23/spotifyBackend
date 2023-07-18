package spotifyServices

import (
	"encoding/json"
	spotifyRequest "spotifyApp/httpRequest"
	"spotifyApp/schemas"
)

func GetPopularTracksFromArtist(artistId string, bearerToken string) (*schemas.TracksQuery, error) {
	tracks := &schemas.TracksQuery{}
	url := "https://api.spotify.com/v1/artists/" + artistId + "/top-tracks?market=US"

	headers := make(map[string]string, 0)
	headers["Authorization"] = "Bearer " + bearerToken
	body, err := spotifyRequest.MakeHttpRequest("GET", url, nil, headers)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &tracks)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}
