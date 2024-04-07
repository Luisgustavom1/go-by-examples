package sort_test

import (
	"reflect"
	"testing"

	s "github.com/Luisgustavom1/go-playground/benchmarks"
)

func TestSortSlice(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		output []int
	}{
		{
			name:   "AlreadySorted",
			input:  []int{1, 2, 3, 4, 5},
			output: []int{1, 2, 3, 4, 5},
		},
		{
			name:   "ReverseSorted",
			input:  []int{5, 4, 3, 2, 1},
			output: []int{1, 2, 3, 4, 5},
		},
		{
			name:   "RandomOrder",
			input:  []int{3, 1, 5, 2, 4},
			output: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s.Sort(tc.input)

			if !reflect.DeepEqual(tc.input, tc.output) {
				t.Errorf("Expected sorted slice: %v, but got: %v", tc.output, tc.input)
			}
		})
	}
}

// go test -bench=. -cpu=8 -benchmem -count 3 -benchtime=5s
func BenchmarkProcessSlice(b *testing.B) {
	sliceSize := 1000
	slice := s.GenerateRandomSlice(sliceSize)

	for i := 0; i < b.N; i++ {
		s.Sort(slice)
	}
}
