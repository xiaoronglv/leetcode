/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */
func reverse(x int) int {
	var origin, reverse, max, min int
	origin = x
	reverse = 0
	max = math.MaxInt32
	min = math.MinInt32
	
	for origin != 0 {
		pop := origin % 10
		origin /= 10
	
		// positive number
		if reverse > max/10 || (reverse == max/10 && pop > 7) {
			return 0
		}
	
		// negative number
		if reverse < min/10 || (reverse == min/10 && pop < -8) {
			return 0
		}
		reverse = reverse*10 + pop
	}
	
	return reverse
}





