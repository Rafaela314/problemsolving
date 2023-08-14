package codility

/*Example

For a = [24, 85, 0], the output should be
solution(a) = 21784.

An array [24, 85, 0] looks like [00011000, 01010101, 00000000] in binary.
After packing these into one number we get 00000000 01010101 00011000 (spaces are placed for convenience), which equals to 21784.*/

import (
	"fmt"
	"strconv"
)

func solution(a []int) int {
	s := ""

	for _, v := range a {
		fmt.Println(v)
		fmt.Println(fmt.Sprintf("%b", v))
		s = fmt.Sprintf("%08b", v) + s
	}

	x, _ := strconv.ParseInt(s, 2, 0)
	return int(x)
}
