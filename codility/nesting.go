package codility

func Solution(S string) int {

	if len(S)%2 != 0 {
		return 0
	}

	var count int

	for _, v := range S {

		if v == rune('(') {
			count++
		}
		if v == rune(')') {
			if count == 0 {
				return 0
			}
			count--
		}
	}
	if count == 0 {
		return 1
	}
	return 0
}
