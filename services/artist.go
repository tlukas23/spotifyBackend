package spotifyServices

import (
	"encoding/json"
	spotifyRequest "spotifyApp/httpRequest"
	"spotifyApp/schemas"
	"spotifyApp/utils"
)

func FindArtistIdByName(name string, bearerToken string) (*schemas.ArtistInfo, error) {
	artistRsp := &schemas.QueryArtistRsp{}
	params := make(map[string]string, 0)
	params["q"] = name
	params["type"] = "artist"
	params["limit"] = "5"
	headers := make(map[string]string, 0)
	headers["Authorization"] = "Bearer " + bearerToken

	body, err := spotifyRequest.MakeHttpQuery("GET", "https://api.spotify.com/v1/search",
		params, headers)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &artistRsp)
	if err != nil {
		return nil, err
	}

	return utils.GetMostPopularArtist(artistRsp.ArtistRsp), nil
}
