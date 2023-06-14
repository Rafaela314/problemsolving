package codility

func PassingCars(A []int) int {

	m := make(map[int][]int)

	LenA := len(A)

	m[A[0]] = []int{}

	for k, v := range A {

		if v == A[0] {
			m[A[0]] = append(m[A[0]], k)
		}
	}

	LenB := len(m[0])

	counter := 0

	for _, y := range m[0] {
		counter = counter + LenA - y - LenB
		LenB--

		if counter > 1000000000 {
			return -1
		}

	}

	return counter
}
