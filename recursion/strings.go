package recursion

// Reverse a string.
func reverse(s string) string {
	if len(s) == 0 {
		return ""
	}
	return string(s[len(s)-1]) + reverse(s[:len(s)-1])
}

// Check if a string is a palindrome.
func isPalindrome(s string) bool {
	return false
}
