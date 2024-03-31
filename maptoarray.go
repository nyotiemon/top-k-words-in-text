package main

func MapNodeToArray(words map[string]*WordNode) []*WordNode {
	result := []*WordNode{}
	for _, word := range words {
		result = append(result, word)
	}
	return result
}

func MapStrToArray(words map[string]int) []*WordNode {
	result := []*WordNode{}
	for key, value := range words {
		result = append(result, &WordNode{str: key, count: value})
	}
	return result
}
