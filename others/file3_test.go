package others

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFile3(t *testing.T) {

	tests := []struct {
		name     string
		input    [][]string
		expected []string
	}{

		{name: "case 1 ", input: [][]string{{"ADD", "1"}, {"ADD", "2"}, []string{"ADD", "2"}, []string{"ADD", "3"},
			[]string{"EXISTS", "1"}, []string{"EXISTS", "2"}, []string{"EXISTS", "3"},
			[]string{"REMOVE", "2"}, []string{"REMOVE", "1"},
			[]string{"EXISTS", "2"}, []string{"EXISTS", "1"}}, expected: []string{"", "", "", "", "true", "true", "true", "true", "true", "true", "false"}},
		{name: "case 2 ", input: [][]string{{"ADD", "1"}, {"ADD", "2"}, {"ADD", "2"}, []string{"ADD", "4"},
			[]string{"GET_NEXT", "1"}, []string{"GET_NEXT", "2"}, {"GET_NEXT", "3"}, []string{"GET_NEXT", "4"},
			[]string{"REMOVE", "2"}, []string{"GET_NEXT", "1"}, {"GET_NEXT", "2"}, []string{"GET_NEXT", "3"}, {"GET_NEXT", "4"},
		}, expected: []string{"", "", "", "", "2", "4", "4", "", "true", "2", "4", "4", ""}},
		{name: "case 3 ", input: [][]string{{"ADD", "1"}, {"ADD", "2"}, {"ADD", "2"}, {"ADD", "4"},
			{"GET_NEXT", "1"}, {"GET_NEXT", "2"}, {"GET_NEXT", "3"}, {"GET_NEXT", "4"},
		}, expected: []string{"", "", "", "", "2", "4", "4", ""}},
	}

	for _, test := range tests {

		result := File3(test.input)

		assert.Equal(t, test.expected, result)

	}
}
