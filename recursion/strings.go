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

	switch len(s) {
	case 0:
		return true
	case 1:
		return true
	case 2:
		return s[0] == s[1]
	default:
		return s[0] == s[len(s)-1] && isPalindrome(s[1:len(s)-1])
	}
}
