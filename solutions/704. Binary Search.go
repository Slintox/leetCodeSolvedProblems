package solutions

func Search(nums []int, target int) int {
	return search(nums, target)
}

func search(nums []int, target int) int {
	var left, right, mid int
	left, right = 0, len(nums)-1

	for left <= right {
		mid = left + (right-left)/2

		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
