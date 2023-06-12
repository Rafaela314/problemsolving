package codility

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
