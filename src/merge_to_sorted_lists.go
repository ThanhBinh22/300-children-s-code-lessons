package main

// Bài 21
// https://leetcode.com/problems/merge-two-sorted-lists/description/
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	currentList := list1
	if list1.Val > list2.Val {
		currentList = list2
		list2 = list2.Next
	} else {
		list1 = list1.Next
	}
	head := currentList
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			currentList.Next = list2
			list2 = list2.Next
		} else {
			currentList.Next = list1
			list1 = list1.Next
		}
		currentList = currentList.Next
	}
	if list1 != nil {
		currentList.Next = list1
	}
	if list2 != nil {
		currentList.Next = list2
	}
	return head
}
