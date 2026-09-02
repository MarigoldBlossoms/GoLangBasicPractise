package practice

import "strings"

/*
Exercise 010:

Accept a whitespace-separated string of words, remove duplicates and return
the remaining words sorted alphanumerically and joined by a single space.

Example:
  Ex010("hello world and practice makes perfect and hello world again")
    -> "again and hello makes perfect practice world"

Tip: run `go test ./...` from this folder.
*/

// Ex010 should deduplicate, sort and re-join the words.
func Ex010(input string) string {
	wordList := strings.Split(input, " ")

	outputList := RemoveDuplicates(wordList)
	outputList = SortSlice(outputList)

	return strings.Join(outputList, " ")
}

func RemoveDuplicates(input []string) []string {
	uniqueList := make(map[string]struct{})

	var outputList []string

	for _, word := range input {
		if _, ok := uniqueList[word]; !ok {
			uniqueList[word] = struct{}{}
			outputList = append(outputList, word)
		}
	}

	return outputList
}

func SortSlice(input []string) []string {
	sorting := true
	for sorting {
		sorting = false
		for i := 0; i < len(input)-1; i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
				sorting = true
			}
		}
	}
	return input
}
