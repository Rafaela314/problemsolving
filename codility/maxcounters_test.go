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
		{name: "case 1 ", input: []int{3, 4, 4, 6, 1, 4, 4, 6}, k: 5, expected: []int{4, 4, 4, 4, 4}},
		{name: "case 1 ", input: []int{1}, k: 1, expected: []int{1}},
	}

	for _, test := range tests {

		actual := MaxCounters(test.input, test.k)
		assert.Equal(t, test.expected, actual)

	}
}
