/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Find the Index of the First Occurrence in a String
 */

// @lc code=start
func strStr(haystack string, needle string) int {

	if len(haystack) < len(needle) {
		return -1
	}

	for i, c := range haystack {

		if len(haystack) < i+len(needle) {
			break
		}

		if c == rune(needle[0]) {

			if needle == haystack[i:i+len(needle)] {
				return i
			}
		}
	}

	return -1
}

// @lc code=end

