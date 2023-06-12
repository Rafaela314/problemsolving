package codility

import (
	"testing"

	"rand"

	"github.com/stretchr/testify/assert"
)

func TestOddOccurrencesInArray(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{name: "case 1 ", input: []int{9, 3, 9, 3, 9, 7, 9}, expected: 7},
		{name: "case 2 ", input: []int{3, 4, 3, 4, 6, 1, 4, 4, 6}, expected: 1},
		{name: "case 3 ", input: []int{15, 3, 4, 3, 4, 6, 1, 4, 1, 15, 4, 6, 22}, expected: 22},
	}

	for _, test := range tests {

		actual := OddOccurrencesInArray(test.input)
		assert.Equal(t, test.expected, actual)

	}
}
func insertXSlice(itemsCount int, b *testing.B) {
	testSlice := []int{}

	for i := 0; i < itemsCount; i += 1 {
		testSlice = append(testSlice, i)
	}
}

func createArray(arraySize int) []int {
	testArray := make([]int, arraySize)

	for i := 1; i <= arraySize; i += 1 {
		testArray[i] = i
	}

	return testArray
}

/*
func BenchmarkOddOccurrencesInArray(b *testing.B) {

	rand.Seed(int(b.N))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		OddOccurrencesInArray(num)
	}
}*/
