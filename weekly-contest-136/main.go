package main

import (
	"log"
	"sort"
)

func isRobotBounded(instructions string) bool {
	point := []int{0, 0}
	direction := 0

	for i := 0; i < 4; i++ {
		for j := 0; j < len(instructions); j++ {
			switch instructions[j] {
			case 'G':
				switch direction {
				case 0:
					point[1]++
				case 1:
					point[0]--
				case 2:
					point[1]--
				case 3:
					point[0]++
				default:
					log.Panic("error!")
				}
			case 'L':
				direction++
				if direction > 3 {
					direction = 0
				}
			case 'R':
				direction--
				if direction < 0 {
					direction = 3
				}
			}
		}

		if point[0] == 0 && point[1] == 0 {
			return true
		}
	}

	return false
}

type garden struct {
	number    int
	color     int
	adjascent map[int]*garden
}

func gardenNoAdj(N int, paths [][]int) []int {
	answer := make([]int, N)
	gardens := map[int]*garden{}
	for i := 1; i <= N; i++ {
		gardens[i] = &garden{
			number:    i,
			adjascent: map[int]*garden{},
		}
	}
	for _, path := range paths {
		gardens[path[0]].adjascent[path[1]] = gardens[path[1]]
		gardens[path[1]].adjascent[path[0]] = gardens[path[0]]
	}

	for i := 1; i <= N; i++ {
		g := gardens[i]
		colors := map[int]bool{1: true, 2: true, 3: true, 4: true}
		for _, val := range g.adjascent {
			if val.color != 0 {
				delete(colors, val.color)
			}
		}

		for key := range colors {
			g.color = key
			answer[g.number-1] = key
			break
		}
	}

	return answer
}

func maxSumAfterPartitioning(A []int, K int) int {
	dp := make([]int, len(A)+1)
	dp[0] = 0
	dp[1] = A[0]

	for i := 2; i <= len(A); i++ {
		var max int
		for j := 1; j <= K && i-j >= 0; j++ {
			var m int
			for l := 1; l <= j; l++ {
				if A[i-l] > m {
					m = A[i-l]
				}
			}
			v := dp[i-j] + m*j
			if v > max {
				max = v
			}
		}
		dp[i] = max
	}

	return dp[len(A)]
}

type node struct {
	count  int
	length int
	str    string
	next   map[byte]*node
}

func longestDupSubstring(S string) string {
	suffixArr := buildSuffixArray(S)
	lcp := buildLongestCommonPrefix(S, suffixArr)

	var longest string
	for i := 0; i < len(S); i++ {
		if lcp[i] > len(longest) {
			longest = S[suffixArr[i]:][:lcp[i]]
		}
	}

	return longest
}

type suffix struct {
	index int
	rank  [2]int
}

func buildSuffixArray(S string) []int {
	n := len(S)
	suffixes := make([]suffix, n)

	for i := 0; i < n; i++ {
		var r int
		if i < n-1 {
			r = int(S[i+1] - 'a')
		} else {
			r = -1
		}
		suffixes[i] = suffix{
			index: i,
			rank:  [2]int{int(S[i] - 'a'), r},
		}
	}

	sort.Slice(suffixes, func(i, j int) bool {
		if suffixes[i].rank[0] == suffixes[j].rank[0] {
			return suffixes[i].rank[1] < suffixes[j].rank[1]
		}
		return suffixes[i].rank[0] < suffixes[j].rank[0]
	})

	index := make([]int, n)
	for k := 4; k < 2*n; k *= 2 {
		rank := 0
		prevRank := suffixes[0].rank[0]
		suffixes[0].rank[0] = rank
		index[suffixes[0].index] = 0

		for i := 1; i < n; i++ {
			if suffixes[i].rank[0] == prevRank && suffixes[i].rank[1] == suffixes[i-1].rank[1] {
				suffixes[i].rank[0] = rank
			} else {
				prevRank = suffixes[i].rank[0]
				rank++
				suffixes[i].rank[0] = rank
			}
			index[suffixes[i].index] = i
		}

		for i := 0; i < n; i++ {
			nextIndex := suffixes[i].index + k/2
			if nextIndex < n {
				suffixes[i].rank[1] = suffixes[index[nextIndex]].rank[0]
			} else {
				suffixes[i].rank[1] = -1
			}
		}

		sort.Slice(suffixes, func(i, j int) bool {
			if suffixes[i].rank[0] == suffixes[j].rank[0] {
				return suffixes[i].rank[1] < suffixes[j].rank[1]
			}
			return suffixes[i].rank[0] < suffixes[j].rank[0]
		})
	}

	suffixArr := make([]int, n)
	for i := 0; i < n; i++ {
		suffixArr[i] = suffixes[i].index
	}

	return suffixArr
}

func buildLongestCommonPrefix(S string, suffixArr []int) []int {
	n := len(S)
	lcp := make([]int, n)

	invSuffix := make([]int, n)
	for i := 0; i < n; i++ {
		invSuffix[suffixArr[i]] = i
	}

	var k int
	for i := 0; i < n; i++ {
		if invSuffix[i] == n-1 {
			k = 0
			continue
		}

		j := suffixArr[invSuffix[i]+1]
		for i+k < n && j+k < n && S[i+k] == S[j+k] {
			k++
		}
		lcp[invSuffix[i]] = k

		if k > 0 {
			k--
		}
	}

	return lcp
}
