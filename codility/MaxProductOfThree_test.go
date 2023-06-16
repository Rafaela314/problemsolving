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
		{name: "case 3 ", x: []int{-1, -2, 1}, expected: 2},
		{name: "case 4 ", x: []int{-10, -2, -20, 2}, expected: 400},
		{name: "case 5 ", x: []int{10, 10, 10}, expected: 1000},
		{name: "case 6 ", x: []int{-10, -2, -4}, expected: -80},
		{name: "case 7 ", x: []int{-5, -6, -4, -7, -10}, expected: -120},
		{name: "case 8 ", x: []int{-5, 0, 1, 2}, expected: 0},
	}

	for _, test := range tests {

		actual := MaxProductOfThree(test.x)
		assert.Equal(t, test.expected, actual)

	}

}
