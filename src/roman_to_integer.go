package main
//Bài 13
// https://leetcode.com/problems/roman-to-integer/description/
func romanToInt(s string) int {
	romanMap := map[rune]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}

	result, prev := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		curr := romanMap[rune(s[i])]
		if curr < prev {
			result -= curr
		} else {
			result += curr
		}
		prev = curr
	}

	return result
}

