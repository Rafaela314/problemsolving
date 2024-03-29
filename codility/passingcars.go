package codility

/*

A non-empty array A consisting of N integers is given. The consecutive elements of array A
represent consecutive cars on a road.

Array A contains only 0s and/or 1s:

        0 represents a car traveling east,
        1 represents a car traveling west.

The goal is to count passing cars. We say that a pair of cars (P, Q), where 0 ≤ P < Q < N,
is passing when P is traveling to the east and Q is traveling to the west.

For example, consider array A such that:
  A[0] = 0
  A[1] = 1
  A[2] = 0
  A[3] = 1
  A[4] = 1

We have five pairs of passing cars: (0, 1), (0, 3), (0, 4), (2, 3), (2, 4).

Write a function:

    class Solution { public int solution(int[] A); }

that, given a non-empty array A of N integers, returns the number of pairs of passing cars.

The function should return −1 if the number of pairs of passing cars exceeds 1,000,000,000.

For example, given:
  A[0] = 0
  A[1] = 1
  A[2] = 0
  A[3] = 1
  A[4] = 1

the function should return 5, as explained above.

Write an efficient algorithm for the following assumptions:

        N is an integer within the range [1..100,000];
        each element of array A is an integer that can have one of the following values: 0, 1.

*/

func PassingCars(A []int) int {

	m := make(map[int][]int)

	LenA := len(A)

	m[A[0]] = []int{}

	for k, v := range A {

		if v == A[0] {
			m[A[0]] = append(m[A[0]], k)
		}
	}

	LenB := len(m[0])

	counter := 0

	for _, y := range m[0] {
		counter = counter + LenA - y - LenB
		LenB--

		if counter > 1000000000 {
			return -1
		}

	}

	return counter
}

/*
func PassingCars(A []int) int {

	m := make(map[int][]int)
	m[0] = []int{}
	m[1] = []int{}

	for k, v := range A {
		if v == 1 {
			m[1] = append(m[1], k)
			continue
		}
		m[0] = append(m[0], k)

	}

	var counter int

	for _, v := range m[0] {
		for i := 1; i < 5; i++ {


		for _, l := range m[1] {
			if v < l {
				counter++

				//fmt.Printf("\n %v v \n", v)
				//fmt.Printf("\n %v l \n", l)

				//fmt.Printf("\n %v final \n", counter)

				if counter > 1000000000 {
					return -1
				}
			}

		}

	}

	return counter

}*/

/*
func PassingCars(A []int) int {

	m := make(map[int][]int)
	m[0] = []int{}
	m[1] = []int{}

	for k, v := range A {
		if v == 1 {
			m[1] = append(m[1], k)
			continue
		}
		m[0] = append(m[0], k)

	}

	var counter int

	for _, v := range m[0] {
		for _, l := range m[1] {
			if v < l {
				counter++

				//fmt.Printf("\n %v v \n", v)
				//fmt.Printf("\n %v l \n", l)

				//fmt.Printf("\n %v final \n", counter)

				if counter > 1000000000 {
					return -1
				}
			}

		}

	}

	return counter

}*/

/*
func PassingCars(A []int) int {

	m := make(map[int]int)

	for k, v := range A {
		if v != 0 {
			m[k] = v
		}
	}

	var counter int

	for k, v := range A {

		if v == 0 {

			for x, _ := range m {
				if x > k {
					counter++
				}

				if counter > 1000000000 {
					return -1
				}

			}

		}

	}

	return counter

}
*/
/*time out


func Solution(A []int) int {
    // write your code in Go 1.4
    var first int

	var counter int

	for k, v := range A {
		if k == 0 {
			first = v
		}

		if v == first {
			for i := k; i < len(A); i++ {
				if A[i] != first {
					counter++
				}

				if counter > 1000000000 {
					return -1
				}
			}

		}

	}

	return counter

}*/
