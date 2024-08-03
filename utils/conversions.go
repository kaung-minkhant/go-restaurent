package utils

import "math"

func ToPointer[T any](value T) *T {
	return &value
}

func ToFixed(value float32, precision uint) float32 {
	multiplier := math.Pow(10, float64(precision))
	output := math.Round(float64(value) * multiplier)
	return float32(output / multiplier)
}
