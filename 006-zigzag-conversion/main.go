package main

import "bytes"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	buf := bytes.NewBufferString("")
	for i := 0; i < numRows; i++ {
		j := i
		k := 0
		evenDiff := (numRows - i - 1) * 2
		oddDiff := i * 2
		for j < len(s) {
			buf.WriteByte(s[j])
			if (k % 2 == 0 && i != numRows-1) || i == 0 {
				j += evenDiff
			} else {
				j += oddDiff
			}
			k++
		}
	}
	return buf.String()
}
