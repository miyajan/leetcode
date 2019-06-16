package main

import (
	"sort"
)

func duplicateZeros(arr []int) {
	var ret []int
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			ret = append(ret, 0)
			if len(ret) == len(arr) {
				break
			}
			ret = append(ret, 0)
			if len(ret) == len(arr) {
				break
			}
		} else {
			ret = append(ret, arr[i])
			if len(ret) == len(arr) {
				break
			}
		}
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = ret[i]
	}
}

func largestValsFromLabels(values []int, labels []int, num_wanted int, use_limit int) int {
	var valueLabels [][]int
	labelMap := make(map[int][]int)
	for i := 0; i < len(values); i++ {
		if _, ok := labelMap[labels[i]]; !ok {
			labelMap[labels[i]] = []int{}
		}
		labelMap[labels[i]] = append(labelMap[labels[i]], values[i])
		valueLabels = append(valueLabels, []int{values[i], labels[i]})
	}
	sort.Slice(valueLabels, func(i, j int) bool {
		return valueLabels[i][0] > valueLabels[j][0]
	})

	used := make(map[int]int)
	var sum int
	var count int
	for i := 0; i < len(valueLabels); i++ {
		if used[valueLabels[i][1]] >= use_limit {
			continue
		}

		used[valueLabels[i][1]]++
		sum += valueLabels[i][0]
		count++
		if count == num_wanted {
			break
		}
	}

	return sum
}

type stack struct {
	location []int
	length   int
}

func shortestPathBinaryMatrix(grid [][]int) int {
	searched := make(map[int]map[int]bool)
	n := len(grid)
	for i := 0; i < n; i++ {
		searched[i] = make(map[int]bool)
	}

	stacks := []stack{{
		location: []int{0, 0},
		length:   1,
	}}
	searched[0][0] = true

	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}

	for len(stacks) > 0 {
		s := stacks[0]
		x := s.location[0]
		y := s.location[1]

		if x == n-1 && y == n-1 {
			return s.length
		}

		length := s.length + 1

		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}

				nx := x + i
				ny := y + j
				if nx >= 0 && nx < n && ny >= 0 && ny < n && grid[nx][ny] == 0 && !searched[nx][ny] {
					stacks = append(stacks, stack{
						location: []int{nx, ny},
						length:   length,
					})
					searched[nx][ny] = true
				}
			}
		}

		stacks = stacks[1:]
	}

	return -1
}

func shortestCommonSupersequence(str1 string, str2 string) string {
	dp := make([][]int, len(str1)+1)
	for i := 0; i <= len(str1); i++ {
		dp[i] = make([]int, len(str2)+1)
		for j := 0; j <= len(str2); j++ {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else if str1[i-1] == str2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				if dp[i-1][j] < dp[i][j-1] {
					dp[i][j] = 1 + dp[i-1][j]
				} else {
					dp[i][j] = 1 + dp[i][j-1]
				}
			}
		}
	}

	length := dp[len(str1)][len(str2)]
	var ret string
	i := len(str1)
	j := len(str2)
	for i > 0 && j > 0 {
		if str1[i-1] == str2[j-1] {
			ret = string(str1[i-1]) + ret
			i--
			j--
			length--
		} else if dp[i-1][j] >= dp[i][j-1] {
			ret = string(str2[j-1]) + ret
			j--
			length--
		} else {
			ret = string(str1[i-1]) + ret
			i--
			length--
		}
	}

	for i > 0 {
		ret = string(str1[i-1]) + ret
		i--
		length--
	}
	for j > 0 {
		ret = string(str2[j-1]) + ret
		j--
		length--
	}

	return ret
}
