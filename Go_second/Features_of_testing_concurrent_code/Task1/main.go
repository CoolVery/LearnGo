package main

import (
	"slices"
	"testing"
)

func  SortIntegers(numbers []int) 

func TestSortNum(t *testing.T) {
	cases := []struct {
		name string
		values []int
		workSlice []int
		want []int
	} {
		{
			name: "positive sort",
			values: []int{4, 1, 5, 6},
			workSlice: []int{4, 1, 5, 6},
			want: []int{1, 4, 5, 6},
		},
		{
			name: "negative sort",
			values: []int{-4, -1, -5, -6},
			workSlice: []int{-4, -1, -5, -6},
			want: []int{-6, -5, -4, -1},
		},
		{
			name: "equal sort",
			values: []int{1, 1, 1, 1},
			workSlice: []int{1, 1, 1, 1},
			want: []int{1, 1, 1, 1},
		},
		{
			name: "sort with double",
			values: []int{3, 1, 5, 2, 4, 4},
			workSlice: []int{3, 1, 5, 2, 4, 4},
			want: []int{1, 2, 3, 4, 4, 5},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			SortIntegers(c.workSlice)
			if !slices.Equal(c.workSlice, c.want) {
				t.Errorf("Ожидали %v, а получили %v", c.want, c.workSlice)
			}
		})
	}
}