package codility

import (
	"fmt"
	"sort"
)

func PermMissingElem(A []int) int {

	if len(A) == 0 {
		return 1
	}

	sort.Ints(A)

	for i := 1; i < len(A); i++ {
		if A[i] != A[i-1]+1 {
			fmt.Printf("\n A[i] %v  \n", A[i])
			return A[i] - 1
		}
	}

	if A[0] != 1 {
		return 1
	}

	return A[len(A)-1] + 1
}

func PermMissingElem2(A []int) int {

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
