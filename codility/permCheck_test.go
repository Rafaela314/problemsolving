package codility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermCheck(t *testing.T) {

	tests := []struct {
		name     string
		a        []int
		expected int
	}{
		{name: "case 1 ", a: []int{4, 3, 2, 1}, expected: 1},
		{name: "case 2 ", a: []int{4, 3, 1}, expected: 0},
	}

	for _, test := range tests {

		actual := PermCheck(test.a)
		assert.Equal(t, test.expected, actual)

	}

}
