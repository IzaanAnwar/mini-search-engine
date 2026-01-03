package search

import "search-engine/internal/index"

func Query(words []string, idx *index.InvertedIndex) []int {
	var result []int = nil

	for _, word := range words {
		postings, ok := (*idx)[word]
		if !ok {
			return []int{}
		}
		if result == nil {
			result = append([]int(nil), postings...)
		} else {
			result = Intersect(result, postings)
			if len(result) == 0 {
				return result
			}
		}
	}

	return result
}

func Intersect(a, b []int) []int {
	i, j := 0, 0
	result := []int{}
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
			j++
			result = append(result, a[i-1])
		} else if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		}
	}
	return result
}
