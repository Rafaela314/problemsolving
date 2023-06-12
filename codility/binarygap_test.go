package codility

import (
	"testing"
)

func TestSolutionIteration(t *testing.T) {

	cases := []struct {
		name     string
		input    int
		expected int
	}{
		{name: "case 1 ", input: 9, expected: 2},
		{name: "case 2 ", input: 529, expected: 4},
		{name: "case 3 ", input: 20, expected: 1},
		{name: "case 4 ", input: 15, expected: 0},
		{name: "case 5 ", input: 32, expected: 0},
	}

	for _, c := range cases {
		actual := SolutionIteration(c.input)
		if actual != c.expected {
			t.Fatalf("Input %d. Expected: %b, actual: %b\n", c.input, c.expected, actual)
		}

	}
}

var num = 100

func BenchmarkSolutionIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SolutionIteration(num)
	}
}

func BenchmarkSolutionIteration2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SolutionIteration2(num)
	}
}
