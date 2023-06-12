package codility

import (
	"fmt"
	"sort"
)

// optimized
func OddOccurrencesInArray(A []int) int {

	sort.Ints(A)

	mark := A[0]
	count := 1

	for i := 1; i < len(A); i++ {
		fmt.Printf("\n mark %v \n", mark)
		fmt.Printf("\n count %v \n", count)

		if A[i] == mark {
			count++
		} else {
			if count%2 != 0 {
				return mark
			}
			count = 1
			mark = A[i]
		}
	}

	return mark
}

func OddOccurrencesInArray2(A []int) int {

	var result int

	m := make(map[int]int) // [number]count

	for _, v := range A {
		//exists
		value, ok := m[v]
		if ok {
			m[v] = value + 1
		} else {
			m[v] = 1
		}
	}

	for key, vv := range m {
		if vv%2 != 0 {
			result = key

			break
		}
	}

	fmt.Printf("\n M %v \n", m)

	return result
}
