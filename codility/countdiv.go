package codility

func CountDiv(A, B, K int) int {

	c := 0

	if A == 0 {
		c++
	}

	if K >= A {
		return (B / K) + c
	}

	if A%K == 0 {
		c++
	}

	return ((B - A) / K) + c

}
