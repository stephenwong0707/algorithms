package binarysearch

func searchRange(nums []int, target int) []int {
	first := searchFirst(nums, target)
	last := searchLast(nums, target)
	return []int{first, last}
}

func searchFirst(nums []int, target int) int {
	idx := -1
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			idx = mid
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return idx
}

func searchLast(nums []int, target int) int {
	idx := -1
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			idx = mid
			low = mid + 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return idx
}
