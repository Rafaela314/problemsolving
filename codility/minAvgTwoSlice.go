package codility

func MinAvgTwoSlice(A []int) int {

	//{name: "case 1 ", x: []int{4, 2, 2, 5, 1, 5, 8}, expected: 1},

	/*	for i := 0; i < len(A)-1; i++ {

			fmt.Printf("\nI %v, AI %v AIcom1 %v \n", i, A[i], A[i+1])
			if A[i+1] < A[i] {
				continue
			}
			for j := i + 1; j < len(A); j++ {

				fmt.Printf("\nJ: i + 1 %v, A[j] %v\n", j, A[j])

				if A[j] <= A[j+1] {
					continue
				} else {

					i = j + 1

					break
				}
			}

		}

		//fmt.Printf("\nBEGIN %v, END: %v\n", begin, end)*/

	return 0
}

/*func MinAvgTwoSlice(A []int) int {

	//{name: "case 1 ", x: []int{4, 2, 2, 5, 1, 5, 8}, expected: 1},

	if len(A) == 2 {
		return (A[0] + A[1]) / 2
	}

	var position int

	sum := 0
	min := sum / 2
	count := 2
	med := sum / 2

	for i := 0; i < len(A)-1; i++ {

		sum = A[i] + A[i+1]
		count = 2
		med = sum / count

		fmt.Printf("\n SUM %v MED: %v\n", sum, med)

		if i == 0 {
			min = med
		}

		if i != 0 && A[i+1] < A[i] {
			continue
		}

		for j := i + 2; j < len(A)-1; j++ {

			fmt.Printf("\n I %v J: %v\n", i, j)
			if A[j] < med {
				sum += A[j]
				count += 1
				med = (sum / count)

				if med < min {
					min = med
					position = i
				}
			}

			fmt.Printf("\n AFTER JJJJJ SUM:  %v, count: %v MIN %v Med %v \n", sum, count, min, med)

			if A[j] > med {
				fmt.Printf("\n ACIMA, CORTA J %v\n", j)
				break
			}

		}

		fmt.Printf("\n SUM:  %v, count: %v MIN %v Med %v \n", sum, count, min, med)

	}

	return position
}

/*

func MinAvgTwoSlice(A []int) int {

	if len(A) == 2 {
		return (A[0] + A[1]) / 2
	}

	sum := A[0] + A[1]
	min := sum / 2
	position := 0

	for i := 0; i < len(A)-1; i++ {
		//{name: "case 1 ", x: []int{4, 2, 2, 5, 1, 5, 8}, expected: 1},
		fmt.Printf("\n I: %v MIN : %v Positon:  %v\n", i, min, position)
		if i != 0 {
			if A[i-1] < A[i] {
				continue
			}
		}

		for j := 0; j < len(A)-1; j++ {
			if i == 0 {
				j = 2
			}

			fmt.Printf("\n SUM3: %v J + 1 : %v \n", sum+A[j], (j + 1))

			if (sum+A[j])/(j+1) > min {
				fmt.Printf("\n Aumenta a média, pode para:  %v média: %v\n", j, (sum+A[j+1])/(j+1))
				j = len(A) - 2
				//break
			} else {
				sum = sum + A[j]
				min = sum / (j + 1)
				position = i
			}
		}

	}

	return position
}*/
