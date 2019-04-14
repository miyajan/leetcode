package main

import (
	"math"
	"strconv"
)

func divisorGame(N int) bool {
	results := make([]bool, N+1)
	results[1] = false
	for i := 2; i <= N; i++ {
		for j := 1; j < i; j++ {
			if i%j == 0 && !results[i-j] {
				results[i] = true
				break
			}
		}
	}
	return results[N]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	return maxAncestorDiffRecursively(root, 0, []int{root.Val})
}

func maxAncestorDiffRecursively(node *TreeNode, currentMaxDiff int, ancestors []int) int {
	maxDiff := currentMaxDiff
	for _, ancestor := range ancestors {
		diff := int(math.Abs(float64(node.Val - ancestor)))
		if diff > maxDiff {
			maxDiff = diff
		}
	}

	nextAncestors := append(ancestors, node.Val)

	if node.Left != nil {
		leftMax := maxAncestorDiffRecursively(node.Left, maxDiff, nextAncestors)
		if leftMax > maxDiff {
			maxDiff = leftMax
		}
	}

	if node.Right != nil {
		rightMax := maxAncestorDiffRecursively(node.Right, maxDiff, nextAncestors)
		if rightMax > maxDiff {
			maxDiff = rightMax
		}
	}

	return maxDiff
}

func longestArithSeqLength(A []int) int {
	var longest int

	subsequenceLengths := make(map[int]map[int]int)
	for i := 0; i < len(A); i++ {
		for j := 0; j < i; j++ {
			diff := A[j] - A[i]
			if _, ok := subsequenceLengths[j]; !ok {
				subsequenceLengths[j] = make(map[int]int)
			}
			if _, ok := subsequenceLengths[j][diff]; !ok {
				subsequenceLengths[j][diff] = 1
			}
			if _, ok := subsequenceLengths[i]; !ok {
				subsequenceLengths[i] = make(map[int]int)
			}
			subsequenceLengths[i][diff] = subsequenceLengths[j][diff] + 1
			if subsequenceLengths[i][diff] > longest {
				longest = subsequenceLengths[i][diff]
			}
		}
	}

	return longest
}

func recoverFromPreorder(S string) *TreeNode {
	dashByte := "-"[0]
	i := 0

	var rootVal []byte
	for i < len(S) && S[i] != dashByte {
		rootVal = append(rootVal, S[i])
		i++
	}
	val, _ := strconv.Atoi(string(rootVal))
	root := &TreeNode{
		Val: val,
	}

	currentNode := make(map[int]*TreeNode)
	currentNode[0] = root
	j := i
	for j < len(S) {
		depth := 0
		for S[j] == dashByte {
			depth++
			j++
		}

		var valBytes []byte
		for j < len(S) && S[j] != dashByte {
			valBytes = append(valBytes, S[j])
			j++
		}

		val, _ := strconv.Atoi(string(valBytes))
		node := &TreeNode{
			Val: val,
		}

		parentNode := currentNode[depth-1]
		if parentNode.Left == nil {
			parentNode.Left = node
		} else {
			parentNode.Right = node
		}

		currentNode[depth] = node
	}

	return root
}
