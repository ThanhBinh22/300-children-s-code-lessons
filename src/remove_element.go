package main

// Bài 27
// https://leetcode.com/problems/remove-element/description/
func removeElement(nums []int, val int) int {
	writeIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[writeIndex] = nums[i]
			writeIndex++
		}
	}
	return writeIndex
}