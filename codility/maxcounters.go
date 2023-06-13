package codility

import "fmt"

// Optimized: O(M+N)
func MaxCounters(N int, A []int) []int {

	m := make(map[int]int)
	n := make(map[int]bool)

	max := 0

	ref := 0

	for i := 1; i <= N; i++ {
		m[i] = 0
	}

	for _, v := range A {

		if 1 <= v && v <= N {

			if ref != 0 && n[v] != true {
				m[v] = ref
				n[v] = true
			}
			m[v]++

			if m[v] > max {
				max = m[v]
			}
		}

		if v == N+1 {
			ref = max
			n = map[int]bool{}
		}

		fmt.Printf("\n %v M \n", m)
	}

	result := []int{}

	for i := 0; i < N; i++ {

		if m[i+1] < ref {
			m[i+1] = ref
		}
		result = append(result, m[i+1])
	}

	return result
}

// O(M*N)
func MaxCounters2(N int, A []int) []int {

	counters := make([]int, N)

	var bigger int

	for i, v := range A {
		if v >= 1 && v <= N {
			x := counters[v-1] + 1
			counters[v-1] = x

			if x > bigger {
				bigger = x
			}

		}

		if A[i] == N+1 {

			for i := 0; i < N; i++ {
				counters[i] = bigger
			}

		}
	}

	return counters

}
