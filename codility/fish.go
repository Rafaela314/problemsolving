package codility

import "fmt"

func Fish(A []int, B []int) int {

	var alive int
	biggest := 0
	temp := []int{}

	for k, v := range B {

		if v == 0 {
			if len(temp) == 0 {
				alive++

				continue
			}

			if A[k] > biggest {
				alive = alive + 1 - len(temp)
				temp = []int{}
				biggest = 0

				continue
			} else {
				count := 0
				alive++
				for i := len(temp); i > 0; i-- {
					if A[k] > temp[i-1] {
						count++

						continue
					} else {

						alive--
						break

					}

				}

				temp = temp[:len(temp)-count]
				alive = alive - count
			}

		} else {

			temp = append(temp, A[k])
			if A[k] > biggest {
				biggest = A[k]
			}
			alive++

		}

	}
	fmt.Printf("\n K: %v ALIVE: %v, TEMP: %v, BIGGEST: %v \n", k, alive, temp, biggest)

	return alive

}
