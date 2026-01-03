package helpers

import (
	"bufio"
	"io/fs"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	DOCS_DIR = "internal/data/docs"
)

func GetFiles(dir string) []string {

	root := os.DirFS(dir)
	txtFiles, err := fs.Glob(root, "*.txt")
	if err != nil {
		log.Fatalf("Failed to glob files: %v", err)
	}
	return txtFiles
}

func SanitizeWord(word string) string {
	word = strings.TrimSpace(word)
	if word == "" {
		return word
	}
	// Lowercase the word for case-insensitive search
	word = strings.ToLower(word)
	// Remove punctuation from the word
	return strings.Map(func(r rune) rune {
		if strings.ContainsRune(".,!?;:\"'()[]{}", r) {
			return -1
		}
		return r
	}, word)

}

func ReadUserInput() []string {
	var reader = bufio.NewReader(os.Stdin)
	log.Print("Enter search query: ")
	message, err := reader.ReadString('\n')
	if err != nil {
		panic("Failed to read user input")
	}
	words := strings.Split(strings.TrimSpace(message), " ")
	var sanitizedWords []string
	for _, word := range words {
		sanitizedWord := SanitizeWord(word)
		if sanitizedWord != "" {
			sanitizedWords = append(sanitizedWords, sanitizedWord)
		}
	}
	return sanitizedWords
}

func GetFileContent(filePath string) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", filePath, err)
	}
	return string(content)
}

func GetSearchedContent(docID []int) []string {

	wd, _ := os.Getwd()
	log.Printf("Working directory: %s", wd)
	var dirPath = wd + "/" + DOCS_DIR
	results := []string{}
	for _, id := range docID {
		var fileName = "doc" + strconv.Itoa(id) + ".txt"
		content := GetFileContent(dirPath + "/" + fileName)
		results = append(results, content)
	}
	return results
}
