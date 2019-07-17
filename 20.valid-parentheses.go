/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */
func isValid(s string) bool {
	stack := []rune{}

	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
		} else {
			if len(stack) < 1 {
				return false
			}
			if ch == ')' && stack[len(stack)-1] != '(' {
				return false
			}
			if ch == '}' && stack[len(stack)-1] != '{' {
				return false
			}
			if ch == ']' && stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

