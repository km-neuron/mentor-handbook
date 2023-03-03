package main

import (
	"strings"
)

func CountRepeatedWords(input string) int {
	if input == "" {
		return 0
	}

	words := strings.Split(input, " ")
	if len(words) == 1 {
		return 1
	}

	firstWord := words[0]
	repeatedCount := 0

	for _, word := range words {
		if word == firstWord {
			repeatedCount++
		} else {
			break
		}
	}

	return repeatedCount
}
