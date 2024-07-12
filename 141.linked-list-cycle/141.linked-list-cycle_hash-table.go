/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {

	visitedNodes := make(map[*ListNode]bool)
	currentNode := head

	for currentNode != nil {
		if visitedNodes[currentNode] {
			return true
		}

		visitedNodes[currentNode] = true
		currentNode = currentNode.Next
	}

	return false
}

// @lc code=end

