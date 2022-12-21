// Package utils provides general utility functions
package utils

import "math"

// RoundFloat rounds a float value using a provided precision.
func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
