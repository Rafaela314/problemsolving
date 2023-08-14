package codility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPeak(t *testing.T) {

	tests := []struct {
		name  string
		input []int

		expected int
	}{
		{name: "case 1 ", input: []int{0, 1, 0}, expected: 1},
		{name: "case 2 ", input: []int{1, 8, 1, 5, 1, 5, 1, 1, 1, 1, 10, 1}, expected: 3},
		{name: "case 3 ", input: []int{1, 2}, expected: 0},
	}

	for _, test := range tests {

		actual := Peak(test.input)
		assert.Equal(t, test.expected, actual)

	}
}
