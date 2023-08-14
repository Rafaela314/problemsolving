package codility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalidrome(t *testing.T) {

	tests := []struct {
		name string
		A    string

		expected bool
	}{
		{name: "case 1 ", A: "aabaa", expected: true},
		{name: "case 2 ", A: "abac", expected: false},
		{name: "case 2 ", A: "aaabaaaa", expected: false},
		{name: "case 2 ", A: "a", expected: true},
		{name: "case 2 ", A: "z", expected: true},
		{name: "case 2 ", A: "az", expected: false},
	}

	for _, test := range tests {

		actual := IsPalidrome(test.A)
		assert.Equal(t, test.expected, actual)

	}

}
