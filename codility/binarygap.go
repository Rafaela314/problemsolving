package codility

import (
	"fmt"
	"strings"
)

func SolutionIteration(N int) int {
	//find gaps, return 0 if there is no gap
	binary := fmt.Sprintf("%b", N)

	count := 0
	bigger := 0

	for _, v := range binary {
		if v == rune('1') {
			if count > bigger {
				bigger = count
			}
			count = 0
		} else {
			count++
		}
	}

	return bigger
}

func SolutionIteration2(N int) int {
	//find gaps, return 0 if there is no gap
	binary := fmt.Sprintf("%b", N)

	var first bool

	//split string into letters
	s := strings.Split(binary, "")

	var big int
	var last int

	for k, v := range s {
		if v == "1" {
			if !first {
				first = true
				last = k
			} else {
				delta := k - last - 1
				last = k
				if big < delta && delta > 0 {
					big = delta
				}
			}

		}
	}

	// return longest
	return big
}
