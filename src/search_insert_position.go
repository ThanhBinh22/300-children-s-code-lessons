package main

import "sort"

//Bài 35
// https://leetcode.com/problems/search-insert-position/description/

func searchInsert(nums []int, target int) int {
    nums = append(nums, target)
	sort.Ints(nums)
	index := 0
	for i := 0; i < len(nums); i++{
		if nums[i] == target{
			index = i
			break
		}
	}
	return index
}