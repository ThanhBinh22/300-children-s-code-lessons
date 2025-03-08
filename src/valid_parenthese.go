package main

// https://leetcode.com/problems/valid-parentheses/
// Bài 20
func isValid(s string) bool {
    stack := []rune{}
    brackets := map[rune]rune{')': '(', '}': '{', ']': '['}

    for _, char := range s {
        if char == '(' || char == '{' || char == '[' {
            stack = append(stack, char)
        } else {
            if len(stack) == 0 {
                return false
            }
            top := stack[len(stack)-1] 
            stack = stack[:len(stack)-1]
            if brackets[char] != top {
                return false
            }
        }
    }
    return len(stack) == 0
}
