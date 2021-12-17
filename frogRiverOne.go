package main

/*
A small frog wants to get to the other side of a river. The frog is initially located on one bank of the river (position 0)
and wants to get to the opposite bank (position X+1). Leaves fall from a tree onto the surface of the river.

You are given an array A consisting of N integers representing the falling leaves. A[K] represents the position where one leaf falls at time K,
measured in seconds.

The goal is to find the earliest time when the frog can jump to the other side of the river.
The frog can cross only when leaves appear at every position across the river from 1 to X (that is,
we want to find the earliest moment when all the positions from 1 to X are covered by leaves).
You may assume that the speed of the current in the river is negligibly small,
i.e. the leaves do not change their positions once they fall in the river.

For example, you are given integer X = 5 and array A such that:
  A[0] = 1
  A[1] = 3
  A[2] = 1
  A[3] = 4
  A[4] = 2
  A[5] = 3
  A[6] = 5
  A[7] = 4

In second 6, a leaf falls into position 5. This is the earliest time when leaves appear in every position across the river.

Write a function:

    class Solution { public int solution(int X, int[] A); }

that, given a non-empty array A consisting of N integers and integer X, returns the earliest time when the frog can jump to the other side of the river.

If the frog is never able to jump to the other side of the river, the function should return −1.

For example, given X = 5 and array A such that:
  A[0] = 1
  A[1] = 3
  A[2] = 1
  A[3] = 4
  A[4] = 2
  A[5] = 3
  A[6] = 5
  A[7] = 4

the function should return 6, as explained above.*/
import "fmt"

func FrogRiverOne(X int, A []int) int {

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

				fmt.Printf("\n X : %v \n", X)
				fmt.Printf("\n COUNT : %v \n", counter)
				fmt.Printf("\n K : %v \n", k)

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

/*
func main() {
	x := FrogRiverOne(5, []int{1, 3, 1, 4, 2, 3, 5, 4})
	y := FrogRiverOne(8, []int{1, 3, 1, 4, 2, 3, 5, 4})
	z := FrogRiverOne(1, []int{1})
	h := FrogRiverOne(2, []int{2, 2, 2, 2, 2})

	fmt.Println(x, y, z, h)
}*/
