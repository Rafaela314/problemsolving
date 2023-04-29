package codility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistinct(t *testing.T) {

	tests := []struct {
		name  string
		input []int

		expected int
	}{
		{name: "case 1 ", input: []int{2, 1, 1, 2, 3, 1}, expected: 3},
	}

	for _, test := range tests {

		actual := Distinct(test.input)
		assert.Equal(t, test.expected, actual)

	}

}
