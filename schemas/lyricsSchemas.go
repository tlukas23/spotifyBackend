package schemas

type Lyrics struct {
	Title  string      `json:"title"`
	Artist string      `json:"artist"`
	Lines  []LyricLine `json:"lines"`
}

type LyricLine struct {
	Words string `json:"words"`
}
