/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
// O(n^2)
func twoSum(nums []int, target int) []int {
	// O(n)
	for i := 0; i < len(nums)-1; i++ {

		// O(n)
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if sum == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// @lc code=end

