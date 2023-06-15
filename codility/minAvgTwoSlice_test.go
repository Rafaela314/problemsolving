package codility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinAvgTwoSlice(t *testing.T) {

	tests := []struct {
		name     string
		x        []int
		expected int
	}{
		{name: "case 1 ", x: []int{4, 2, 2, 5, 1, 5, 8}, expected: 60},
	}

	for _, test := range tests {

		actual := MinAvgTwoSlice(test.x)
		assert.Equal(t, test.expected, actual)

	}

}
