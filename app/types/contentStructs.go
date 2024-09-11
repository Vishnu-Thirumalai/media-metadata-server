package types

type ContentMeta struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	ContentType string `json:"type,omitempty"`
}

type SeriesContentFields struct {
	Parent        string `json:"parent,omitempty"`
	EpisodeNumber string `json:"episodeNumber,omitempty"`
}

type ContentCommon struct {
	Genres        []string `json:"genres,omitempty"`
	AgeRating     string   `json:"ageRating,omitempty"`
	DurationInMin int      `json:"durationInMin,omitempty"`
	Year          int      `json:"year,omitempty"`
	Directors     []string `json:"directors,omitempty"`
	Stars         []string `json:"stars,omitempty"`
}

type ContentItem struct {
	ContentMeta
	ContentCommon
	SeriesContentFields
}
