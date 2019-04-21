package main

import (
	"math"
	"sort"
)

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	ret := [][]int{}
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			ret = append(ret, []int{i, j})
		}
	}

	sort.Slice(ret, func(i, j int) bool {
		return distance(ret[i], r0, c0) < distance(ret[j], r0, c0)
	})

	return ret
}

func distance(cell []int, r0 int, c0 int) int {
	return int(math.Abs(float64(cell[0]-r0)) + math.Abs(float64(cell[1]-c0)))
}

func twoCitySchedCost(costs [][]int) int {
	length := len(costs)
	var min int

	sort.Slice(costs, func(i, j int) bool {
		return costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1]
	})

	aCosts := costs[0 : length/2]
	bCosts := costs[length/2:]

	for i := 0; i < len(aCosts); i++ {
		min += aCosts[i][0]
		min += bCosts[i][1]
	}

	return min
}

func maxSumTwoNoOverlap(A []int, L int, M int) int {
	var max int
	for i := 0; i < len(A)-L+1; i++ {
		sumL := 0
		for j := i; j < i+L; j++ {
			sumL += A[j]
		}

		for j := 0; j < i-M+1; j++ {
			sumM := 0
			for k := j; k < j+M; k++ {
				sumM += A[k]
			}

			if sumL+sumM > max {
				max = sumL + sumM
			}
		}

		for j := i + L; j < len(A)-M+1; j++ {
			sumM := 0
			for k := j; k < j+M; k++ {
				sumM += A[k]
			}

			if sumL+sumM > max {
				max = sumL + sumM
			}
		}
	}
	return max
}

type StreamChecker struct {
	nodes   map[byte]*CharNode
	history []byte
}

type CharNode struct {
	nodes    map[byte]*CharNode
	terminal bool
}

func Constructor(words []string) StreamChecker {
	checker := StreamChecker{
		history: []byte{},
	}
	nodes := make(map[byte]*CharNode)
	for i := 0; i < len(words); i++ {
		word := words[i]
		currentNodes := nodes
		for j := len(word) - 1; j >= 0; j-- {
			char := word[j]
			if _, ok := currentNodes[char]; !ok {
				currentNodes[char] = &CharNode{
					nodes: make(map[byte]*CharNode),
				}
			}
			if j == 0 {
				node := currentNodes[char]
				node.terminal = true
			}
			currentNodes = currentNodes[char].nodes
		}
	}
	checker.nodes = nodes
	return checker
}

func (this *StreamChecker) Query(letter byte) bool {
	this.history = append(this.history, letter)
	if _, ok := this.nodes[letter]; !ok {
		return false
	}

	currentNode := this.nodes[letter]
	if currentNode.terminal {
		return true
	}

	for i := len(this.history) - 2; i >= 0; i-- {
		if _, ok := currentNode.nodes[this.history[i]]; !ok {
			return false
		}
		currentNode = currentNode.nodes[this.history[i]]
		if currentNode.terminal {
			return true
		}
	}

	return currentNode.terminal
}
