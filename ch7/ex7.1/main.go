package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	i := 0
	for scanner.Scan() {
		*c++
		i++
	}

	return i, nil

}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	i := 0
	for _, val := range p {
		if val == '\n' {
			*c++
			i++
		}
	}

	return i, nil
}
func main() {
	var c WordCounter
	c.Write([]byte("hello world"))
	fmt.Println(c)

	var w LineCounter
	w.Write([]byte("hello world \n helloworld"))
	fmt.Println(w)
}
