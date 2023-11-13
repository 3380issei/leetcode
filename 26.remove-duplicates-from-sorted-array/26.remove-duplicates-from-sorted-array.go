/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {

	prev := nums[0]
	uniqueNum := 1

	for i := 1; i < len(nums); i++ {
		if prev != nums[i] {
			nums[uniqueNum] = nums[i]
			prev = nums[i]
			uniqueNum++
		}
	}

	return uniqueNum
}

// @lc code=end

