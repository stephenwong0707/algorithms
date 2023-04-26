package array

func numSubarrayProductLessThanK(nums []int, k int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		product := 1
		for j := i; j < len(nums); j++ {
			product *= nums[j]
			if product < k {
				result++
			} else {
				j = len(nums)
			}
		}
	}
	return result
}
