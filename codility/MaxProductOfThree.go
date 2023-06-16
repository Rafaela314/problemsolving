package codility

import "fmt"

func MaxProductOfThree(A []int) int {

	var p1, p2, p3, n1, n2, m1, m2, m3 int

	// p : biggest positives
	// n smallest negatives
	// m biggest negatives

	if len(A) == 3 {
		return A[0] * A[1] * A[2]
	}

	var hasPositive bool

	for _, v := range A {

		if v >= 0 {

			hasPositive = true
			if v >= p1 {
				p3 = p2
				p2 = p1
				p1 = v
			} else if v >= p2 {
				p3 = p2
				p2 = v
			} else if v > p3 {
				p3 = v
			}
		}

		if v < 0 {
			if v <= n1 {
				n2 = n1
				n1 = v
			} else if v <= n2 {
				n2 = v
			}

			if v >= m1 || m1 == 0 {
				m3 = m2
				m2 = m1
				m1 = v
			} else if v >= m2 || m2 == 0 {
				m3 = m2
				m2 = v
			} else if v > m3 {
				m3 = v
			}

		}
		fmt.Printf(" \n p1: %v p2: %v p3: %v n1: %v n2: %v m1: %v m2: %v m3: %v\n", p1, p2, p3, n1, n2, m1, m2, m3)

	}

	if !hasPositive {
		return m1 * m2 * m3
	}

	a := p1 * p2 * p3
	b := p1 * n1 * n2

	if a > b {
		return a
	}

	return b

}
