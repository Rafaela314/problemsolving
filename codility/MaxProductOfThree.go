package codility

import (
	"sort"
)

/*
A non-empty array A consisting of N integers is given. The product of triplet (P, Q, R) equates to A[P] * A[Q] * A[R] (0 ≤ P < Q < R < N).

For example, array A such that:

  A[0] = -3
  A[1] = 1
  A[2] = 2
  A[3] = -2
  A[4] = 5
  A[5] = 6
contains the following example triplets:

(0, 1, 2), product is −3 * 1 * 2 = −6
(1, 2, 4), product is 1 * 2 * 5 = 10
(2, 4, 5), product is 2 * 5 * 6 = 60
Your goal is to find the maximal product of any triplet.

Write a function:

func Solution(A []int) int

that, given a non-empty array A, returns the value of the maximal product of any triplet.

For example, given array A such that:

  A[0] = -3
  A[1] = 1
  A[2] = 2
  A[3] = -2
  A[4] = 5
  A[5] = 6
the function should return 60, as the product of triplet (2, 4, 5) is maximal.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [3..100,000];
each element of array A is an integer within the range [−1,000..1,000].
*/

func MaxProductOfThree(A []int) int {

	sort.Ints(A)

	x := A[len(A)-1] * A[len(A)-2] * A[len(A)-3]
	y := A[len(A)-3] * A[0] * A[1]

	if x > y {
		return x
	}

	return y

}

/*
func MaxProductOfThree(A []int) int {

	sort.Ints(A)

	fmt.Printf(" \n A[]: %v \n", A[(len(A)-3):])

	result := A[(len(A)-1)] * A[(len(A)-2)] * A[(len(A)-3)]

	fmt.Printf(" \n result: %d \n", result)

	return result
}*/
/*
func MaxProductOfThree(A []int) int {

	m := make(map[int]int, len(A))

	for k, v := range A {
		m[v] = k
	}

	counter := 0
	a := 0
	b := 0
	c := 0

	sort.Ints(A)
	fmt.Printf(" \n A: %v \n", A)
	fmt.Printf(" \n M: %v \n", m)

	for i := len(A) - 1; i > 0; i-- {

		if m[A[i]] > m[A[i-1]] {
			if counter == 0 {
				a = A[i]

			}
			if counter == 1 {
				b = A[i]

			}
			if counter == 2 {
				c = A[i]

			}
			counter++

			if counter == 3 {
				break
			}
		}

	}

	fmt.Printf(" \n FINAL a: %d b %d c %d \n", a, b, c)
	return a * b * c
}
*/
/*func MaxProductOfThree(A []int) int {

	var result int

	first := A[0]
	firstIdx := 0
	second := 0
	secondIdx := 0
	third := 0
	thirdIdx := 0

	for i := 1; i < len(A); i++ {

		if A[i] > first && i > firstIdx {
			third = second
			thirdIdx = secondIdx
			second = first
			secondIdx = firstIdx
			first = A[i]
			firstIdx = i

			if (first * second * third) > result {
				result = first * second * third
			}
		}

		if first > A[i] && A[i] > second && i < firstIdx {
			third = second
			thirdIdx = secondIdx
			second = A[i]
			if (first * second * third) > result {
				result = first * second * third
			}
		}

		if second > A[i] && A[i] > thirdIdx && i < secondIdx {
			third = A[i]
			if (first * second * third) > result {
				result = first * second * third
			}
		}
		fmt.Printf(" \n first: %d second: %d third: %d \n", first, second, third)
		fmt.Printf(" \n firstIdx: %d secondIdx: %d thirdIdx: %d \n", firstIdx, secondIdx, thirdIdx)
	}
	return first * second * third
}*/
