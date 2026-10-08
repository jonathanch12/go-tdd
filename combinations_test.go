package main

import (
	"reflect"
	"testing"
)

func TestCombine(t *testing.T) {
	testCases := []struct {
		name      string
		n         int
		k         int
		expected  [][]int
		expectErr bool
	}{
		{
			name:     "Basic case (n=4, k=2)",
			n:        4,
			k:        2,
			expected: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			name:     "Single element (k=1)",
			n:        4,
			k:        1,
			expected: [][]int{{1}, {2}, {3}, {4}},
		},
		{
			name:     "Full combination (n=k)",
			n:        4,
			k:        4,
			expected: [][]int{{1, 2, 3, 4}},
		},
		{
			name:     "Empty result (k=0)",
			n:        4,
			k:        0,
			expected: nil,
		},
		{
			name:      "Invalid input (n=0)",
			n:         0,
			k:         2,
			expectErr: true,
		},
		{
			name:     "Larger case (n=5, k=3)",
			n:        5,
			k:        3,
			expected: [][]int{{1, 2, 3}, {1, 2, 4}, {1, 2, 5}, {1, 3, 4}, {1, 3, 5}, {1, 4, 5}, {2, 3, 4}, {2, 3, 5}, {2, 4, 5}, {3, 4, 5}},
		},
		{
			name:      "k greater than n (n=3, k=5)",
			n:         3,
			k:         5,
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := combine(tc.n, tc.k)

			if tc.expectErr {
				if err == nil {
					t.Errorf("n = %d; k = %d; expected an error but got none (result = %v)", tc.n, tc.k, result)
				}
				return
			}

			if err != nil {
				t.Errorf("n = %d; k = %d; unexpected error: %v", tc.n, tc.k, err)
			}
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("n = %d; k = %d; result = %v; expect %v", tc.n, tc.k, result, tc.expected)
			}
		})
	}
}
