package functions

import "math"

func Variance(nums []float64) int {
	avg := float64(Average(nums))
	var gap []float64
	for _, num := range nums {
		gap = append(gap, num-avg)
	}
	var square []float64
	for _, num := range gap {
		square = append(square, num*num)
	}
	sum := 0.0
	for _, num := range square {
		sum += num
	}
	variance := sum / float64(len(nums))
	return int(math.Round(variance))
}