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
	return n%10 + digitSum(n/10)
}

// number of digits in a number.
func numDigits(n int) int {
	if n < 10 {
		return 1
	}
	return 1 + numDigits(n/10)
}
