package codility

import "fmt"

func Peak(A []int) int {
	// Implement your solution here

	if len(A) <= 2 {
		return 0
	}

	if len(A) == 3 {
		if A[0] < A[1] && A[1] > A[2] {
			return 1
		}
		return 0
	}

	peaks := []int{}
	r := 0
	sum := 0
	for i := 1; i < len(A); i++ {

		if A[i] > A[i-1] && A[i] > A[i+1] {
			if r == 0 {
				r = i

			} else {
				peaks = append(peaks, i-r)
				sum = sum + i - r
				r = i

			}
		}
	}

	if len(peaks) == 0 {
		return 0
	}

	if len(peaks) == 1 {
		return 1
	}

	flag := 0
	x := 1
	teto := sum / len(peaks)
	max := 0

	track := 0

	for i := 0; i < len(peaks); i++ {

		track = track + peaks[i]
		if peaks[i] < x {
			peaks[i+1] = peaks[i+1] + peaks[i]
			peaks = peaks[i:]
			i = 0
			fmt.Printf("\n MENOR Peak %v  flag: %v track %v max %v x %v\n", peaks, flag, track, max, x)
		} else {
			fmt.Printf("\n MAIOR Peak %v  flag: %v track %v max %v x %v\n", peaks, flag, track, max, x)
			flag++
			if flag == x {
				fmt.Printf("\n CHEGOU X Peak %v  flag: %v track %v max %v x %v\n", peaks, flag, track, max, x)
				if x == teto {
					return x
				}
				if flag > max {
					max = flag
				}
				i = 0
				flag = 0
				x++
				fmt.Printf("\n VAI NO PROXPeak %v  flag: %v track %v max %v x %v\n", peaks, flag, track, max, x)
			}

		}
		fmt.Printf("\n Peak %v  flag: %v track %v max %v x %v\n", peaks, flag, track, max, x)
	}

	flag = len(peaks) + 1

	return flag

}
