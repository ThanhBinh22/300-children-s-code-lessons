package main

import (
	"strconv"
)

//Bài 9
//https://leetcode.com/problems/palindrome-number/description/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else {
		xStr := strconv.Itoa(x)
		for i := 0; i < len(xStr)/2; i++ {
			if xStr[i] != xStr[len(xStr)-1-i] {
				return false
			}
		}
	}

	return true
}
