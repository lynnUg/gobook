package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(expand("test,$foo", the))
}

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "$foo", f("foo"), -1)
}

func the(s string) string {
	return "the"
}
