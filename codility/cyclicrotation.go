package codility

import "fmt"

func CyclicRotation(a []int, k int) []int {
	r := make([]int, len(a))

	limit := len(a) // ex 5 : 0-4 - 13

	for i, v := range a {

		if i+k+1 > limit {
			reminder := (i + k) % limit
			r[reminder] = v

		} else {

			r[i+k] = v

		}
	}

	fmt.Printf("\n %v final \n", r)

	return r

}
