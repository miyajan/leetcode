package main

import (
	"container/heap"
)

type Item struct {
	value    []int
	current  int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return [][]int{}
	}

	pq := PriorityQueue{}
	for i := 0; i < len(nums1) && i < k; i++ {
		pq = append(pq, &Item{
			value:    []int{nums1[i], nums2[0]},
			current:  0,
			priority: nums1[i] + nums2[0],
			index:    i,
		})
	}
	heap.Init(&pq)

	var ret [][]int
	for k > 0 && len(pq) > 0 {
		item := heap.Pop(&pq).(*Item)
		ret = append(ret, item.value)
		k--
		if item.current < len(nums2)-1 {
			heap.Push(&pq, &Item{
				value:    []int{item.value[0], nums2[item.current+1]},
				current:  item.current + 1,
				priority: item.value[0] + nums2[item.current+1],
			})
		}
	}

	return ret
}
