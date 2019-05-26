package main

import (
	"math"
	"sort"
)

func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Ints(stones)
		x := stones[len(stones)-2]
		y := stones[len(stones)-1]
		stones = stones[0 : len(stones)-2]
		if y > x {
			stones = append(stones, y-x)
		}
	}
	if len(stones) == 0 {
		return 0
	}
	return stones[0]
}

func removeDuplicates(S string) string {
	var stack []byte
	for i := 0; i < len(S); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == S[i] {
			stack = stack[0 : len(stack)-1]
		} else {
			stack = append(stack, S[i])
		}
	}
	return string(stack)
}

func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) <= len(words[j])
	})

	chainLength := map[string]int{}
	strLen := map[int][]string{}

	var longest int
	for _, word := range words {
		maxLength := 1
		for _, s := range strLen[len(word)-1] {
			if isPredecessor(s, word) && chainLength[s]+1 > maxLength {
				maxLength = chainLength[s] + 1
			}
		}
		if maxLength > longest {
			longest = maxLength
		}
		chainLength[word] = maxLength
		if _, ok := strLen[len(word)]; !ok {
			strLen[len(word)] = []string{}
		}
		strLen[len(word)] = append(strLen[len(word)], word)
	}

	return longest
}

func isPredecessor(a, b string) bool {
	j := 0
	for i := 0; i < len(a); i++ {
		for {
			if a[i] == b[j] {
				break
			}
			j++
			if j >= len(b) || j-i > 1 {
				return false
			}
		}
	}
	return true
}

func lastStoneWeightII(stones []int) int {
	var sum int
	dp := make([]bool, 1501)
	dp[0] = true
	for i := 0; i < len(stones); i++ {
		sum += stones[i]
		for j := 1500; j >= stones[i]; j-- {
			dp[j] = dp[j] || dp[j-stones[i]]
		}
	}

	min := math.MaxInt32
	for i := 0; i <= 1500; i++ {
		if dp[i] {
			a := sum - i
			b := i
			var w int
			if a > b {
				w = a - b
			} else {
				w = b - a
			}
			if w < min {
				min = w
			}
		}
	}

	return min
}
