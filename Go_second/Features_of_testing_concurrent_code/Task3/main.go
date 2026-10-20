package main

import "testing"

func ReverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func TestReverseString(t *testing.T) {
	cases := []struct {
		name    string
		str string
		want    string
	}{
		{
			name:	"success reverse",
			str: 	"dwrt",
			want:	"trwd",
		},
		{
			name:	"empty reverse",
			str: 	"",
			want:	"",
		},
		{
			name:	"one word reverse",
			str: 	"d",
			want:	"d",
		},
		{
			name:	"equal reverse",
			str: 	"ddd",
			want:	"ddd",
		},
		{
			name:	"symbol reverse",
			str: 	"(-;@){",
			want:	"{)@;-(",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := ReverseString(c.str)
			if result != c.want {
				t.Errorf("Ожидали %s, а получили %s", c.want, result)
			}
		})
	}
}
