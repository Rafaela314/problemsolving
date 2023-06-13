package codility

import (
	"testing"
)

func TestIsUnique(t *testing.T) {
	cases := []struct {
		input    []int
		num      int
		expected int
	}{
		{[]int{1, 3, 1, 4, 2, 3, 5, 4}, 5, 6},
		{[]int{1, 1, 1}, 1, -1},
		{[]int{1}, 1, -1},
		{[]int{1, 3, 1, 4, 2, 3, 5, 4, 8, 3, 7, 6, 3, 2}, 8, 11},
	}
	for _, c := range cases {
		actual := FrogRiverOne(c.num, c.input)
		if actual != c.expected {
			t.Fatalf("Input %v. Expected: %b, actual: %b\n", c.input, c.expected, actual)
		}
	}
}
