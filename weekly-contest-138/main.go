package main

import (
	"container/heap"
	"sort"
)

func heightChecker(heights []int) int {
	expected := make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		expected[i] = heights[i]
	}
	sort.Ints(expected)

	var ret int
	for i := 0; i < len(heights); i++ {
		if heights[i] != expected[i] {
			ret++
		}
	}

	return ret
}

func maxSatisfied(customers []int, grumpy []int, X int) int {
	unsatisfiedPrefixSum := make([]int, len(customers)+1)
	unsatisfiedPrefixSum[0] = 0
	var satisfiedSum int
	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 1 {
			unsatisfiedPrefixSum[i+1] = unsatisfiedPrefixSum[i] + customers[i]
		} else {
			unsatisfiedPrefixSum[i+1] = unsatisfiedPrefixSum[i]
			satisfiedSum += customers[i]
		}
	}

	var maxKeepSatisfied int
	for i := X; i <= len(customers); i++ {
		if unsatisfiedPrefixSum[i]-unsatisfiedPrefixSum[i-X] > maxKeepSatisfied {
			maxKeepSatisfied = unsatisfiedPrefixSum[i] - unsatisfiedPrefixSum[i-X]
		}
	}

	return satisfiedSum + maxKeepSatisfied
}

func prevPermOpt1(A []int) []int {
	if len(A) == 1 {
		return A
	}

	min := make([][]int, len(A))
	for i := 0; i < len(A); i++ {
		min[i] = []int{0, -1}
	}
	for i := 0; i < len(A); i++ {
		for j := i - 1; j >= 0; j-- {
			if A[j] > A[i] && A[i] >= min[j][0] {
				min[j] = []int{A[i], i}
				break
			}
		}
	}
	for i := len(A) - 2; i >= 0; i-- {
		if min[i][0] != 0 {
			A[i], A[min[i][1]] = A[min[i][1]], A[i]
			break
		}
	}
	return A
}

type LengthHeap [][]int

func (h LengthHeap) Len() int           { return len(h) }
func (h LengthHeap) Less(i, j int) bool { return h[i][1] > h[j][1] }
func (h LengthHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *LengthHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *LengthHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func rearrangeBarcodes(barcodes []int) []int {
	codeLength := &LengthHeap{}
	heap.Init(codeLength)
	codeLengthMap := map[int][]int{}
	for i := 0; i < len(barcodes); i++ {
		if _, ok := codeLengthMap[barcodes[i]]; !ok {
			length := []int{barcodes[i], 1}
			codeLengthMap[barcodes[i]] = length
		} else {
			codeLengthMap[barcodes[i]][1]++
		}
	}

	for _, v := range codeLengthMap {
		heap.Push(codeLength, v)
	}

	var ret []int
	var previous int
	for i := 0; i < len(barcodes); i++ {
		l1 := heap.Pop(codeLength).([]int)
		if l1[0] != previous {
			ret = append(ret, l1[0])
			previous = l1[0]
			l1[1]--
			if l1[1] != 0 {
				heap.Push(codeLength, l1)
			}
		} else {
			l2 := heap.Pop(codeLength).([]int)
			ret = append(ret, l2[0])
			previous = l2[0]
			l2[1]--
			heap.Push(codeLength, l1)
			if l2[1] != 0 {
				heap.Push(codeLength, l2)
			}
		}
	}

	return ret
}
