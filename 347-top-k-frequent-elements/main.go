package main

func topKFrequent(nums []int, k int) []int {
	elements := make(map[int]int)
	for _, num := range nums {
		elements[num]++
	}

	frequencies := make([][]int, len(nums)+1)
	for key := range elements {
		frequencies[elements[key]] = append(frequencies[elements[key]], key)
	}

	var ret []int
	for i := len(nums); i > 0 && k > 0; i-- {
		if len(frequencies[i]) > 0 {
			var count int
			if len(frequencies[i]) > k {
				count = k
			} else {
				count = len(frequencies[i])
			}
			ret = append(ret, frequencies[i][0:count]...)
			k -= count
		}
	}

	return ret
}
