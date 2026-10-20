package main

import "testing"

func Contains(numbers []int, target int) bool {
	for _, num := range numbers {
		if num == target {
			return true
		}
	}
	return false
}

func TestContains(t *testing.T) {
	cases := []struct {
		name string
		numbers []int
		target int
		want bool
	} {
		{
			name: "success contains",
			numbers: []int{1, 2, 3},
			target: 3,
			want: true,
		},
		{
			name: "not success contains",
			numbers: []int{1, 2, 3},
			target: 4,
			want: false,
		},
		{
			name: "nil slice contains",
			numbers: nil,
			target: 3,
			want: false,
		},
		{
			name: "negative contains",
			numbers: []int{3, 2, -3},
			target: -3,
			want: true,
		},
		{
			name: "equal contains",
			numbers: []int{1, 1, 1},
			target: 1,
			want: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := Contains(c.numbers, c.target)
			if result != c.want {
				t.Fatalf("Ожидали %t, а пришло %t", c.want, result)
			}
		})
	}
}