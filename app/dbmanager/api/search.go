package api

import (
	"encoding/json"
	"media/mediaserver/dbmanager/postgre"
	"media/mediaserver/types"
	"net/http"
)

func searchDB(w http.ResponseWriter, r *http.Request) {
	var request types.SearchRequest
	json.NewDecoder(r.Body).Decode(&request)

	response := types.ContentResponse{}
	content, err := postgre.SearchContent(request.Genre, request.AgeRating, request.ContentType, request.SearchString, request.SortByYear)
	if err != nil {
		response.Error = err.Error()
	} else {
		response.Content = content
	}
	json.NewEncoder(w).Encode(response)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	var request types.ContentRequest
	json.NewDecoder(r.Body).Decode(&request)

	response := types.ContentResponse{}
	content, err := postgre.GetSingleContent(request.ID)
	if err != nil {
		response.Error = err.Error()
	} else {
		response.Content = content
	}

	json.NewEncoder(w).Encode(response)
}

func getSeriesEpisodes(w http.ResponseWriter, r *http.Request) {
	var request types.ContentRequest
	json.NewDecoder(r.Body).Decode(&request)

	response := types.ContentResponse{}
	content, err := postgre.GetEpisodes(request.ID)
	if err != nil {
		response.Error = err.Error()
	} else {
		response.Content = content
	}

	json.NewEncoder(w).Encode(response)
}
