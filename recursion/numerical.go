package recursion

// power of a number.
func power(base int, exponent int) int {
	if exponent == 0 {
		return 1
	}
	return base * power(base, exponent-1)
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

// calculate the nth triangular number.
func triangularNth(n int) int {
	if n == 1 {
		return 1
	}
	return n + triangularNth(n-1)
}

// calculate the sum of the first n natural numbers.
func sumNatural(n int) int {
	if n <= 1 {
		return n
	}
	return n + sumNatural(n-1)
}

// calculate the sum of the first n even numbers.
func sumOfFirstNEven(n int) int {
	if n == 0 {
		return 0
	}
	return 2*n + sumOfFirstNEven(n-1)
}

// calculate the sum of the first n odd numbers.
func sumOfFirstNOdd(n int) int {
	if n <= 1 {
		return n
	}
	return 2*n - 1 + sumOfFirstNOdd(n-1)
}

// iterate the numbers from 1 to n.
func forEachUptoN(n int, fn func(int)) {
	if n == 0 {
		return
	}
	forEachUptoN(n-1, fn)
	fn(n)
}

// iterate the numbers from n to 1.
func forEachFromNDownTo1(n int, fn func(int)) {
	if n == 0 {
		return
	}
	fn(n)
	forEachFromNDownTo1(n-1, fn)
}

// calculate the product of the first n natural numbers.
func productNatural(n int) int {
	if n <= 1 {
		return 1
	}
	return n * productNatural(n-1)
}

// calculate the product of the first n even numbers.
func productEvenNatural(n int) int {
	if n == 0 {
		return 1
	}
	return (2 * n) * productEvenNatural(n-1)
}

// calculate the product of the first n odd numbers.
func productOddNatural(n int) int {
	if n <= 1 {
		return 1
	}
	return (2*n - 1) * productOddNatural(n-1)
}

// calculate the sum of the digits of a binary number.
func sumDigitsBinary(n int) int {
	if n == 0 {
		return 0
	}
	return (n % 10) + sumDigitsBinary(n/10)
}
