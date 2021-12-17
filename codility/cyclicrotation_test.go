package codility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCyclicRotation(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		k        int
		expected []int
	}{
		{name: "case 1 ", input: []int{3, 8, 9, 7, 6}, k: 3, expected: []int{9, 7, 6, 3, 8}},
		{name: "case 2 ", input: []int{0, 0, 0}, k: 1, expected: []int{0, 0, 0}},
		{name: "case 2 ", input: []int{1, 2, 3, 4}, k: 4, expected: []int{1, 2, 3, 4}},
	}

	for _, test := range tests {

		actual := CyclicRotation(test.input, test.k)
		assert.Equal(t, test.expected, actual)

		//if result != test.expected {
		//	t.Errorf("FAIL --- %s, expected '%v', received '%v'", test.name, test.expected, result)
		//}
	}

	/*for _, test := range tests {
		actual := SolutionIteration(test.input)
		assert.Equal(t, test.expected, actual)

	}*/

}
