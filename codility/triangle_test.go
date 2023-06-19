package codility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTriangle(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{name: "case 1 ", input: []int{10, 2, 5, 1, 8, 20}, expected: 1},
		{name: "case 2 ", input: []int{10, 50, 5, 1}, expected: 0},
		{name: "case 3 ", input: []int{}, expected: 0},
		{name: "case 4 ", input: []int{10, 5, 8}, expected: 1},
		{name: "case 5 ", input: []int{5, 3, 3}, expected: 1},
	}

	for _, test := range tests {

		actual := Triangle(test.input)
		assert.Equal(t, test.expected, actual)

	}

}
