package main

import (
	"log"
	"search-engine/internal/helpers"
	"search-engine/internal/index"
	"search-engine/internal/search"
)

func main() {
	idx := make(index.InvertedIndex)
	index.FillData(&idx)
	log.Printf("Inverted Index built with %v terms", len(idx))

	userInput := helpers.ReadUserInput()

	res := search.Query(userInput, &idx)
	log.Printf("Search results: %v", res)
	matches := helpers.GetSearchedContent(res)
	for i, content := range matches {
		log.Printf("Document ID: %d\nContent:\n%s\n", res[i], content)
	}
}
