package hacker

func CountApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) []int {
	// Write your code here

	var list []int

	var countA int
	var countB int

	for _, v := range apples {
		p := v + a

		if p >= s && p <= t {

			countA = countA + 1
		}
	}

	for _, v2 := range oranges {
		p := v2 + b

		if p >= s && p <= t {

			countB = countB + 1
		}

	}

	list = append(list, countA)
	list = append(list, countB)

	return list

}
