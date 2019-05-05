package main

import (
	"math"
	"sort"
)

func isBoomerang(points [][]int) bool {
	if (points[0][0] == points[1][0] && points[0][1] == points[1][1]) ||
		(points[0][0] == points[2][0] && points[0][1] == points[2][1]) ||
		(points[1][0] == points[2][0] && points[1][1] == points[2][1]) {
		return false
	}

	if (points[0][0] == points[1][0] && points[0][0] == points[2][0]) ||
		(points[0][1] == points[1][1] && points[0][1] == points[2][1]) {
		return false
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] <= points[j][0]
	})

	return (float64(points[1][1]-points[0][1]) / float64(points[1][0]-points[0][0])) !=
		(float64(points[2][1]-points[1][1]) / float64(points[2][0]-points[1][0]))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	nodes := collect(root)
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Val < nodes[j].Val
	})
	var sum int
	for i := 0; i < len(nodes); i++ {
		node := nodes[len(nodes)-1-i]
		sum += node.Val
		node.Val = sum
	}
	return root
}

func collect(node *TreeNode) []*TreeNode {
	if node == nil {
		return []*TreeNode{}
	}
	nodes := []*TreeNode{node}
	nodes = append(nodes, collect(node.Left)...)
	nodes = append(nodes, collect(node.Right)...)
	return nodes
}

func minScoreTriangulation(A []int) int {
	dp := make(map[int]map[int]int)
	for i := 0; i < len(A); i++ {
		dp[i] = make(map[int]int)
	}
	return minScore(0, len(A)-1, A, dp)
}

func minScore(p1 int, p2 int, A []int, dp map[int]map[int]int) int {
	if p2-p1 <= 1 {
		return 0
	}

	if _, ok := dp[p1][p2]; ok {
		return dp[p1][p2]
	}

	min := math.MaxInt32
	for i := p1 + 1; i < p2; i++ {
		score := minScore(p1, i, A, dp) + minScore(i, p2, A, dp) + A[p1]*A[p2]*A[i]
		if score < min {
			min = score
		}
	}
	dp[p1][p2] = min

	return min
}

func numMovesStonesII(stones []int) []int {
	sort.Slice(stones, func(i, j int) bool {
		return stones[i] < stones[j]
	})

	var j int
	var maxWindowLength int
	var start, end int
	for i := 0; i < len(stones); i++ {
		for j < len(stones) && stones[j]-stones[i] <= len(stones)-1 {
			j++
		}

		if j-i > maxWindowLength || (j-i == maxWindowLength && stones[j-1]-stones[i] > end-start) {
			maxWindowLength = j - i
			start = stones[i]
			end = stones[j-1]
		}
	}

	min := len(stones) - maxWindowLength
	if maxWindowLength == len(stones)-1 && end-start == len(stones)-2 {
		min++
	}

	interval1 := stones[len(stones)-1] - stones[1] - 1
	interval2 := stones[len(stones)-2] - stones[0] - 1
	var maxInterval int
	if interval1 > interval2 {
		maxInterval = interval1
	} else {
		maxInterval = interval2
	}
	max := maxInterval - (len(stones) - 3)

	return []int{min, max}
}
