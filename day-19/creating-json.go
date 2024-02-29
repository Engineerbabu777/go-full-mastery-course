package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Name  string `json:"name"`
	Actor string `json:"actor"`
	Year  int    `json:"year"`
}

func main() {
	EncodeJSON()
	DecodeJSON()
}

func EncodeJSON() {

	movies := []Movie{
		{"The Matrix 1", "Keanu Reeves 2", 20000},
		{"The Matrix 2", "Keanu Reeves", 1999},
	}

	jsonFormData, _ := json.MarshalIndent(movies, "", "\t")
	fmt.Printf("Data:%s ", jsonFormData)
}

func DecodeJSON() {
	jsonData := []byte(`{"name":"The Matrix 1","actor":"Keanu Reeves 2","year":20000}`)

	var movies Movie

	json.Unmarshal(jsonData, &movies)

	fmt.Printf("Movies:%#v\n", movies)
}
