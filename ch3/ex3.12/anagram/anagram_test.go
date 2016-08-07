package anagram

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	cases := []struct{
		in1 , in2 string  
		want bool 
	}{
		{"msas", "mass",  true},
	}

	for _, c := range cases {
		got := IsAnagram(c.in1,c.in2)
		if  got != c.want {
			t.Errorf("IsAnagram(%s , %s) != %v got %v", c.in1, c.in2, c.want, got)
		}
	}
}
