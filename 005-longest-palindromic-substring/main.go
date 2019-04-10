package main

func longestPalindrome(s string) string {
	var ret string
	for i := 0; i < len(s); i++ {
		s1 := getLongestPalindrome(s, i, i)
		if len(s1) > len(ret) {
			ret = s1
		}
		s2 := getLongestPalindrome(s, i, i+1)
		if len(s2) > len(ret) {
			ret = s2
		}
	}
	return ret
}

func getLongestPalindrome(s string, left int, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}
