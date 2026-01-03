package index

import (
	"log"
	"os"
	"search-engine/internal/helpers"
	"strconv"
	"strings"
)

const (
	DOCS_DIR = "internal/data/docs"
)

func FillData(idx *InvertedIndex) {

	wd, _ := os.Getwd()
	log.Printf("Working directory: %s", wd)
	var dirPath = wd + "/" + DOCS_DIR

	files := helpers.GetFiles(dirPath)
	for _, fileName := range files {
		content := helpers.GetFileContent(dirPath + "/" + fileName)
		// Split content into words factor in new lines
		words := strings.Fields(string(content))

		docIDStr := strings.TrimSuffix(fileName, ".txt")
		docIDStr = strings.TrimPrefix(docIDStr, "doc")
		docID, err := strconv.Atoi(docIDStr)
		if err != nil {
			log.Fatalf("Failed to convert docID %s to int: %v", docIDStr, err)
		}

		seen := make(map[string]bool)
		for _, word := range words {
			sanitizedWord := helpers.SanitizeWord(word)
			if sanitizedWord == "" || seen[sanitizedWord] {
				continue
			}
			seen[sanitizedWord] = true
			(*idx)[sanitizedWord] = append((*idx)[sanitizedWord], docID)

		}

	}
}
