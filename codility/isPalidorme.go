package codility

import "fmt"

func IsPalidrome(inputString string) bool {

	reference := len(inputString) - 1

	for k, v := range inputString {
		fmt.Printf("\n TEMP : %v ref %v\n", v, inputString[reference])

		if v != rune(inputString[reference]) {
			return false
		}

		if len(inputString)%2 == 0 && k == len(inputString)/2 {
			break
		}

		if len(inputString)%2 != 0 && k == (len(inputString)-1)/2 {
			break
		}

		reference--

	}

	return true

}
