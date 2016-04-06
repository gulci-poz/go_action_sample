package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// plik będzie zamknięty kiedy funkcja cos zwróci lub nawet w wypadku nieoczekiwanego zakończenia
	defer file.Close()
	
	// slice zbudowany ze wskaźników do Feed
	var feeds [] *Feed
	err = json.NewDecoder(file).Decode($feeds)
	
	return feeds, err
}
