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
		//{name: "case 1 ", input: []int{3, 4, 4, 6, 1, 4, 4, 6}, k: 5, expected: []int{4, 4, 4, 4, 4}},
		//{name: "case 1 ", input: []int{1}, k: 1, expected: []int{1}},
	}

	for _, test := range tests {

		actual := OddOccurrencesInArray(test.input)
		assert.Equal(t, test.expected, actual)

	}
}
