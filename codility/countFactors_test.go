package codility

import (
	"testing"
)

func TestCountFactors(t *testing.T) {

	tests := []struct {
		name     string
		a        int
		expected int
	}{
		//{name: "case 1 ", a: 24, expected: 8},
		{name: "case 1 ", a: 24, expected: 0},
	}

	for _, test := range tests {

		result := CountFactors(test.a)

		if result != test.expected {
			t.Errorf("FAIL --- %s, expected '%v', received '%v'", test.name, test.expected, result)
		}
	}

}
