package codility

import (
	"fmt"
	"strings"
)

func Pow1(A, B int) int {

	result := 0

	binary := fmt.Sprintf("%b", (A * B))

	s := strings.Split(binary, "")

	for _, v := range s {
		if v == "1" {
			result++
		}
	}

	return result

}
