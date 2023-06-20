package codility

import "fmt"

func Pow2(A []int) int {

	m := make(map[int]bool)

	var counter int

	for _, v := range A {

		_, ok := m[v]
		if !ok {
			m[v] = false
		}
	}

	counter = len(m)

	min := 0

	c := 0

	days := 0

	goal := 0

	for c <= (len(A) - len(m)) {

		days = 0
		goal = 0

		for j := range m {
			m[j] = false
		}

		for i := c; i < len(A); i++ {

			days = days + 1

			fmt.Printf("\n DAY : %v \n", days)

			if m[A[i]] == false {
				m[A[i]] = true
				goal = goal + 1
				fmt.Printf("\n GOAL : %v \n", goal)
			}

			if goal == counter {

				//fmt.Printf("\n DAY : %v Min %v\n", days, min)

				if min == 0 {
					min = days
				} else {
					if days < min {
						min = days
					}

				}
				break
			}
		}

		c++
	}

	fmt.Printf("\n MIN RESULT : %v \n", min)

	return min
}

/*
func find(a []int, i int) {
	for i := 1; i < 5; i++ {
		sum += i
	}

}*/
