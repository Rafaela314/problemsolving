package codility

/*
A non-empty array A consisting of N integers is given. Array A represents numbers on a tape.

Any integer P, such that 0 < P < N, splits this tape into two non-empty parts: A[0], A[1], ..., A[P − 1] and A[P], A[P + 1], ..., A[N − 1].

The difference between the two parts is the value of: |(A[0] + A[1] + ... + A[P − 1]) − (A[P] + A[P + 1] + ... + A[N − 1])|

In other words, it is the absolute difference between the sum of the first part and the sum of the second part.

For example, consider array A such that:
  A[0] = 3
  A[1] = 1
  A[2] = 2
  A[3] = 4
  A[4] = 3

We can split this tape in four places:

        P = 1, difference = |3 − 10| = 7
        P = 2, difference = |4 − 9| = 5
        P = 3, difference = |6 − 7| = 1
        P = 4, difference = |10 − 3| = 7

Write a function:

    class Solution { public int solution(int[] A); }

that, given a non-empty array A of N integers, returns the minimal difference that can be achieved.

For example, given:
  A[0] = 3
  A[1] = 1
  A[2] = 2
  A[3] = 4
  A[4] = 3

the function should return 1, as explained above.*/

import (
	"fmt"
	"math"
)

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
