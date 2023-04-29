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
		{s: "CAGCCTA", expected: 1},
		//{name: "case 1 ", s: "CAGCCTATGC", p: []int{2, 5, 0}, q: []int{4, 5, 6}, expected: []int{2, 4, 1}},
	}

	for _, test := range tests {

		result := Brackets(test.s, test.p, test.q)

		assert.Equal(t, test.expected, result)

	}

}
