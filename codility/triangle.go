package codility

import "sort"

//A triplet (P, Q, R) is triangular if 0 ≤ P < Q < R < N

func Triangle(A []int) int { //10,5,8

	if len(A) == 0 {
		return 0
	}

	sort.Ints(A)

	for i := 0; i < len(A); i++ {
		a := A[i]
		b := A[i+1]
		c := A[i+2]

		if (a+b > c) && (b+c > a) && (a+c > b) {
			return 1
		}
	}
	return 0
}
