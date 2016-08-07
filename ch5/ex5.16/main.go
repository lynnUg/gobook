package main

import (
	"fmt"
)

func Join(sep string, words ...string) (out string) {

	for i, word := range words {
		out += word
		if i != len(words)-1 {
			out += sep
		}
	}

	return
}

func main() {
	fmt.Println(Join(",", "A", "b", "c"))
}
