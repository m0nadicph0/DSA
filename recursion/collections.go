package recursion

// Calculate the sum of an array of numbers.
func sum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return nums[0] + sum(nums[1:])
}

// Calculate the product of an array of numbers
func product(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	default:
		return nums[0] * product(nums[1:])
	}
}
