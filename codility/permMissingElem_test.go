package codility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermMissingElem(t *testing.T) {

	tests := []struct {
		name     string
		a        []int
		expected int
	}{
		{name: "case 1 ", a: []int{2, 3, 1, 5}, expected: 4},
		{name: "case 2 ", a: []int{}, expected: 1},
		{name: "case 3 ", a: []int{1, 2, 3}, expected: 4},
		{name: "case 4 ", a: []int{2, 3}, expected: 1},
		{name: "case 5 ", a: []int{2, 3, 1, 5, 4}, expected: 6},
	}

	for _, test := range tests {

		actual := PermMissingElem(test.a)
		assert.Equal(t, test.expected, actual)

	}

}
