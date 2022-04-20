package hacker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBomberMan(t *testing.T) {

	tests := []struct {
		name  string
		input []string
		n     int32

		expected []string
	}{
		{name: "case 1 ", input: []string{".......", "...O...", "....O..", ".......", "OO.....", "OO....."}, n: 3, expected: []string{"OOO.OOO", "OO...OO", "OOO...O", "..OO.OO", "...OOOO", "...OOOO"}},
	}

	for _, test := range tests {

		actual := BomberMan(test.n, test.input)
		assert.Equal(t, test.expected, actual)

	}
}
