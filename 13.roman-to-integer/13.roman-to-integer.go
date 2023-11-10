/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {

	var current, previous, result int

	symbolInt := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := len(s) - 1; i >= 0; i-- {
		current = symbolInt[s[i]]

		if current < previous {
			result -= current
		} else {
			result += current
		}

		previous = current
	}

	return result

}

// @lc code=end

