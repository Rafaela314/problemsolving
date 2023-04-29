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
		//{name: "case 1 ", s: "CAGCCTATGC", p: []int{2, 5, 0}, q: []int{4, 5, 6}, expected: []int{2, 4, 1}},
	}

	for _, test := range tests {

		result := GenomicRangeQuery(test.s, test.p, test.q)

		assert.Equal(t, test.expected, result)

		//if result != test.expected {
		//	t.Errorf("FAIL --- %s, expected '%v', received '%v'", test.name, test.expected, result)
		//}
	}

}
