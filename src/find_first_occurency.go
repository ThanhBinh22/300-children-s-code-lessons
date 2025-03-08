package main

import "strings"

// Bài 28
// https://leetcode.com/problems/implement-strstr/description/
func strStr(haystack string, needle string) int {
	if !strings.Contains(haystack, needle) {
		return -1
	} 
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}