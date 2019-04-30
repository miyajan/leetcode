package main

func intToRoman(num int) string {
	return toRoman(num/1000, "M", "", "") +
		toRoman((num%1000)/100, "C", "D", "M") +
		toRoman((num%100)/10, "X", "L", "C") +
		toRoman(num%10, "I", "V", "X")
}

func toRoman(num int, oneChar string, fiveChar string, tenChar string) string {
	switch num {
	case 0:
		return ""
	case 4:
		return oneChar + fiveChar
	case 5:
		return fiveChar
	case 9:
		return oneChar + tenChar
	default:
		var roman string
		if num > 5 {
			roman = fiveChar
			num -= 5
		}
		for num > 0 {
			roman = roman + oneChar
			num--
		}
		return roman
	}
}
