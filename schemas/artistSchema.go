package schemas

type QueryArtistRsp struct {
	ArtistRsp *ArtistQuery `json:"artists"`
}

type ArtistQuery struct {
	ArtistList []ArtistInfo `json:"items"`
}
type ArtistInfo struct {
	SpotifyID   string             `json:"id"`
	ExternalURL *ArtistExternalURL `json:"external_urls"`
	Genres      []string           `json:"genres"`
	Name        string             `json:"name"`
	Followers   *ArtistFollowers   `json:"followers"`
	Href        string             `json:"href"`
	Images      []ArtistImages     `json:"images"`
	Popularity  int64              `json:"popularity"`
	Uri         string             `json:"uri"`
}

type ArtistFollowers struct {
	Href  string `json:"href"`
	Total int64  `json:"total"`
}

type ArtistImages struct {
	Height int32  `json:"height"`
	Url    string `json:"url"`
	Width  int32  `json:"width"`
}

type ArtistExternalURL struct {
	Spotify string `json:"spotify"`
}
