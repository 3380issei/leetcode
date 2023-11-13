/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {

	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	var stack []rune

	closingToOpening := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, c := range s {
		// c is opening
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
			continue
		}

		// c is closing
		if len(stack) != 0 && stack[len(stack)-1] == closingToOpening[c] {
			stack = stack[:len(stack)-1]
			continue
		} else {
			return false
		}
	}

	// is stack empty?
	if len(stack) == 0 {
		return true
	} else {
		return false
	}

}

// @lc code=end

