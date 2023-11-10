/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
import "strconv"

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)

	n := len(str)

	for i := 0; i < n/2; i++ {
		if str[i] != str[n-i-1] {
			return false
		}
	}

	return true
}

// @lc code=end
