/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// @lc code=start
func lengthOfLastWord(s string) int {

	length := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && length == 0 {
			continue
		} else if s[i] != ' ' {
			length++
			continue
		} else {
			break
		}
	}

	return length
}

// @lc code=end

