package hacker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueensAttackII(t *testing.T) {

	tests := []struct {
		c_q       int32
		r_q       int32
		n         int32
		k         int32
		obstacles [][]int32
		expected  int32
	}{
		{c_q: 1, r_q: 1, n: 1, k: 0, obstacles: [][]int32{}, expected: 0},
		{c_q: 4, r_q: 4, n: 4, k: 0, obstacles: [][]int32{}, expected: 9},
		{c_q: 3, r_q: 4, n: 5, k: 3, obstacles: [][]int32{{5, 5}, {4, 2}, {2, 3}}, expected: 10},
	}

	for _, test := range tests {

		actual := QueensAttackII(test.n, test.k, test.r_q, test.c_q, test.obstacles)
		assert.Equal(t, test.expected, actual)

	}
}
