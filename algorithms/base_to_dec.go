package algorithms

import (
	"math"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	const charset = "0123456789ABCDEF"

	var res float64
	for c, i := range Reverse(value) {
		for j, v := range charset {
			if i == v {
				res += float64(j) * math.Pow(float64(base), float64(c))
			}
		}
	}
	return int(res)
}
