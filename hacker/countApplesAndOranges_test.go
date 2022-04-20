package hacker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountApplesAndOranges(t *testing.T) {

	tests := []struct {
		name    string
		init    int32
		end     int32
		a       int32
		b       int32
		apples  []int32
		oranges []int32

		expected []int
	}{
		{name: "case 1 ", init: 7, end: 10, a: 4, b: 12, apples: []int32{2, 3, -4}, oranges: []int32{3, -2, -4}, expected: []int{1, 2}},
	}

	for _, test := range tests {

		actual := CountApplesAndOranges(test.init, test.end, test.a, test.b, test.apples, test.oranges)
		assert.Equal(t, test.expected, actual)

	}
}
