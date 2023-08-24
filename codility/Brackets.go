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

/*

var m = map[string]string{
    "{": "}",
    "(": ")",
    "[": "]",
}

func ValidBraces(str string) bool {
    s := make([]string, 0)
    for _, r := range str {
        r := string(r)
        if len(s) > 0 && m[s[len(s) - 1]] == r {
            s = s[:len(s) - 1]
        } else {
            s = append(s, r)
        }
    }
    return len(s) == 0
}

// other way

func ValidBraces(str string) bool {
  stack := make([]rune, 0)
	for _, c := range str {
		switch c {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		case ')', ']', '}':
      if len(stack) == 0 {
        return false
      }
			if stack[len(stack)-1] != c {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
  return len(stack) == 0
}
*/
