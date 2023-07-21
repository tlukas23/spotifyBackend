package spotifyServices

import (
	"encoding/json"
	"errors"
	"log"
	spotifyRequest "spotifyApp/httpRequest"
	"spotifyApp/schemas"
)

func GetTrackInfoFromName(trackInput *schemas.SearchTrackInput, bearerToken string) (*schemas.TrackInfo, error) {
	tracks := &schemas.TracksQuery{}
	params := make(map[string]string, 0)
	params["q"] = trackInput.TrackName + " " + trackInput.ArtistName
	params["type"] = "track"
	params["limit"] = "1"
	headers := make(map[string]string, 0)
	headers["Authorization"] = "Bearer " + "BQDk7t0cWW2lwUoDAItE_ml2_D0A7n3Pk-ENx1cNLfE33R7D7XeX9ihzN68f-Eqf7E3DVX-_10fmr-R5jGS8KjVEMHLfWr257nOw_madAmPrWNJifNqPZ9VQathUDKZFtll5WsjwO3NAk65cplpa6RhNOb1FaFREO_ucj7-pkvgrnXoyFOYZd52UPZ4Z8zZF"

	body, err := spotifyRequest.MakeHttpQuery("GET", "https://api.spotify.com/v1/search",
		params, headers)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = json.Unmarshal(body, &tracks)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if tracks.TrackItems == nil || len(tracks.TrackItems.Tracks) == 0 {
		log.Println("no track found")
		return nil, errors.New("no track found")
	}

	return &tracks.TrackItems.Tracks[0], nil
}

func GetTopTrackInfoForUser(input *schemas.LoginUser, bearerToken string) (*schemas.TrackItems, error) {
	tracks := &schemas.TrackItems{}
	params := make(map[string]string, 0)
	params["time_range"] = "medium_term"
	params["limit"] = "10"
	headers := make(map[string]string, 0)
	headers["Authorization"] = "Bearer " + input.AuthToken

	body, err := spotifyRequest.MakeHttpQuery("GET", "https://api.spotify.com/v1/me/top/tracks",
		params, headers)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = json.Unmarshal(body, &tracks)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if len(tracks.Tracks) == 0 {
		log.Println("no track found")
		return nil, errors.New("no track found")
	}

	return tracks, nil
}
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
