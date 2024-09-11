# Media Metadata Server

## What and Why
Parses a bunch of json data into a table and makes it queriable via a REST api

Take-home interview for a company that I put enough work into that I'd like to keep around.


## Requirements
* Docker

## To Start

1. cd to app
2. Run _docker compose build_
3. run _docker compose up_

# Endpoints:
All endpoints return a json struct witht he following:
```
{
    error string
    content []{
        ID            string
	    Title         string
	    Description   string
	    ContentType   string
        Parent        string
        Genres        []string
        EpisodeNumber string
        AgeRating     string  
        DurationInMin int     
        Year          int     
        Directors     []string
        Stars         []string
    }
}
```

1. http://localhost:8080/listItems - returns all non-episode items filtered with the given fields
    * Accepts a json with the following fields:
        * Genre        string 
        * AgeRating    string
        * ContentType  string
        * SearchString string
        * SortByYear   bool - if true, returns results ordered ascending by year
    * e.g. curl -s -XPOST -d'{"contentType":"video","searchString":"Jackman"}' http://localhost:8080/listItems
2. http://localhost:8080/getEpisodesForSeries - returns all episodes whose parent series has the given ID
    * Accepts a json with the following field:
        * ID string
    * e.g. curl -s -XPOST -d'{"id":"podcast_1"}' http://localhost:8080/getEpisodesForSeries
3. http://localhost:8080/contentInfo - returns information of content item with the given ID
    * Accepts a json with the following field:
        * ID string
    *  e.g. curl -s -XPOST -d'{"id":"movie_2"}' http://localhost:8080/contentInfo