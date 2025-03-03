package main
// https://leetcode.com/problems/two-sum/
//Bài 1

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if index, found := seen[complement]; found {
			return []int{index, i}
		}
		seen[num] = i
	}
	return nil
}
