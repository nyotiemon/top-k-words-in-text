package main

import (
	"sort"
	"strings"
)

func OldGetWordCount(str string) map[string]int {
	res := make(map[string]int, 0)

	// split the words by spaces
	// remove the (.) at the end of the word character
	// put it into the dict, inc the value by 1

	temp := strings.Split(str, " ")
	for _, word := range temp {
		if word[len(word)-1] == '.' {
			res[strings.ToLower(word[:len(word)-1])]++
		} else {
			res[strings.ToLower(word)]++
		}
	}

	return res
}

/**
Time complexity
O(n) with n as number of words in array
**/

func OldGetTopK(words map[string]int, k int) []*WordNode {
	result := []*WordNode{}
	for key, value := range words {
		result = append(result, &WordNode{str: key, count: value})
	}

	// only preserve top 3 words in res
	sort.Slice(result, func(a, b int) bool {
		return result[a].count > result[b].count
	})

	return result[:k]
}

/**
Time complexity
O(n) for converting map to array
O(n log n) for sorting
**/
