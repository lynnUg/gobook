package charcount

import (
	"testing"
	"strings"
	"bufio"
)
var tests = []struct{
	input string
	want map[string]int
}{
	{"",map[string]int{}},
	{"世界", map[string]int{"letter": 2}},
	{"\x90", map[string]int{"invalid": 1}},
	{"hello, world!", map[string]int{
		"space":  1,
		"letter": 10,
		"punct":  2,
	}},
        
}
func TestCharCount(t *testing.T) {
	for  _, test := range tests {
		r := strings.NewReader(test.input)
		got := CharCount(bufio.NewReader(r))
		if !MapCompare(got, test.want) {
			t.Errorf("got %v expected %v",got, test.want)
		}
		 
	}
}

func MapCompare(x , y map[string]int) bool {
	if len(x) !=len(y) {
		return false
	}

	for key, value :=range x {
		if  a,ok := y[key] ; !ok || a != value {
		return false
	 }
	}

	return true
}
