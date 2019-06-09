package main

import (
	"sort"
	"strings"
)

func findOcurrences(text string, first string, second string) []string {
	texts := strings.Split(text, " ")
	var prePrev string
	var prev string
	var ret []string
	for _, t := range texts {
		if prePrev == first && prev == second {
			ret = append(ret, t)
		}
		prePrev = prev
		prev = t
	}
	return ret
}

func buildPossibilities(cur string, tiles string, seen map[string]bool) {
	seen[cur] = true

	for i := 0; i < len(tiles); i++ {
		buildPossibilities(cur+string(tiles[i]), tiles[0:i]+tiles[i+1:], seen)
	}
}

func numTilePossibilities(tiles string) int {
	seen := make(map[string]bool)
	for i := 0; i < len(tiles); i++ {
		buildPossibilities(string(tiles[i]), tiles[0:i]+tiles[i+1:], seen)
	}

	return len(seen)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	node, sum := process(root, 0, limit)
	if sum < limit {
		return nil
	}
	return node
}

func process(node *TreeNode, parentSum int, limit int) (*TreeNode, int) {
	if node == nil {
		return nil, parentSum
	}

	sum := parentSum + node.Val

	leftNode, leftSum := process(node.Left, sum, limit)
	rightNode, rightSum := process(node.Right, sum, limit)

	node.Left = leftNode
	node.Right = rightNode

	if leftSum > rightSum {
		sum = leftSum
	} else {
		sum = rightSum
	}

	if sum < limit {
		return nil, sum
	}

	return node, sum
}

func smallestSubsequence(text string) string {
	indexes := make(map[byte][]int)
	var keys []byte
	for i := 0; i < len(text); i++ {
		if _, ok := indexes[text[i]]; !ok {
			indexes[text[i]] = []int{}
			keys = append(keys, text[i])
		}
		indexes[text[i]] = append(indexes[text[i]], i)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	var smallest string
	smallestLength := len(keys)
	for len(smallest) < smallestLength {
		for i := 0; i < len(keys); i++ {
			firstOccur := indexes[keys[i]][0]
			allKeysExistAfterFirstOccur := true
			for j := i + 1; j < len(keys); j++ {
				if firstOccur > indexes[keys[j]][len(indexes[keys[j]])-1] {
					allKeysExistAfterFirstOccur = false
				}
			}

			if allKeysExistAfterFirstOccur {
				smallest += string(keys[i])
				keys = append(keys[0:i], keys[i+1:]...)
				for k, v := range indexes {
					var next []int
					for _, index := range v {
						if index > firstOccur {
							next = append(next, index)
						}
					}
					indexes[k] = next
				}
				break
			}
		}
	}

	return smallest
}
