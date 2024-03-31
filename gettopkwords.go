package main

import "sort"

func TopKBinarySearchMapStr(words map[string]int, k int) []*WordNode {
	result := make([]*WordNode, 0, k+1) // +1 to prevent append fn do reallocation
	min := 0
	for str, count := range words {
		if len(result) < 1 {
			result = append(result, &WordNode{str: str, count: count})
			min = getmin(min, count)
		} else if len(result) < k || count >= min {
			// binary search to find insert position

			l, r := 0, len(result)
			for l < r {
				mid := l + (r-l)/2
				if count >= result[mid].count {
					r = mid
				} else {
					l = mid + 1
				}
			}

			result = append(result[:l+1], result[l:]...)
			result[l] = &WordNode{str: str, count: count}

			result = result[:getmin(len(result), k)]

			min = result[len(result)-1].count
		}
	}
	return result
}

func TopKBinarySearchMapNode(words map[string]*WordNode, k int) []*WordNode {
	result := make([]*WordNode, 0, k+1) // +1 to prevent append fn do reallocation
	min := 0
	for _, word := range words {
		if len(result) < 1 {
			result = append(result, word)
			min = getmin(min, word.count)
		} else if len(result) < k || word.count >= min {
			// binary search to find insert position

			l, r := 0, len(result)
			for l < r {
				mid := l + (r-l)/2
				if word.count >= result[mid].count {
					r = mid
				} else {
					l = mid + 1
				}
			}
			result = append(result[:l+1], result[l:]...)
			result[l] = word

			result = result[:getmin(len(result), k)]

			min = result[len(result)-1].count
		}
	}
	return result
}

/**
Time complexity
O(n^logk) with n as number of words in map, and k as top-k words

Space complexity
O(k+1) -> O(k) with k as top-k words
**/

func TopKSortedArrayFromMapStr(words map[string]int, k int) []*WordNode {
	result := []*WordNode{}
	for key, value := range words {
		result = append(result, &WordNode{str: key, count: value})
	}

	sort.Slice(result, func(a, b int) bool {
		return result[a].count > result[b].count
	})

	return result[:k]
}

func TopKSortedArrayFromMapNode(words map[string]*WordNode, k int) []*WordNode {
	result := []*WordNode{}
	for _, word := range words {
		result = append(result, word)
	}

	sort.Slice(result, func(a, b int) bool {
		return result[a].count > result[b].count
	})

	return result[:k]
}

/**
Time complexity
O(n) for converting map to array
O(n log n) for sorting

Space complexity
O(n) with n as number of word in map
**/
