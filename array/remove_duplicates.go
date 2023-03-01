package array

func removeDuplicates(nums []int) int {
	seen := make(map[int]bool)
	k := 0
	for _, val := range nums {
		if !seen[val] {
			seen[val] = true
			nums[k] = val
			k++
		}
	}
	return k
}
