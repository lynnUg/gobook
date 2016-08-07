package main

import "fmt"

func PopCount(x uint64) int {
  n := 0
  for i :=uint(0) ; i<64 ; i++ {
    if x & (1 << i) != 0 {
      n++
    }
  }
  return n
}

func main() {
    values  := []uint64{1,2,3,4,255}
    for  _,value := range values {
	    fmt.Printf("\n %d is popcount is %d : %[1]b\n" ,value , PopCount(value))
    }
}
