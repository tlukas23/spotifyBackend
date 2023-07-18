package schemas

type TracksQuery struct {
	Tracks []TrackInfo `json:"tracks"`
}

type TrackInfo struct {
	Album       *AlbumInfo `json:"album"`
	Id          string     `json:"id"`
	Name        string     `json:"name"`
	Popularity  int        `json:"popularity"`
	TrackNumber int        `json:"track_number"`
}

type AlbumInfo struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	ReleaseDate string `json:"release_date"`
	TotalTracks int    `json:"total_tracks"`
}
