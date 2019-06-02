package main

import (
	"strings"
)

func gcdOfStrings(str1 string, str2 string) string {
	str1Len := len(str1)
	str2Len := len(str2)

	if str2Len > str1Len {
		str1, str2 = str2, str1
		str1Len, str2Len = str2Len, str1Len
	}

	for i := str2Len; i > 0; i-- {
		if str1Len%i != 0 || str2Len%i != 0 {
			continue
		}

		n1 := str1Len / i
		n2 := str2Len / i
		flag := true
		for j := 0; j < n1; j++ {
			if str1[i*j:i*(j+1)] != str2[0:i] {
				flag = false
				break
			}
			if j < n2 && str2[i*j:i*(j+1)] != str2[0:i] {
				flag = false
				break
			}
		}
		if flag {
			return str2[0:i]
		}
	}
	return ""
}

func maxEqualRowsAfterFlips(matrix [][]int) int {
	sum := make(map[string]int)
	for i := 0; i < len(matrix); i++ {
		var flipsZero []string
		var flipsOne []string
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != 0 {
				flipsZero = append(flipsZero, string(j))
			} else {
				flipsOne = append(flipsOne, string(j))
			}
		}
		sum[strings.Join(flipsZero, ",")]++
		sum[strings.Join(flipsOne, ",")]++
	}

	var max int
	for _, v := range sum {
		if v > max {
			max = v
		}
	}
	return max
}

func addNegabinary(arr1 []int, arr2 []int) []int {
	var ret []int
	var carry int
	for len(arr1) > 0 || len(arr2) > 0 || carry != 0 {
		if len(arr1) > 0 {
			carry += arr1[len(arr1)-1]
			arr1 = arr1[0 : len(arr1)-1]
		}
		if len(arr2) > 0 {
			carry += arr2[len(arr2)-1]
			arr2 = arr2[0 : len(arr2)-1]
		}
		ret = append([]int{carry & 1}, ret...)
		carry = -(carry >> 1)
	}
	for len(ret) > 1 && ret[0] == 0 {
		ret = ret[1:]
	}
	return ret
}

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	prefixSumMatrix := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		prefixSumMatrix[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if j == 0 {
				prefixSumMatrix[i][j] = matrix[i][j]
			} else {
				prefixSumMatrix[i][j] = prefixSumMatrix[i][j-1] + matrix[i][j]
			}
		}
	}

	var ret int
	for i := 0; i < len(matrix[0]); i++ {
		for j := i; j < len(matrix[0]); j++ {
			countMap := map[int]int{0: 1}
			var current int
			for k := 0; k < len(matrix); k++ {
				current += prefixSumMatrix[k][j]
				if i > 0 {
					current -= prefixSumMatrix[k][i-1]
				}
				ret += countMap[current-target]
				countMap[current]++
			}
		}
	}

	return ret
}
