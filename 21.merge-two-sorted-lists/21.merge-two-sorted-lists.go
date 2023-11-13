/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	margedList := &ListNode{}
	currentNode := margedList

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			currentNode.Next = list1
			list1 = list1.Next
		} else {
			currentNode.Next = list2
			list2 = list2.Next
		}
		currentNode = currentNode.Next
	}

	if list1 == nil {
		currentNode.Next = list2
	} else {
		currentNode.Next = list1
	}

	return margedList.Next
}

// @lc code=end

