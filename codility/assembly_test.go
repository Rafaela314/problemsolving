package codility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssembly(t *testing.T) {

	tests := []struct {
		name  string
		input []int
		x     int
		y     int

		expected int
	}{
		{name: "case 1 ", input: []int{1, 1, 3}, x: 1, y: 1, expected: 2},
		{name: "case 2", input: []int{6, 5, 5, 4, 3}, x: 8, y: 9, expected: 4},
		{name: "case 3", input: []int{6, 5, 2, 1, 8}, x: 17, y: 5, expected: 5},
		{name: "case 4 ", input: []int{1}, x: 1, y: 1, expected: 1},
	}

	for _, test := range tests {

		actual := Assembly(test.input, test.x, test.y)
		assert.Equal(t, test.expected, actual)

	}

}
