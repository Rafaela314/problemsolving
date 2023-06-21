package codility

import "fmt"

//O(N)

func Brackets(S string) int {

	var temp string

	if len(S) == 0 {
		return 1
	}

	if len(S)%2 != 0 {
		return 0
	}

	for _, v := range S {

		if v == rune('[') || v == rune('{') || v == rune('(') {
			temp += string(v)
		}

		if v == rune('}') {
			if len(temp) == 0 {
				return 0
			} else if string(temp[len(temp)-1]) != "{" {
				return 0
			} else {
				temp = temp[:len(temp)-1]
			}

		}

		if v == rune(')') {
			if len(temp) == 0 {
				return 0
			} else if string(temp[len(temp)-1]) != "(" {
				return 0
			} else {
				temp = temp[:len(temp)-1]
			}

		}

		if v == rune(']') {
			if len(temp) == 0 {
				return 0
			} else if string(temp[len(temp)-1]) != "[" {
				return 0
			} else {
				temp = temp[:len(temp)-1]
			}

		}

		fmt.Printf("\n TEMP : %v \n", temp)

	}

	if temp != "" {
		return 0
	}

	return 1
}
