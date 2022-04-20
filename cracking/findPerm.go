package cracking

import (
	"sort"
	"strings"
)

/* Given a smaller string s and a bigger string b, design an algorith to find all permutaitions of the shorter within the longer one.
Print the location of each permutation. */

//Binary Gap

func FindPerm(A, B string) int {

	// keep count
	counter := 0

	//organize B
	b := strings.Split(B, "")
	sort.Strings(b)

	a := strings.Split(A, "")

	// make windows with len(b)
	start := 0
	end := len(b)

	var isPerm bool

	for j := 0; j < len(a)-len(b)+1; j++ {

		window := make([]string, len(b))

		_ = copy(window, a[start:end])

		sort.Strings(window)

		for i, v := range window {

			isPerm = true

			if v != b[i] {

				isPerm = false
				// optimize
				break
			}
		}

		if isPerm {

			counter++
			isPerm = false
		}

		start++
		end++
	}

	return counter
}
