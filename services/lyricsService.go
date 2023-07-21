package spotifyServices

import (
	"encoding/json"
	"log"
	spotifyRequest "spotifyApp/httpRequest"
	"spotifyApp/schemas"
)

var lyricsApiRoute string = "https://spotify-lyric-api.herokuapp.com/?trackid="

func GetLyricsForTrack(track *schemas.TrackInfo) (*schemas.Lyrics, error) {
	lyrics := &schemas.Lyrics{}
	url := lyricsApiRoute + track.Id

	body, err := spotifyRequest.MakeHttpRequest("GET", url, nil, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = json.Unmarshal(body, &lyrics)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if track.Artist != nil && len(track.Artist) > 0 {
		lyrics.Artist = track.Artist[0].Name
	}
	lyrics.Title = track.Name
	return lyrics, nil
}
