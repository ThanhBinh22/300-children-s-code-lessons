package main

import "math"
//https://leetcode.com/problems/sqrtx
// Bài 69

func mySqrt(x int) int {
    result := math.Sqrt(float64(x))
	return int(result)
}