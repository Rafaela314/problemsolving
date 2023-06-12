package others

import (
	"fmt"
	"strconv"
)

func File3(queries [][]string) []string {
	m := make(map[string]int)

	result := []string{}

	for _, v := range queries {
		if v[0] == "ADD" {
			result = append(result, "")
			if _, ok := m[v[1]]; ok {
				m[v[1]] = m[v[1]] + 1
			} else {
				m[v[1]] = 1
			}

		}

		if v[0] == "EXISTS" {
			if _, ok := m[v[1]]; ok {
				result = append(result, "true")

			} else {

				result = append(result, "false")
			}
			fmt.Printf("\n M %v  \n", m)
			fmt.Printf("\n M %v  \n", result)

		}

		if v[0] == "REMOVE" {
			if _, ok := m[v[1]]; ok {
				result = append(result, "true")

				m[v[1]] = m[v[1]] - 1

				if m[v[1]] == 0 {
					delete(m, v[1])
				}
			} else {
				result = append(result, "false")
			}

		}

		if v[0] == "GET_NEXT" {

			found := false

			target, _ := strconv.Atoi(v[1])
			for i := 0; i < len(m); i++ {
				target += 1
				t := strconv.Itoa(target)

				if _, ok := m[t]; ok {
					result = append(result, t)
					found = true
					break
				}

				fmt.Printf("\n Result %v target %v  \n", result, target)

			}

			if !found {
				result = append(result, "")
			}

		}

		//fmt.Printf("\n M %v  \n", m[v[1]])
		//fmt.Printf("\n A %v  \n", result)
		//fmt.Printf("\n M %v  \n", m)
	}

	return result
}
