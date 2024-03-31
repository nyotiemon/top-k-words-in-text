package main

import (
	"fmt"
	"strings"
)

type WordNode struct {
	str   string
	count int
}

func getmin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lowerAndTrimDot(str string) string {
	str = strings.ToLower(str)
	if str[len(str)-1] == '.' {
		str = str[:len(str)-1]
	}
	return str
}

func main() {
	wordNode := SliceTextToMapNode("test end with dot.")
	fmt.Printf("#1 len %v, should be %v, correct:%v\n", len(wordNode), 4, len(wordNode) == 4)
	wordNode = SliceTextToMapNode("test end with space ")
	fmt.Printf("#2 len %v, should be %v, correct:%v\n", len(wordNode), 4, len(wordNode) == 4)
	wordNode = SliceTextToMapNode("test end with nothing")
	fmt.Printf("#3 len %v, should be %v, correct:%v\n", len(wordNode), 4, len(wordNode) == 4)
	wordNode = SplitTextToMapNode("test end with nothing")
	fmt.Printf("#4 len %v, should be %v, correct:%v\n", len(wordNode), 4, len(wordNode) == 4)
	wordStr := SplitTextToMapStr("test end with nothing")
	fmt.Printf("#5 len %v, should be %v, correct:%v\n", len(wordStr), 4, len(wordStr) == 4)

	top := 3 // input word count is guaranteed have more than top value
	wordNode = SplitTextToMapNode("Hello world. Isn't isn't isn't This is some 3 sample text. Text is short and simple.")
	fmt.Println("wordNode", wordNode)

	result := TopKBinarySearchMapNode(wordNode, top)
	fmt.Printf("top1 should be [%s:%v], actual [%s:%v], correct: %v\n",
		"isn't", 3,
		result[0].str, result[0].count,
		result[0].str == "isn't" && result[0].count == 3)
	fmt.Printf("top2 should be [%s] or [%s]:%v, actual [%s:%v], correct: %v\n",
		"text", "is", 2,
		result[1].str, result[1].count,
		(result[1].str == "text" || result[1].str == "is") && result[1].count == 2)
	fmt.Printf("top3 should be [%s] or [%s]:%v, actual [%s:%v], correct: %v\n",
		"text", "is", 2,
		result[2].str, result[2].count,
		(result[2].str == "text" || result[2].str == "is") && result[2].count == 2)

	result = TopKSortedArrayFromMapNode(wordNode, top)
	fmt.Printf("top1 should be [%s:%v], actual [%s:%v], correct: %v\n",
		"isn't", 3,
		result[0].str, result[0].count,
		result[0].str == "isn't" && result[0].count == 3)
	fmt.Printf("top2 should be [%s] or [%s]:%v, actual [%s:%v], correct: %v\n",
		"text", "is", 2,
		result[1].str, result[1].count,
		(result[1].str == "text" || result[1].str == "is") && result[1].count == 2)
	fmt.Printf("top3 should be [%s] or [%s]:%v, actual [%s:%v], correct: %v\n",
		"text", "is", 2,
		result[2].str, result[2].count,
		(result[2].str == "text" || result[2].str == "is") && result[2].count == 2)
}
