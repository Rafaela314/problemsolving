package codility

import (
	"fmt"
	"sort"
)

/*
A non-empty array A consisting of N integers is given. The array contains an odd number of elements, and each element of the array can be paired with another element that has the same value, except for one element that is left unpaired.

For example, in array A such that:
  A[0] = 9  A[1] = 3  A[2] = 9
  A[3] = 3  A[4] = 9  A[5] = 7
  A[6] = 9

        the elements at indexes 0 and 2 have value 9,
        the elements at indexes 1 and 3 have value 3,
        the elements at indexes 4 and 6 have value 9,
        the element at index 5 has value 7 and is unpaired.

Write a function:

    class Solution { public int solution(int[] A); }

that, given an array A consisting of N integers fulfilling the above conditions, returns the value of the unpaired element.

For example, given array A such that:
  A[0] = 9  A[1] = 3  A[2] = 9
  A[3] = 3  A[4] = 9  A[5] = 7
  A[6] = 9

the function should return 7, as explained in the example above.
*/

func OddOccurrencesInArray(A []int) int {

	var result int

	m := make(map[int]int) // [number]count

	sort.Ints(A)
	fmt.Printf("\n %v a \n", A)

	for _, v := range A {
		//exists
		value, ok := m[v]
		if ok {
			m[v] = value + 1
		} else {
			m[v] = 1
		}
	}

	for key, vv := range m {
		if vv%2 != 0 {
			result = key
		}
	}

	fmt.Printf("\n M %v \n", m)

	return result

}

/*
func main() {
	x := OddOccurrencesInArray([]int{3, 7, 4242, 5, 5, 7, 7, 7, 4242, 5, 5})

	fmt.Println(x)
}
*/
