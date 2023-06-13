package codility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTapeEquilibrium(t *testing.T) {

	tests := []struct {
		name  string
		input []int

		expected int
	}{
		{name: "case 1 ", input: []int{3, 1, 2, 4, 3}, expected: 1},
		{name: "case 2 ", input: []int{2, 4}, expected: 2},
		{name: "case 3 ", input: []int{1, 2, 7, 12}, expected: 2},
	}

	for _, test := range tests {

		actual := TapeEquilibrium(test.input)
		assert.Equal(t, test.expected, actual)

	}
}
