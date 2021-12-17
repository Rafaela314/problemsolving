package main

/*

Write a function:

func Solution(A []int) int

that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.

For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

Given A = [1, 2, 3], the function should return 4.

Given A = [−1, −3], the function should return 1.
*/

import (
	"sort"
)

func firstP(a int) {}

func MissingInteger(A []int) int {

	// write your code in Go 1.4

	sort.Ints(A)

	var hasPositives bool
	var firstP int

	var last int

	var next int

	for _, v := range A {

		if v > 0 {
			if !hasPositives {
				hasPositives = true
				firstP = v
				next = firstP + 1
				last = v
			}

			if firstP != 1 {
				return 1
			}

			if v != last {

				if v != next {
					return next
				}

				last = v
				next = v + 1
			}
		}
	}

	if !hasPositives {
		return 1
	}

	return A[len(A)-1] + 1

}

/*
func main() {
	a := MissingInteger([]int{})                 // 1
	b := MissingInteger([]int{2, 3, 4})          // 1
	c := MissingInteger([]int{2})                // 1
	d := MissingInteger([]int{1, 3, 6, 4, 1, 2}) //5
	e := MissingInteger([]int{4, 5, 6, 2})       //3
	f := MissingInteger([]int{1, 2, 3})          //4
	g := MissingInteger([]int{-1, -3})           //1

	h := MissingInteger([]int{1, 3, 6, 4, 1, 2}) //5
	i := MissingInteger([]int{-1})               //1
	j := MissingInteger([]int{-1, 0, 1})         //2
	k := MissingInteger([]int{-3, -2 - 1, 0, 1}) //2
	l := MissingInteger([]int{1})                //2
	m := MissingInteger([]int{-3, -2 - 1, 0})    //1
	n := MissingInteger([]int{-3, -1, 0})        //1
	o := MissingInteger([]int{-3, -1, 0, 1})     //2
	p := MissingInteger([]int{4, 5, 6, 2})       //3

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
}*/
