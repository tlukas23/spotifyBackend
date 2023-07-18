package utils

import (
	"spotifyApp/schemas"
)

func GetMostPopularArtist(artists *schemas.ArtistQuery) *schemas.ArtistInfo {
	mostPopularArtist := new(schemas.ArtistInfo)
	for x := range artists.ArtistList {
		if artists.ArtistList[x].Popularity > mostPopularArtist.Popularity {
			mostPopularArtist = &artists.ArtistList[x]
		}
	}
	return mostPopularArtist
}
