package api

import (
	"fmt"
	"log"
	"net/http"
)

func InitServer() {
	http.HandleFunc("/search", searchDB)
	http.HandleFunc("/getItem", getItem)
	http.HandleFunc("/getEpisodes", getSeriesEpisodes)
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "0\n")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
