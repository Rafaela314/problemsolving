package codility

import (
	"fmt"
	"sort"
)

//O(N*log(N))

func Triangle(A []int) int { //10,5,8

	if len(A) < 3 {
		return 0
	}

	sort.Ints(A)

	fmt.Printf("\n %v A \n", A)

	for i := 0; i < len(A)-2; i++ {
		a := A[i]
		b := A[i+1]
		c := A[i+2]

		if (a+b > c) && (b+c > a) && (a+c > b) {
			return 1
		}
	}
	return 0
}
