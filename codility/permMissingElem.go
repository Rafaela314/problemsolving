package codility

/*
An array A consisting of N different integers is given.
The array contains integers in the range [1..(N + 1)],
which means that exactly one element is missing.

Your goal is to find that missing element.

Write a function:

    class Solution { public int solution(int[] A); }

that, given an array A, returns the value of the missing element.

For example, given array A such that:
  A[0] = 2
  A[1] = 3
  A[2] = 1
  A[3] = 5

the function should return 4, as it is the missing element
*/

import (
	"fmt"
	"sort"
)

func PermMissingElem(A []int) int {

	sort.Ints(A)

	var first bool

	var result int

	var last int

	fmt.Printf("\n A %v  \n", A)

	if len(A) == 0 {
		result = 1
	}

	for i, v := range A {

		if !first {
			first = true
			if v != 1 {
				result = 1
				break
			}

			if len(A) == 1 {
				result = 2
				break
			}

			last = v
		} else if i <= len(A) {

			fmt.Printf("\n last %v  \n", last)

			delta := v - last

			fmt.Printf("\n delta %v  \n", delta)
			if delta > 1 {
				result = v - 1
				break
			}

			last = v
		}
	}

	if result == 0 {
		result = len(A) + 1
	}
	return result
}

/*
func main() {
	x := PermMissingElem([]int{})
	y := PermMissingElem([]int{1})
	z := PermMissingElem([]int{2})
	w := PermMissingElem([]int{1, 2, 3, 4})
	v := PermMissingElem([]int{1, 2, 3, 4, 6})

	fmt.Println(x, y, z, w, z, v)
}
*/
