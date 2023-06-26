package codility

import (
	"fmt"
	"sort"
)

func Assembly(H []int, X int, Y int) int {

	sort.Ints(H)

	stop := 0
	sum := 0
	count := 0

	for k, v := range H {

		sum = sum + v

		if sum > X+Y {
			stop = k - 1
			break
		}

		if sum == X+Y {
			stop = k
			break
		}
	}

	var l1 int
	var l2 int

	if X > Y {
		l1 = X
		l2 = Y
	} else {

		l1 = Y
		l2 = X
	}

	fmt.Printf("\n stop: %v l1 :%v l2: %v \n", stop, l1, l2)

	for i := stop; i != -1; i-- {

		fmt.Printf("\n stop: %v i :%v x: %v y: %v\n", stop, i, X, Y)

		if l1 >= H[i] {
			l1 = l1 - H[i]
			count++
			continue
		}

		if l2 >= H[i] {
			l2 = l2 - H[i]
			count++

		}
	}

	//fmt.Printf("\n %v sum %v stop\n", sum, stop)

	//fmt.Printf("\n %v count \n", count)

	return count
}
