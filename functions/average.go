package functions

import "math"

func Average(nums []float64) int {
	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	avg := sum / float64(len(nums))
	return int(math.Round(avg))
}