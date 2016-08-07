package main

import (
	"fmt"
)

func rotate(moves int, s []int, right bool) []int {
	if  moves == 0 || len(s) ==0 {
	 return s
	}


	x := moves % len(s)
	y := len(s) - x
	if right {
		x, y = y, x
	}
	temp := make([]int, len(s))
	copy(temp ,s[x:])
	copy(temp[y:],  s[:x])
	return temp

}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	
	t := rotate(2, s , true)
	fmt.Println(t)
	f := rotate(2, s , false)
	fmt.Println(f)
}
