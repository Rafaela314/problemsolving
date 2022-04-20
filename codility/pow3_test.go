package codility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPow3(t *testing.T) {

	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{name: "case 1 ", input: []string{"...X..", "....XX", "..X..."}, expected: 6},
		{name: "case 1 ", input: []string{"."}, expected: 1},
	}

	for _, test := range tests {

		actual := Pow3(test.input)
		assert.Equal(t, test.expected, actual)

	}

}
