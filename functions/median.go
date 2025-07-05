package functions

import (
	"math"
	"sort"
)

func Median(nums []float64) int {
	copyNums := make([]float64, len(nums))
	copy(copyNums, nums)
	sort.Float64s(copyNums)
	med := 0.0
	if len(copyNums)%2 != 0 {
		med = copyNums[len(copyNums)/2]
	} else {
		med = (copyNums[len(copyNums)/2] + copyNums[(len(copyNums)/2)-1]) / 2
	}
	return int(math.Round(med))
}