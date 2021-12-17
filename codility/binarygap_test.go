package codility

import (
	"testing"
)

func TestSolutionIteration(t *testing.T) {

	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{name: "case 1 ", input: 9, expected: 2},
		{name: "case 2 ", input: 529, expected: 4},
		{name: "case 3 ", input: 20, expected: 1},
		{name: "case 4 ", input: 15, expected: 0},
		{name: "case 5 ", input: 32, expected: 0},
	}

	for _, test := range tests {

		result := SolutionIteration(test.input)

		if result != test.expected {
			t.Errorf("FAIL --- %s, expected '%v', received '%v'", test.name, test.expected, result)
		}
	}

	/*for _, test := range tests {
		actual := SolutionIteration(test.input)
		assert.Equal(t, test.expected, actual)

	}*/

}
