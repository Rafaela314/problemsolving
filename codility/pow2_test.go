package codility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPow2(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{name: "case 1 ", input: []int{7, 3, 7, 3, 1, 3, 4, 1}, expected: 5},
		{name: "case 2 ", input: []int{2, 1, 1, 3, 2, 1, 1, 3}, expected: 3},
		{name: "case 3 ", input: []int{7, 5, 2, 7, 2, 7, 4, 7}, expected: 6},
	}

	for _, test := range tests {

		actual := Pow2(test.input)
		assert.Equal(t, test.expected, actual)

	}

}
