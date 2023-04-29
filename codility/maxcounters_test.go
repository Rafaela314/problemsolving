package codility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxCounters(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		k        int
		expected []int
	}{
		{name: "case 1 ", input: []int{3, 4, 4, 6, 1, 4, 4}, k: 5, expected: []int{3, 2, 2, 4, 2}},
		{name: "case 2 ", input: []int{3, 4, 4, 6, 1, 4, 4}, k: 5, expected: []int{3, 2, 2, 4, 2}},
		{name: "case 3 ", input: []int{7, 7, 7}, k: 2, expected: []int{0, 0}},
		{name: "case 4 ", input: []int{1}, k: 1, expected: []int{1}},
	}

	for _, test := range tests {

		actual := MaxCounters(test.k, test.input)
		assert.Equal(t, test.expected, actual)

	}
}
