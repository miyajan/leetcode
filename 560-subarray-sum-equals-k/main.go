package main

func subarraySum(nums []int, k int) int {
	m := map[int]int{0: 1}
	var ret, sum int
	for _, num := range nums {
		sum += num
		ret += m[sum - k]
		m[sum]++
	}
	return ret
}
