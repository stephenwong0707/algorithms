package binarysearch

func findPeakElement(nums []int) int {
	low := 0
	high := len(nums) - 1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] > nums[low] && nums[mid] > nums[mid-1] {
			low = mid
		} else if nums[mid] > nums[high] && nums[mid] > nums[mid+1] {
			high = mid
		} else {
			low++
		}
	}
	return low
}
