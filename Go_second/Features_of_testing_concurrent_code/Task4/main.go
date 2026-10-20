package main

import "testing"

func AreAnagrams(str1, str2 string) bool {
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	if len(str1) != len(str2) {
		return false
	}

	// Convert strings to slices of runes for sorting
	r1 := []rune(str1)
	r2 := []rune(str2)

	sort.Slice(r1, func(i, j int) bool {
		return r1[i] < r1[j]
	})

	sort.Slice(r2, func(i, j int) bool {
		return r2[i] < r2[j]
	})

	return string(r1) == string(r2)
}

func TestAreAnagrams(t *testing.T) {
	cases := []struct {
		name 	string
		str1  	string
		str2	string
		want 	bool
	}{
		{
			name: "success anagrams",
			str1:  "dot",
			str2:  "odt",
			want: true,
		},
		{
			name: "empty anagrams",
			str1:  "",
			str2: "",
			want: true,
		},
		{
			name: "not success anagrams",
			str1:  "doe",
			str2: "eor",
			want: false,
		},
		{
			name: "equal anagrams",
			str1:  "ddd",
			str2: "ddd",
			want: true,
		},
		{
			name: "symbol anagrams",
			str1:  "(-;@){",
			str2: "{)@;-(",
			want: true,
		},
		{
			name: "not equal len anagrams",
			str1:  "ee",
			str2: "rrr",
			want: false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := AreAnagrams(c.str1, c.str2)
			if result != c.want {
				t.Errorf("Ожидали %t, а получили %t", c.want, result)
			}
		})
	}
}