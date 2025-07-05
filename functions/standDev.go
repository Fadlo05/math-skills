package functions

import "math"

func StandDev(nums []float64) int {
	variance := float64(Variance(nums))
	square := math.Sqrt(variance)
	standDev := math.Round(square)
	return int(standDev)
}