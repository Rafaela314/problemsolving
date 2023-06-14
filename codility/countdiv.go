package codility

func CountDiv(A, B, K int) int {

	if A == 0 && B == 0 {
		return 1
	}

	var firstDiv int
	var found bool

	for i := A; i <= B; i++ {
		if i%K == 0 {
			firstDiv = i
			found = true
			break
		}
	}

	if !found && firstDiv == 0 {
		return firstDiv
	}

	count := ((B - firstDiv) / K) + 1

	return count

}
