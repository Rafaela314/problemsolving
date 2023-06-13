package codility

import (
	"sort"
)

func PermCheck(A []int) int {

	counter := 1

	sort.Ints(A)

	for _, v := range A {

		if v != counter {
			return 0
		}

		counter++

	}

	return 1

}
