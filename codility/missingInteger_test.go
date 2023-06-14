package codility

import (
	"testing"
)

func TestMissingInteger(t *testing.T) {
	cases := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 3, 6, 4, 1, 2}, 5},
		{[]int{1, 2, 3}, 4},
		{[]int{-1, -3}, 1},
	}
	for _, c := range cases {
		actual := MissingInteger(c.input)
		if actual != c.expected {
			t.Fatalf("Input %v. Expected: %b, actual: %b\n", c.input, c.expected, actual)
		}
	}
}
