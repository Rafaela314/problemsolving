package codility

import (
	"testing"
)

func TestFrogJmp(t *testing.T) {

	cases := []struct {
		name     string
		x        int
		y        int
		d        int
		expected int
	}{
		{name: "case 1 ", x: 10, y: 85, d: 30, expected: 3},
	}

	for _, c := range cases {

		actual := FrogJmp(c.x, c.y, c.d)
		if actual != c.expected {
			t.Fatalf("Input x: %d y:%d d: %d. Expected: %b, actual: %b\n", c.x, c.y, c.d, c.expected, actual)
		}

	}

}
