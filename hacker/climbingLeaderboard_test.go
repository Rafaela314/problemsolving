package hacker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbingLeaderBoard(t *testing.T) {

	tests := []struct {
		name        string
		input       []int32
		inputPlayer []int32

		expected []int32
	}{
		{name: "case 1 ", input: []int32{100, 100, 50, 40, 40, 20, 10}, inputPlayer: []int32{5, 25, 50, 120}, expected: []int32{6, 4, 2, 1}},
		{name: "case 2 ", input: []int32{100, 90, 90, 80, 75, 60}, inputPlayer: []int32{50, 65, 77, 90, 102}, expected: []int32{6, 5, 4, 2, 1}},
	}

	for _, test := range tests {

		actual := ClimbingLeaderBoard(test.input, test.inputPlayer)
		assert.Equal(t, test.expected, actual)

	}
}
