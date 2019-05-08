package main

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		b := []byte(strs[i])
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		s := string(b)

		if _, ok := m[s]; !ok {
			m[s] = []string{strs[i]}
		} else {
			m[s] = append(m[s], strs[i])
		}
	}

	var ret [][]string
	for key := range m {
		ret = append(ret, m[key])
	}

	return ret
}
