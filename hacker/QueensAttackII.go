package hacker

import (
	"fmt"
	"sync"
)

func QueensAttackII(n, k, r_q, c_q int32, obstacles [][]int32) int32 {
	lookup := make(map[string]int)

	for _, v := range obstacles {
		key := fmt.Sprintf("%d,%d", v[0], v[1])
		lookup[key] = 1
	}
	var wg sync.WaitGroup

	var sumRight int32
	var sumLeft int32
	var sumUp int32
	var sumDown int32
	var sumDiagonalTopRight int32
	var sumDiagonalTopLeft int32
	var sumDiagonalBottomLeft int32
	var sumDiagonalBottomRight int32

	wg.Add(1)
	go func(r, c int32) {

		// to the right
		for i := c + 1; i <= n; i++ {
			key := fmt.Sprintf("%d,%d", r, i)
			if _, ok := lookup[key]; ok {
				break
			}
			sumRight++
		}
		defer wg.Done()
	}(r_q, c_q)

	wg.Add(1)
	go func(r, c int32) {

		// to the left
		for i := c - 1; i >= 1; i-- {
			key := fmt.Sprintf("%d,%d", r, i)
			if _, ok := lookup[key]; ok {
				break
			}
			sumLeft++
		}
		defer wg.Done()

	}(r_q, c_q)

	wg.Add(1)
	go func(r, c int32) {

		// down
		for i := r - 1; i >= 1; i-- {
			key := fmt.Sprintf("%d,%d", i, c)
			if _, ok := lookup[key]; ok {
				break
			}
			sumDown++
		}

		defer wg.Done()
	}(r_q, c_q)

	wg.Add(1)
	go func(r, c int32) {

		// up
		for i := r + 1; i <= n; i++ {
			key := fmt.Sprintf("%d,%d", i, c)
			if _, ok := lookup[key]; ok {
				break
			}
			sumUp++
		}
		defer wg.Done()

	}(r_q, c_q)
	wg.Add(1)
	go func(r, c int32) {

		//d1 - top left
		c1 := c - 1
		r1 := r + 1
		for r1 <= n && c1 >= 1 {
			key := fmt.Sprintf("%d,%d", r1, c1)
			if _, ok := lookup[key]; ok {
				break
			}
			c1--
			r1++
			sumDiagonalTopLeft++
		}
		defer wg.Done()
	}(r_q, c_q)
	wg.Add(1)
	go func(r, c int32) {

		//d2 - bottom left
		c2 := c - 1
		r2 := r - 1
		for r2 >= 1 && c2 >= 1 {
			key := fmt.Sprintf("%d,%d", r2, c2)
			if _, ok := lookup[key]; ok {
				break
			}
			c2--
			r2--
			sumDiagonalBottomLeft++
		}

		defer wg.Done()
	}(r_q, c_q)

	wg.Add(1)
	go func(r, c int32) {

		//d3 - bottom right
		c3 := c + 1
		r3 := r - 1
		for r3 >= 1 && c3 <= n {
			key := fmt.Sprintf("%d,%d", r3, c3)
			if _, ok := lookup[key]; ok {
				break
			}

			c3++
			r3--
			sumDiagonalBottomRight++
		}

		defer wg.Done()
	}(r_q, c_q)

	wg.Add(1)
	go func(r, c int32) {

		//d4 - top right
		c4 := c + 1
		r4 := r + 1
		for r4 <= n && c4 <= n {
			key := fmt.Sprintf("%d,%d", r4, c4)
			fmt.Println(key)
			if _, ok := lookup[key]; ok {
				break
			}
			c4++
			r4++
			sumDiagonalTopRight++
		}

		defer wg.Done()
	}(r_q, c_q)

	wg.Wait()

	total := sumRight + sumLeft + sumDown + sumUp + sumDiagonalTopLeft + sumDiagonalTopRight + sumDiagonalBottomRight + sumDiagonalBottomLeft

	return total

}
