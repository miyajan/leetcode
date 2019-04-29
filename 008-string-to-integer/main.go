package main

import (
	"bytes"
	"log"
	"math"
	"strconv"
)

func myAtoi(str string) int {
	var i int
	var minus bool
	buf := bytes.NewBuffer([]byte{})
	for i < len(str) {
		if str[i] == ' ' {
			if buf.Len() > 0 {
				break
			}
			i++
		} else if str[i] == '+' {
			if buf.Len() > 0 {
				break
			}
			if i == len(str)-1 || !isNumeric(str[i+1]) {
				return 0
			}
			i++
		} else if str[i] == '-' {
			if buf.Len() > 0 {
				break
			}
			if i == len(str)-1 || !isNumeric(str[i+1]) {
				return 0
			}
			minus = true
			i++
		} else if isNumeric(str[i]) {
			for str[i] == '0' && buf.Len() == 0 && i < len(str)-2 && isNumeric(str[i+1]) {
				i++
			}
			buf.WriteByte(str[i])
			i++
		} else {
			break
		}
	}

	if buf.Len() == 0 {
		return 0
	}

	s := buf.String()
	if len(s) > 10 {
		if minus {
			return math.MinInt32
		}
		return math.MaxInt32
	}

	ret, err := strconv.Atoi(s)
	if err != nil {
		log.Panicf("invalid strconv.Atoi: %s\n", s)
	}

	if minus {
		ret = -ret
	}

	if ret < math.MinInt32 {
		return math.MinInt32
	}
	if ret > math.MaxInt32 {
		return math.MaxInt32
	}
	return ret
}

func isNumeric(char byte) bool {
	return '0' <= char && char <= '9'
}
