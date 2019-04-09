package main

func lengthOfLongestSubstring(s string) int {
	var i, ret int
	dup := map[uint8]struct{}{}
	for j := 0; j < len(s); j++ {
		for _, ok := dup[s[j]]; ok; _, ok = dup[s[j]] {
			delete(dup, s[i])
			i++
		}
		dup[s[j]] = struct{}{}

		if j-i+1 > ret {
			ret = j - i + 1
		}
	}
	return ret
}
