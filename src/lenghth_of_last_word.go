package main

import "strings"

// Bài 58
// https://leetcode.com/problems/length-of-last-word/description/

func lengthOfLastWord(s string) int {
    strSplit := strings.Fields(s)
	lastWord := strSplit[len(strSplit) - 1]
	return len(lastWord)
}