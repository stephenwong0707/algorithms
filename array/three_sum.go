package array

import "sort"

type Triplet struct {
	x int
	y int
	z int
}

func threeSum(nums []int) [][]int {
	tripletMap := map[Triplet]struct{}{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		low := i + 1
		high := len(nums) - 1
		for low < high {
			sum := nums[i] + nums[low] + nums[high]
			if sum == 0 {
				triplet := Triplet{nums[i], nums[low], nums[high]}
				if _, ok := tripletMap[triplet]; !ok {
					tripletMap[triplet] = struct{}{}
				}
				low++
			} else if sum > 0 {
				high--
			} else if sum < 0 {
				low++
			}

		}
	}
	var result [][]int
	for triplet := range tripletMap {
		result = append(result, []int{triplet.x, triplet.y, triplet.z})
	}
	return result
}
