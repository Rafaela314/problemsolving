package codility

import (
	"fmt"
)

// O(N + M)
func GenomicRangeQuery(S string, P, Q []int) []int {

	result := make([]int, len(P))

	positions := map[string][]int{"A": {}, "C": {}, "G": {}}

	comp := map[string]int{"A": 1, "C": 2, "G": 3, "T": 4}

	for i, v := range S {
		if v == rune('A') {
			positions["A"] = append(positions["A"], i)

		}
		if v == rune('C') {
			positions["C"] = append(positions["C"], i)

		}
		if v == rune('G') {
			positions["G"] = append(positions["G"], i)

		}
	}

	fmt.Printf("\n POSITIONS:%v \n", positions)

	for i := 0; i < len(P); i++ {

		var found bool

		for _, a := range positions["A"] {

			if P[i] <= a && Q[i] >= a {

				found = true
				break
			}

		}

		if found {
			result[i] = comp["A"]
			continue
		}

		for _, c := range positions["C"] {

			if P[i] <= c && Q[i] >= c {

				found = true
				break
			}
		}

		if found {
			result[i] = comp["C"]
			continue
		}

		for _, g := range positions["G"] {

			if P[i] <= g && Q[i] >= g {
				found = true
				break
			}

		}

		if found {
			result[i] = comp["G"]
			continue
		}

		result[i] = comp["T"]

	}
	fmt.Printf("\n Result:%v \n", result)

	return result
}
