package recursion

func sum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return nums[0] + sum(nums[1:])
}

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
