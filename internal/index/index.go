package index

// InvertedIndex maps a word to a sorted list of document IDs.
type InvertedIndex map[string][]int
