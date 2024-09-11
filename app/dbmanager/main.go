package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"media/mediaserver/dbmanager/api"
	"media/mediaserver/dbmanager/postgre"
	"media/mediaserver/types"

	"os"
	"path/filepath"
)

func main() {
	setupDB()
	fmt.Println("Database initialized")
	api.InitServer()
}

func setupDB() {
	err := postgre.InitDB("/schema")
	if err != nil {
		panic(err.Error())
	}

	err = readFiles()
	if err != nil {
		panic("Data not mounted correctly: " + err.Error())
	}
}

func readFiles() error {
	data, err := os.ReadDir("/data")
	if err != nil {
		return err
	}

	for _, file := range data {
		if filepath.Ext(file.Name()) == ".json" {

			jsonData, err := os.ReadFile("/data/" + file.Name())
			if err != nil {
				return errors.Join(fmt.Errorf("/data/%s could not be read", file.Name()), err)
			}

			err = processItem(jsonData)
			if err != nil {
				return errors.Join(fmt.Errorf("/data/%s could not be inserted into db", file.Name()), err)
			}
		}
	}
	return nil
}

func processItem(jsonData []byte) error {

	var contentItem *types.ContentItem

	json.Unmarshal(jsonData, &contentItem)

	if contentItem.Parent != "" {
		return postgre.InsertEpisodeIntoDB(contentItem)
	} else {
		return postgre.InsertContentIntoDB(contentItem)
	}
}
