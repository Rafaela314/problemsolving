package codility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOddOccurrencesInArray(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{name: "case 1 ", input: []int{9, 3, 9, 3, 9, 7, 9}, expected: 7},
		{name: "case 2 ", input: []int{3, 4, 3, 4, 6, 1, 4, 4, 6}, expected: 1},
		{name: "case 3 ", input: []int{15, 3, 4, 3, 4, 6, 1, 4, 1, 15, 4, 6, 22}, expected: 22},
	}

	for _, test := range tests {

		actual := OddOccurrencesInArray2(test.input)
		assert.Equal(t, test.expected, actual)

	}
}
