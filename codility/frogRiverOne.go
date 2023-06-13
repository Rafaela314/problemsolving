package codility

import "fmt"

func FrogRiverOne(X int, A []int) int {

	var result int

	m := make(map[int]bool) // [position]check

	for i := 1; i <= X; i++ {
		m[i] = false
	}

	checks := 0

	for k, v := range A {
		if m[v] != true {
			m[v] = true
			checks++

			if checks == X {
				result = k

				break
			}
		}
	}

	fmt.Printf("\n result : %v \n", result)
	return result

}

func FrogRiverOne2(X int, A []int) int {

	m := make(map[int]bool) // [position]check

	for i := 1; i <= X; i++ {
		m[i] = false
	}

	var counter int

	var result int

	for k, v := range A {

		if exists, ok := m[v]; ok {

			//already checked?
			if !exists {
				m[v] = true
				//	fmt.Println(m)
				counter++
			}

			if counter == X {

				result = k
				break
			} else {
				result = -1
			}
		}
	}
	fmt.Printf("\n result : %v \n", result)
	return result

}
