package stack

import "strings"

func ReverseString(s string) string {
	stack := NewSliceStack(len(s))
	for _, token := range s {
		stack.Push(string(token))
	}

	reversed := ""
	for !stack.IsEmpty() {
		tok, _ := stack.Pop()
		reversed += tok.(string)
	}
	return reversed
}

func IsParenthesesBalanced(expr string) bool {
	expr = strings.ReplaceAll(expr, " ", "")

	if len(expr) == 0 {
		return false
	}
	stack := NewSliceStack(len(expr))
	for _, token := range expr {
		if isOpen(string(token)) {
			_ = stack.Push(string(token))
		} else {
			if stack.IsEmpty() {
				return false
			}
			topValue, _ := stack.Peek()
			if matchingOpenParenthesis(string(token)) == topValue.(string) {
				_, _ = stack.Pop()
			}
		}
	}

	return stack.IsEmpty()
}

func matchingOpenParenthesis(closing string) string {
	switch closing {
	case ")":
		return "("
	case "]":
		return "["
	case "}":
		return "{"
	default:
		return ""
	}
}

func isOpen(token string) bool {
	switch token {
	case "(", "[", "{":
		return true
	default:
		return false
	}
}
