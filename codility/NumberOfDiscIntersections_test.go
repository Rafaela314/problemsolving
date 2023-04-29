package codility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfDiscIntersections(t *testing.T) {

	tests := []struct {
		name string
		x    []int

		expected int
	}{
		{name: "case 1 ", x: []int{1, 5, 2, 1, 4, 0}, expected: 11},
	}

	for _, test := range tests {

		actual := NumberOfDiscIntersections(test.x)
		assert.Equal(t, test.expected, actual)

	}

}
