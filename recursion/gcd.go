package recursion

//1. Take the two numbers as input.
//2. Check if one of the numbers is zero, in which case the GCD is the other number.
//3. If both numbers are non-zero, recursively call the GCD function with the smaller number and the difference of the two numbers.
//4. Repeat steps 2-3 until one of the numbers becomes zero, at which point we return the other number as the GCD.
func gcd(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a > b {
		return gcd(b, a-b)
	} else {
		return gcd(a, b-a)
	}
}

func lcm(a int, b int) int {
	return (a * b) / gcd(a, b)
}
