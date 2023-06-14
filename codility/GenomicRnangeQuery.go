package codility

import (
	"fmt"
	"strings"
)

func GenomicRangeQuery(S string, P, Q []int) []int {

	result := []int{}

	fmt.Printf("\n S@ %v \n", string(S[2]))

	for i := 0; i < len(P); i++ {

		if P[i] == Q[i] {

			result = append(result, findMatch(string(S[P[i]])))
			break
		}

		k := S[P[i]:Q[i]]

		fmt.Printf("\n k %v \n", k)

		if strings.Contains(k, "A") {
			result = append(result, 1)
			continue
		}

		if strings.Contains(k, "C") {
			result = append(result, 2)
			continue
		}

		if strings.Contains(k, "G") {
			result = append(result, 3)
			continue
		}
		if strings.Contains(k, "T") {
			result = append(result, 4)
		}
		fmt.Printf("\n result %v \n", result)
	}

	return result
}

func findMatch(a string) int {
	switch a {
	case "A":
		return 1
	case "C":
		return 2
	case "G":
		return 3
	default:
		return 4
	}
}

func GenomicRangeQuery2(S string, P, Q []int) []int {

	result := []int{}

	fmt.Printf("\n S@ %v \n", string(S[2]))

	for i := 0; i < len(P); i++ {

		if P[i] == Q[i] {

			result = append(result, findMatch(string(S[P[i]])))
			break
		}

		k := S[P[i]:Q[i]]

		fmt.Printf("\n k %v \n", k)

		if strings.Contains(k, "A") {
			result = append(result, 1)
			continue
		}

		if strings.Contains(k, "C") {
			result = append(result, 2)
			continue
		}

		if strings.Contains(k, "G") {
			result = append(result, 3)
			continue
		}
		if strings.Contains(k, "T") {
			result = append(result, 4)
		}
		fmt.Printf("\n result %v \n", result)
	}

	return result
}
