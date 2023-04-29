package codility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProductOfThree(t *testing.T) {

	tests := []struct {
		name string
		x    []int

		expected int
	}{
		{name: "case 1 ", x: []int{-3, 1, 2, -2, 5, 6}, expected: 60},
		{name: "case 2 ", x: []int{1, 2, 7, 3, 5, 6}, expected: 210},
		{name: "case 2 ", x: []int{-1, -2, 1}, expected: 2},
	}

	for _, test := range tests {

		actual := MaxProductOfThree(test.x)
		assert.Equal(t, test.expected, actual)

	}

}
