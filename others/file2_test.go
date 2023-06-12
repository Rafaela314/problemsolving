package others

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var b byte

func TestFile2(t *testing.T) {

	b := byte('e')

	tests := []struct {
		name     string
		input    string
		t        byte
		expected []int
	}{
		{name: "case 1 ", input: "mkteneetkybe", t: b, expected: []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0}}}

	for _, test := range tests {

		actual := File2(test.input, test.t)
		assert.Equal(t, test.expected, actual)

	}

}
