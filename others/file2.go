package others

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func File2(s string, t byte) []int {

	m := []int{}

	result := []int{}

	letter := string(t)

	str := strings.Split(s, "")

	for x, y := range str {
		if y == letter {
			m = append(m, x)
		}
	}
	sort.Ints(m)

	for i, v := range str {

		fmt.Printf("\n result %v \n", result)

		if v == letter {
			result = append(result, 0)
			continue
		}

		for _, vIdx := range m {
			ref := 0

			if i-vIdx == int(math.Abs(1)) {
				result = append(result, 0)
				continue
			}

			if int(math.Abs(float64(i-vIdx))) > ref {
				ref = int(math.Abs(float64(i - vIdx)))
			}

			result = append(result, ref)

		}

	}

	return result

}
