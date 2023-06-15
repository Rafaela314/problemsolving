package codility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenomicRangeQuery(t *testing.T) {

	tests := []struct {
		name     string
		s        string
		p        []int
		q        []int
		expected []int
	}{
		{name: "case 1 ", s: "CAGCCTA", p: []int{2, 5, 0}, q: []int{4, 5, 6}, expected: []int{2, 4, 1}},
		{name: "case 2 ", s: "CAGCCTATGC", p: []int{2, 5, 0}, q: []int{4, 5, 6}, expected: []int{2, 4, 1}},
		{name: "case 3 ", s: "GT", p: []int{0, 0, 1}, q: []int{0, 1, 1}, expected: []int{3, 3, 4}},
		{name: "case 4 ", s: "AAAAAAAAAAAAAAAAAAAAAA", p: []int{0, 0, 1}, q: []int{0, 1, 1}, expected: []int{1, 1, 1}},
		{name: "case 4 ", s: "CCCTTTTTT", p: []int{0, 0, 1}, q: []int{0, 1, 5}, expected: []int{2, 2, 2}},
	}

	for _, test := range tests {

		result := GenomicRangeQuery(test.s, test.p, test.q)

		assert.Equal(t, test.expected, result)

	}

}
