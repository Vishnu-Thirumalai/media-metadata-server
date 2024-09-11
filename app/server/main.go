package main

import (
	"fmt"
	"log"
	"net/http"
)

func searchAllContent(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Post("http://dbmanager:8080/search", "application/json", r.Body)
	if err != nil {
		panic(err)
	}

	body := make([]byte, resp.ContentLength)
	resp.Body.Read(body)
	w.Write(body)
}

func getSeriesEpisodes(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Post("http://dbmanager:8080/getEpisodes", "application/json", r.Body)
	if err != nil {
		panic(err)
	}

	body := make([]byte, resp.ContentLength)
	resp.Body.Read(body)
	w.Write(body)
}

func getContentInfo(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Post("http://dbmanager:8080/getItem", "application/json", r.Body)
	if err != nil {
		panic(err)
	}

	body := make([]byte, resp.ContentLength)
	resp.Body.Read(body)
	w.Write(body)
}

func main() {
	http.HandleFunc("/listItems", searchAllContent)
	http.HandleFunc("/getEpisodesForSeries", getSeriesEpisodes)
	http.HandleFunc("/contentInfo", getContentInfo)
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "0\n")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
