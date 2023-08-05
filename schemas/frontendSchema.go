package schemas

type SearchArtistInput struct {
	Name string `json:"artist_name"`
}

type SearchTrackInput struct {
	TrackName  string `json:"track_name"`
	ArtistName string `json:"artist_name"`
}

type UserInfo struct {
	AuthToken string `json:"auth_token"`
}

// ****************************************************//
