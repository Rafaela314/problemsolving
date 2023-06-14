package codility

import (
	"sort"
)

func MissingInteger(A []int) int {

	sort.Ints(A)

	var hasPositives bool

	var firstP int

	var last int

	var next int

	for _, v := range A {

		if v > 0 {
			if !hasPositives {
				hasPositives = true
				firstP = v
				next = firstP + 1
				last = v
			}

			if firstP != 1 {
				return 1
			}

			if v != last {

				if v != next {
					return next
				}

				last = v
				next = v + 1
			}
		}
	}

	if !hasPositives {
		return 1
	}

	return A[len(A)-1] + 1

}
