package postgre

import (
	"database/sql"
	"media/mediaserver/types"

	"github.com/lib/pq"
)

const (
	episodeInsert = "INSERT into episodes VALUES( $1,$2,$3,$4,$5,$6, $7)"
	contentInsert = "INSERT into content VALUES( $1,$2,$3,$4,$5,$6,$7,$8,$9,$10)"

	singleEpisodeSelect = "SELECT * FROM episodes WHERE id = $1"
	singleContentSelect = "SELECT * FROM content WHERE id = $1"

	parentEpisodeSelect = "SELECT * FROM episodes WHERE parent = $1"

	baseSearchQuery = "SELECT * FROM content WHERE array_to_string(genres,';','empty')  LIKE '%' || $1 || '%' and COALESCE(ageRating,'empty') LIKE '%' || $2 || '%' and COALESCE(contentType,'') LIKE '%' || $3 || '%' and array_to_string(genres,';') || ';' || description || ';' || title || ';' || array_to_string(stars,';') || ';' || array_to_string(directors,';') like '%' || $4 || '%'"

	sortYearAscending = "ORDER BY year ASC"
)

func InsertContentIntoDB(content *types.ContentItem) (err error) {
	populateArrayFields(content)

	_, err = pool.Exec(contentInsert, content.ID, content.Title, content.Description, content.ContentType, pq.Array(content.Genres), content.AgeRating, content.DurationInMin, content.Year, pq.Array(content.Directors), pq.Array(content.Stars))
	return
}

func InsertEpisodeIntoDB(content *types.ContentItem) (err error) {
	_, err = pool.Exec(episodeInsert, content.ID, content.Title, content.Description, content.ContentType, content.Parent, content.EpisodeNumber, content.DurationInMin)

	return
}

func GetSingleContent(ID string) ([]types.ContentItem, error) {
	content := types.ContentItem{}

	err := pool.QueryRow(singleContentSelect, ID).Scan(&content.ID, &content.Title, &content.Description, &content.ContentType, pq.Array(&content.Genres), &content.AgeRating, &content.DurationInMin, &content.Year, pq.Array(&content.Directors), pq.Array(&content.Stars))
	if err != sql.ErrNoRows { //success or db error
		return []types.ContentItem{content}, err
	}

	err = pool.QueryRow(singleEpisodeSelect, ID).Scan(&content.ID, &content.Title, &content.Description, &content.ContentType, &content.Parent, &content.EpisodeNumber, &content.DurationInMin)

	return []types.ContentItem{content}, err
}

func GetEpisodes(parentID string) ([]types.ContentItem, error) {

	episodes := make([]types.ContentItem, 0, 10)

	rows, err := pool.Query(parentEpisodeSelect, parentID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		episode := types.ContentItem{}
		_ = rows.Scan(&episode.ID, &episode.Title, &episode.Description, &episode.ContentType, &episode.Parent, &episode.EpisodeNumber, &episode.DurationInMin)

		episodes = append(episodes, episode)
	}

	return episodes, nil
}

func SearchContent(genre, ageRating, contentType, searchString string, sortByYearAscending bool) ([]types.ContentItem, error) {
	query := baseSearchQuery
	if sortByYearAscending {
		query += sortYearAscending
	}

	episodes := make([]types.ContentItem, 0, 10)

	rows, err := pool.Query(query, genre, ageRating, contentType, searchString)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		content := types.ContentItem{}
		_ = rows.Scan(&content.ID, &content.Title, &content.Description, &content.ContentType, pq.Array(&content.Genres), &content.AgeRating, &content.DurationInMin, &content.Year, pq.Array(&content.Directors), pq.Array(&content.Stars))

		cleanDefaultArray(&content)

		episodes = append(episodes, content)
	}

	return episodes, nil
}

// not sure of the SQL to deal with empty arrays so just defaulting empty values and removing them later
func populateArrayFields(contentItem *types.ContentItem) {
	if len(contentItem.Genres) == 0 {
		contentItem.Genres = []string{""}
	}
	if len(contentItem.Directors) == 0 {
		contentItem.Directors = []string{""}
	}
	if len(contentItem.Stars) == 0 {
		contentItem.Stars = []string{""}
	}
}

func cleanDefaultArray(contentItem *types.ContentItem) {
	if len(contentItem.Genres) == 1 && contentItem.Genres[0] == "" {
		contentItem.Genres = nil
	}
	if len(contentItem.Directors) == 1 && contentItem.Directors[0] == "" {
		contentItem.Directors = nil
	}
	if len(contentItem.Stars) == 1 && contentItem.Stars[0] == "" {
		contentItem.Stars = nil
	}
}
