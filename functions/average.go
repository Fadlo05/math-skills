package functions

func Average(nums []float64) float64 {
	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	avg := sum / float64(len(nums))
	return avg
}