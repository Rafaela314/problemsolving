package codility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTime(t *testing.T) {

	tests := []struct {
		input    int
		expected string
	}{
		{input: 0, expected: "now"},
		{input: 14, expected: "14 seconds"},
		{input: 722, expected: "12 minutes and 2 seconds"},
		{input: 19202, expected: "5 hour 20 minutes and 2 seconds"},
		{input: 267604, expected: "3 days 2 hour 20 minutes and 4 seconds"},
	}

	for _, test := range tests {

		actual := GetTime(test.input)
		assert.Equal(t, test.expected, actual)

	}

}
