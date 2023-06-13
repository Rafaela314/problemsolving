package others

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDaysOpen(t *testing.T) {

	tests := []struct {
		name     string
		input    []map[string]string
		expected []map[string]string
	}{

		{name: "case 1 ",
			input: []map[string]string{
				{
					"day":   "Monday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
				{
					"day":   "Tuesday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
				{
					"day":   "Wednesday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
				{
					"day":   "Thursday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
				{
					"day":   "Friday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
				{
					"day":   "Saturday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
				{
					"day":   "Sunday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
			},
			expected: []map[string]string{
				{
					"days":  "Monday - Sunday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
			},
		},

		{name: "case 2 ",
			input: []map[string]string{
				{
					"day":   "Monday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
				{
					"day":   "Tuesday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
				{
					"day":   "Wednesday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
				{
					"day":   "Thursday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
				{
					"day":   "Friday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
				{
					"day":   "Saturday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
			},
			expected: []map[string]string{
				{
					"days":  "Monday - Saturday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
				{
					"days":  "Sunday",
					"open":  "",
					"close": "",
				},
			},
		},
		{name: "case 3 ",
			input: []map[string]string{
				{
					"day":   "Monday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
				{
					"day":   "Tuesday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
				{
					"day":   "Wednesday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
				{
					"day":   "Thursday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
				{
					"day":   "Friday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
			},
			expected: []map[string]string{
				{
					"days":  "Monday - Friday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
				{
					"days":  "Saturday - Sunday",
					"open":  "",
					"close": "",
				},
			},
		},
		{name: "case 4 ",
			input: []map[string]string{
				{
					"day":   "Monday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
				{
					"day":   "Tuesday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
				{
					"day":   "Wednesday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
				{
					"day":   "Thursday",
					"open":  "8:00 AM",
					"close": "4:00 PM",
				},
				{
					"day":   "Friday",
					"open":  "8:00 AM",
					"close": "4:00 PM",
				},
			},
			expected: []map[string]string{
				{
					"days":  "Monday - Wednesday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
				{
					"days":  "Thursday - Friday",
					"open":  "8:00 AM",
					"close": "4:00 PM",
				},
				{
					"days":  "Saturday - Sunday",
					"open":  "",
					"close": "",
				},
			},
		},

		{name: "case 5 ",
			input: []map[string]string{
				{
					"day":   "Monday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
				{
					"day":   "Tuesday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
				{
					"day":   "Thursday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
				{
					"day":   "Friday",
					"open":  "8:00 AM",
					"close": "3:00 PM",
				},
				{
					"day":   "Saturday",
					"open":  "8:00 AM",
					"close": "3:00 PM",
				},
			},
			expected: []map[string]string{
				{
					"days":  "Monday - Tuesday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
				{
					"days":  "Wednesday",
					"open":  "",
					"close": "",
				},
				{
					"days":  "Thursday",
					"open":  "8:00 AM",
					"close": "5:00 PM",
				},
				{
					"days":  "Friday - Saturday",
					"open":  "8:00 AM",
					"close": "3:00 PM",
				},
				{
					"days":  "Sunday",
					"open":  "",
					"close": "",
				},
			},
		},
	}

	for _, test := range tests {

		result := DaysOpen(test.input)

		assert.Equal(t, test.expected, result)

	}
}
