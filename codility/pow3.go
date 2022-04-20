package codility

import (
	"fmt"
	"strings"
)

/* A clening robot is cleaning a grid size NxM. represented by an array of N strings.
The robot start cleaning on the top left corner. It goes straight forward until it find an X, then it turns 90 degree clockwise*/

func Pow3(A []string) int {

	//m := make(map[string][]string)

	//var stop bool

	var B [][]string

	// Prepare split

	for _, value := range A {
		s := strings.Split(value, "")
		B = append(B, s)

	}

	// keep the count of clean spots until it hit an already clean spot with same flag direction

	_, col := len(A), len(A[0])

	//starting point
	x, y := 0, 0

	//move forward
	for i := y; i < col; i++ {

		if B[x][i] == "." {

			fmt.Println("COUNT FORWARD")
			//s := fmt.Sprintf("%d,%d", i, y)
			//m[s] = append(m[s], "south")
		} else {
			if i > 0 {
				y = i - 1
			}
			fmt.Println("COUNT STOP")
			break
		}
	}

	/*
		//move down
		for i := y; i < row; i++ {
			fmt.Println(B[i][y])
			if B[i][y] == "." {
				fmt.Println("COUNT SOUTH")

				//s := fmt.Sprintf("%d,%d", i, y)
				//m[s] = append(m[s], "south")
			}
		}

		//move backward
		for i := y; i < col; i++ {
			fmt.Println(B[i][y])
			if B[x][i] == "." {
				fmt.Println("COUNT BACKWARD")

				//s := fmt.Sprintf("%d,%d", i, y)
				//m[s] = append(m[s], "south")
			}
		}
	*/
	return 1

}

/*	var matrix [][]string

	for i, _ := range A {
		s := strings.Split(A[i], "")
		matrix = append(matrix, s)

	}

	totalRows := len(A)
	totalColumns := len(A[0])

	startRow := 0
	endRow := len(matrix)

	startColumn := 0
	endColumn := len(matrix)

	for startRow <= endRow && startColumn <= endColumn {

		for i := startColumn; i <= endColumn; i++ {
			output = append(output, matrix[startRow][i])
		}

	}

	//m := make(map[int]int)

	/*
		count := 0
		position := 0

		for position; i < len(A); i++ {


			for i, y := range s {
				if y == "." {
					count++
					fmt.Printf("\n COUNT : %v \n", count)
				} else {
					fmt.Printf("\n BRESK : %v \n", i)
					position = i-1
					break
				}
			}
		}

		return count*/
