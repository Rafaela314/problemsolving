package codility

import (
	"testing"
)

func TestCountDiv(t *testing.T) {

	tests := []struct {
		name     string
		a        int
		b        int
		k        int
		expected int
	}{
		{name: "case 1 ", a: 6, b: 11, k: 2, expected: 3},
		{name: "case 2 ", a: 11, b: 345, k: 17, expected: 20},
		{name: "case 3 ", a: 0, b: 0, k: 11, expected: 1},
		{name: "case 4 ", a: 1, b: 1, k: 11, expected: 0},
	}

	for _, test := range tests {

		result := CountDiv(test.a, test.b, test.k)

		if result != test.expected {
			t.Errorf("FAIL --- %s, expected '%v', received '%v'", test.name, test.expected, result)
		}
	}

}
