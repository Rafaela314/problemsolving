package codility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrogJmp(t *testing.T) {

	tests := []struct {
		name     string
		x        int
		y        int
		d        int
		expected int
	}{
		{name: "case 1 ", x: 10, y: 85, d: 30, expected: 3},
	}

	for _, test := range tests {

		actual := FrogJmp(test.x, test.y, test.d)
		assert.Equal(t, test.expected, actual)

	}

}
