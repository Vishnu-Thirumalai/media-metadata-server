package types

type responseMeta struct {
	Error string
}

type SearchRequest struct {
	Genre        string
	AgeRating    string
	ContentType  string
	SortByYear   bool
	SearchString string
}

type ContentRequest struct {
	ID string
}

type ContentResponse struct {
	responseMeta
	Content []ContentItem
}
