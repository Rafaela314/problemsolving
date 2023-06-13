package codility

import (
	"fmt"
	"math"
)

func TapeEquilibrium(A []int) int {

	total := 0

	for _, v := range A {
		total += v
	}

	last := A[0]

	min := int(math.Abs(float64(last - (total - last))))

	for i := 1; i < len(A)-1; i++ {

		last += A[i]

		diff := math.Abs(float64(last - (total - last)))

		if diff < float64(min) {
			min = int(diff)

		}

		fmt.Printf("\n last %v min : %v, NUM %v, diff %v,  \n", last, min, A[i], diff)

	}

	return min
}

/*
func TapeEquilibrium(A []int) int {

	m := make(map[int]int) // [P]sum

	tt := A[0]
	ac := 0

	for i := 1; i < len(A); i++ {
		ac = ac + A[i-1]
		m[i] = ac

		tt = tt + A[i]
	}

	var first bool

	var min int

	for k, v := range m {

		if k == 0 {

		} else {

			delta := math.Abs(float64((tt - v) - v))
			if int(delta) < min || !first {
				min = int(delta)
				first = true
			}

			fmt.Printf("\n Delta : %v \n", delta)

		}

	}

	return min
}
*/
