package codility

import "fmt"

func NumberOfDiscIntersections(A []int) int {

	count := 0

	for k, v := range A {
		max := k + v
		min := k - v

		for i := k + 1; i < len(A); i++ {
			max2 := i + A[i]
			min2 := i - A[i]

			if (max >= max2 && min <= max2) || (max2 >= max && min2 <= max) {

				fmt.Printf("\n COUNTTTT!!! MIN2  %d MAX2 %d MAX %d MIN %d\n", min2, max2, max, min)

				count++
				if count > 10000000 {
					return -1
				}
			}

		}

	}

	return count
}
