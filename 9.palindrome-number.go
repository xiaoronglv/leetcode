/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */
func isPalindrome(x int) bool {
	// negative number
	if x < 0 {
		return false
	}

	// single digit number
	if x >=0 && x < 10 {
		return true
	}

	if x % 10 == 0 {
		return false
	}

	reverse := 0
	for x > reverse  {
		digit := x % 10
		x /= 10
		reverse = reverse*10 + digit
	}

	return x == reverse || x == reverse/10
}

