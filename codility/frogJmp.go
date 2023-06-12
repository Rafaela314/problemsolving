package codility

import "fmt"

func FrogJmp(X, Y, D int) int {

	var result int

	distance := Y - X
	if distance%D != 0 {
		result = ((distance / D) + 1)
	} else {
		result = distance / D
	}

	fmt.Printf("\n %v result \n", result)
	return result
}
