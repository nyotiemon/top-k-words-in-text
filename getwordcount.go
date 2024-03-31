package main

import "strings"

func SplitTextToMapNode(text string) map[string]*WordNode {
	words := strings.Split(text, " ")

	result := make(map[string]*WordNode, 0)
	for _, word := range words {
		word = lowerAndTrimDot(word)
		if v, e := result[word]; !e {
			result[word] = &WordNode{str: word, count: 1}
		} else {
			v.count++
		}
	}
	return result
}

func SplitTextToMapStr(text string) map[string]int {
	words := strings.Split(text, " ")

	result := make(map[string]int, 0)
	for _, word := range words {
		result[lowerAndTrimDot(word)]++
	}
	return result
}

func SliceTextToMapNode(input string) map[string]*WordNode {
	dict := make(map[string]*WordNode, 0)
	i, j := 0, 0
	for i < len(input)-1 && j < len(input)-1 {
		switch input[j] {
		case ' ':
			word := lowerAndTrimDot(input[i:j])
			if v, e := dict[word]; !e {
				dict[word] = &WordNode{str: word, count: 1}
			} else {
				v.count++
			}
			i = j + 1
		}

		j++
	}
	if i < j {
		word := lowerAndTrimDot(input[i:j])
		if v, e := dict[word]; !e {
			dict[word] = &WordNode{str: word, count: 1}
		} else {
			v.count++
		}
	}

	return dict
}

/**
Time complexity (ignoring split func)
SplitTextToMapNode = O(n) with n as number of words in array
SplitTextToMapStr = O(n) with n as number of words in array
SliceTextToMapNode = O(m) with m as number of letter in text

Space complexity
Split = O(n + m) = O(n) with n for array length, and m for map length
Slice = O(n) = O(n) with m for map length
**/
