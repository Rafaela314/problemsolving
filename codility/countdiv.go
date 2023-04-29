package codility

//"fmt"
//"sync"

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

/*
type Counter struct {
	counter int
	mu      sync.Mutex
}

func CountDiv(A, B, K int) int {
	c := Counter{counter: 0}
	var wg sync.WaitGroup

	for i := A; i <= B; i++ {
		//fmt.Printf("\n A %v \n", A)
		wg.Add(1)

		go func(b int, s *sync.WaitGroup) {
			fmt.Printf("\n NUM %v \n", b)
			if b%K == 0 {
				//fmt.Printf("\n A %v \n", b)
				c.mu.Lock()
				c.counter++
				c.mu.Unlock()

			}
			defer s.Done()
		}(i, &wg)

	}
	wg.Wait()
	return c.counter
}
*/
/*Write a function:

func Solution(A int, B int, K int) int

that, given three integers A, B and K, returns the number of integers within the range [A..B] that are divisible by K, i.e.:

{ i : A ≤ i ≤ B, i mod K = 0 }

For example, for A = 6, B = 11 and K = 2, your function should return 3, because there are three numbers divisible by 2 within the range [6..11], namely 6, 8 and 10.

Write an efficient algorithm for the following assumptions:

A and B are integers within the range [0..2,000,000,000];
K is an integer within the range [1..2,000,000,000];
A ≤ B.*/
