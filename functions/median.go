package functions

import (
	"sort"
)

func Median(nums []float64) float64 {
	sort.Float64s(nums)
	med := 0.0
	if len(nums)%2 != 0 {
		med = nums[len(nums)/2]
	} else {
		med = (nums[len(nums)/2] + nums[(len(nums)/2)-1]) / 2
	}
	return med
}
