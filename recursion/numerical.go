package recursion

// power of a number.
func power(n int) int {
	return 0
}

// sum of digits of a number.
func digitSum(n int) int {
	if n < 10 {
		return n
	}
	lastDigit := n % 10
	remaining := n / 10
	return lastDigit + digitSum(remaining)
}

// number of digits in a number.
func numDigits(n int) int {
	return 0
}
