package codility

func Distinct(A []int) int {

	m := make(map[int]int)

	for _, v := range A {

		_, ok := m[v]
		if !ok {
			m[v] = 1
		}

	}
	return len(m)
}
