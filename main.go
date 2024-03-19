package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(words []string) [][]string {
	anagramGroups := [][]string{}
	visited := make([]bool, len(words))

	for i := 0; i < len(words); i++ {
		if !visited[i] {
			anagramGroup := []string{words[i]}
			visited[i] = true

			for j := 0; j < len(words); j++ {
				if !visited[j] && isAnagram(words[i], words[j]) {
					anagramGroup = append(anagramGroup, words[j])
					visited[j] = true
				}
			}

			anagramGroups = append(anagramGroups, anagramGroup)
		}
	}

	return anagramGroups
}

func isAnagram(word1, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	// word1Chars := []rune(word1)
	// word2Chars := []rune(word2)

	// sort.Slice(word1Chars, func(i, j int) bool {
	// 	return word1Chars[i] < word1Chars[j]
	// })

	// sort.Slice(word2Chars, func(i, j int) bool {
	// 	return word2Chars[i] < word2Chars[j]
	// })

	word1Chars := strings.Split(word1, "")
	word2Chars := strings.Split(word2, "")

	sort.Strings(word1Chars)
	sort.Strings(word2Chars)

	for i := 0; i < len(word1Chars); i++ {
		if word1Chars[i] != word2Chars[i] {
			return false
		}
	}

	return true
}

func main() {
	words := []string{"cook", "save", "teste", "aves", "vase", "state", "map"}
	result := groupAnagrams(words)
	fmt.Println(result)
}
