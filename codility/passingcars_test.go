package codility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassingCars(t *testing.T) {

	tests := []struct {
		name  string
		input []int

		expected int
	}{
		{name: "case 1 ", input: []int{0, 1, 0}, expected: 1},
		{name: "case 2 ", input: []int{0, 1, 0, 1, 1}, expected: 5},
		//{name: "case 3 ", input: []int{0, 1, 0, 1, 1}, expected: -1},
	}

	for _, test := range tests {

		actual := PassingCars(test.input)
		assert.Equal(t, test.expected, actual)

	}
}
