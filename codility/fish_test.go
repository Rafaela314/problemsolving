package codility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFish(t *testing.T) {

	tests := []struct {
		name     string
		A        []int
		B        []int
		expected int
	}{
		{name: "case 1 ", A: []int{4, 3, 2, 1, 5}, B: []int{0, 1, 0, 0, 0}, expected: 2},
		{name: "case 2 ", A: []int{3, 5, 2, 9, 2, 3}, B: []int{1, 0, 0, 1, 1, 0}, expected: 3},
	}

	for _, test := range tests {

		actual := Fish(test.A, test.B)
		assert.Equal(t, test.expected, actual)

	}

}
