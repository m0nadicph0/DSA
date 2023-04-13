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

//Minimum element in an array.
func minimum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return min(nums[0], minimum(nums[1:]))
}

//Maximum element in an array.
func maximum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(nums[0], maximum(nums[1:]))
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// Check if an array is sorted in ascending order
func isSortedAscending(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	if nums[0] > nums[1] {
		return false
	}

	return isSortedAscending(nums[1:])
}

// Check if an array is sorted in descending order
func isSortedDescending(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	if nums[0] < nums[1] {
		return false
	}

	return isSortedDescending(nums[1:])
}

// Search for an element in an array
func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	if nums[0] == target {
		return true
	}
	return search(nums[1:], target)
}
