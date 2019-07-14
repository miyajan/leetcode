package main

import (
	"math"
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	arr2Map := make(map[int]bool)
	for i := 0; i < len(arr2); i++ {
		arr2Map[arr2[i]] = true
	}

	arr1Map := make(map[int]int)
	for i := 0; i < len(arr1); i++ {
		arr1Map[arr1[i]]++
	}

	var ret []int
	for i := 0; i < len(arr2); i++ {
		count := arr1Map[arr2[i]]
		for j := 0; j < count; j++ {
			ret = append(ret, arr2[i])
		}
	}

	sort.Ints(arr1)
	for i := 0; i < len(arr1); i++ {
		if _, ok := arr2Map[arr1[i]]; !ok {
			ret = append(ret, arr1[i])
		}
	}

	return ret
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	log2n := 11
	parents := make([][]*TreeNode, 1001)
	for i := 0; i < 1001; i++ {
		parents[i] = make([]*TreeNode, log2n)
	}
	depths := make([][]*TreeNode, 1001)
	for i := 0; i < 1001; i++ {
		depths[i] = []*TreeNode{}
	}
	var dfs func(node *TreeNode, parent *TreeNode, depth int)
	dfs = func(node *TreeNode, parent *TreeNode, depth int) {
		parents[node.Val][0] = parent
		depths[depth] = append(depths[depth], node)

		if node.Left != nil {
			dfs(node.Left, node, depth+1)
		}
		if node.Right != nil {
			dfs(node.Right, node, depth+1)
		}
	}
	dfs(root, nil, 0)

	for i := 1; i < log2n; i++ {
		for j := 0; j < len(parents); j++ {
			if parents[j][i-1] == nil {
				continue
			}
			parents[j][i] = parents[parents[j][i-1].Val][i-1]
		}
	}

	var deepestLeaves []*TreeNode
	for i := 0; i < len(depths); i++ {
		if len(depths[i]) == 0 {
			break
		}
		deepestLeaves = depths[i]
	}

	if len(deepestLeaves) == 1 {
		return deepestLeaves[0]
	}

	for i := log2n - 1; i >= 0; i-- {
		a := parents[deepestLeaves[0].Val][i]
		flag := false
		for j := 1; j < len(deepestLeaves); j++ {
			if parents[deepestLeaves[j].Val][i] != a {
				flag = true
			}
		}

		if flag {
			for j := 0; j < len(deepestLeaves); j++ {
				deepestLeaves[j] = parents[deepestLeaves[j].Val][i]
			}
		}
	}
	return parents[deepestLeaves[0].Val][0]
}

func longestWPI(hours []int) int {
	prefixTiringDaySum := make([]int, len(hours)+1)
	prefixNonTiringDaySum := make([]int, len(hours)+1)
	prefixTiringDaySum[0] = 0
	prefixNonTiringDaySum[0] = 0
	for i := 0; i < len(hours); i++ {
		if hours[i] > 8 {
			prefixTiringDaySum[i+1] = prefixTiringDaySum[i] + 1
			prefixNonTiringDaySum[i+1] = prefixNonTiringDaySum[i]
		} else {
			prefixTiringDaySum[i+1] = prefixTiringDaySum[i]
			prefixNonTiringDaySum[i+1] = prefixNonTiringDaySum[i] + 1
		}
	}

	longest := 0
	for i := 0; i < len(hours); i++ {
		for j := i + 1; j <= len(hours); j++ {
			tiringDaySum := prefixTiringDaySum[j] - prefixTiringDaySum[i]
			nonTiringDaySum := prefixNonTiringDaySum[j] - prefixNonTiringDaySum[i]
			if tiringDaySum > nonTiringDaySum && tiringDaySum+nonTiringDaySum > longest {
				longest = tiringDaySum + nonTiringDaySum
			}
		}
	}

	return longest
}

func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	skillNum := len(req_skills)
	skillIndexMap := make(map[string]int)
	for i := 0; i < len(req_skills); i++ {
		skillIndexMap[req_skills[i]] = i
	}

	dp := make(map[int][]int)
	dp[0] = []int{}
	for i := 0; i < len(people); i++ {
		p := people[i]
		pSkill := 0
		for j := 0; j < len(p); j++ {
			skill := p[j]
			pSkill |= 1 << uint(skillIndexMap[skill])
		}
		for skillSet, team := range dp {
			totalSkillSet := pSkill | skillSet
			if dp[totalSkillSet] == nil || len(team)+1 < len(dp[totalSkillSet]) {
				dp[totalSkillSet] = append([]int{i}, team...)
			}
		}
	}

	return dp[int(math.Pow(2, float64(skillNum)))-1]
}
