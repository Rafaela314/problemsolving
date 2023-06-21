package codility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBrackets(t *testing.T) {

	tests := []struct {
		s        string
		expected int
	}{
		{s: "", expected: 1},
		{s: "{[()()]}", expected: 1},
		{s: "([)()]", expected: 0},
		{s: ")(", expected: 0},
		{s: "({{({}[]{})}}[]{})", expected: 1},
	}

	for _, test := range tests {

		result := Brackets(test.s)

		assert.Equal(t, test.expected, result)

	}

}
