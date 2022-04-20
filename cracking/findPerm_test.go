package cracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPerm(t *testing.T) {

	tests := []struct {
		name     string
		a        string
		b        string
		expected int
	}{
		{name: "case 1 ", a: "cbabadcbbabbcbabaabccbabc", b: "abbc", expected: 7},
	}

	for _, test := range tests {

		actual := FindPerm(test.a, test.b)
		assert.Equal(t, test.expected, actual)

	}

}
