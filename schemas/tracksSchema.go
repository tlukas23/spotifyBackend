package schemas

type TracksQueryFromArtist struct {
	Tracks []TrackInfo `json:"tracks"`
}

type TracksQuery struct {
	TrackItems *TrackItems `json:"tracks"`
}

type TrackItems struct {
	Tracks []TrackInfo `json:"items"`
}

type TrackInfo struct {
	Album       *AlbumInfo   `json:"album"`
	Artist      []ArtistInfo `json:"artists"`
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	Popularity  int          `json:"popularity"`
	TrackNumber int          `json:"track_number"`
}

type AlbumInfo struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	ReleaseDate string `json:"release_date"`
	TotalTracks int    `json:"total_tracks"`
}
