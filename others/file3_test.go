package others

import (
	"testing"
)

func TestFile3(t *testing.T) {

	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{name: "case 1 ", input: 9, expected: 2},
	}

	for _, test := range tests {

		result := File3(test.input)

		if result != test.expected {
			t.Errorf("FAIL --- %s, expected '%v', received '%v'", test.name, test.expected, result)
		}
	}

	/*for _, test := range tests {
		actual := SolutionIteration(test.input)
		assert.Equal(t, test.expected, actual)

	}*/

}
