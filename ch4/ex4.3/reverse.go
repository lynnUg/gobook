package main

import (
	"fmt"
)

func reverse(a *[4]int) {
	for i,j := 0, len(a)-1; i<j; i,j= i+1, j-1 {
	    a[i] , a[j] = a[j] , a[i]
	}
}

func main() {
	a := [...]int{0,1,2,3}
	reverse(&a)
	fmt.Println(a)
}
