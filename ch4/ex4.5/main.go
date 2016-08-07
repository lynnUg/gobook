package main

import (
		"fmt"
)
func RemoveDup(s []string) []string{
	i :=0
	last :=""
	for _,z := range s {
		if last!=z{
			last = s[i]
			s[i]=z
			i++
			
		}
	}
	return s[:i]
}
func main() {
	a := []string{"a","a","a"}
	fmt.Println(RemoveDup(a))
}
