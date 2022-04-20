package hacker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAvenue1(t *testing.T) {

	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{name: "case 1 ", a: 3, b: 7, expected: 3},
		{name: "case 1 ", a: 0, b: 0, expected: 0},
	}

	for _, test := range tests {

		actual := Avenue1(test.a, test.b)
		assert.Equal(t, test.expected, actual)

	}

}
