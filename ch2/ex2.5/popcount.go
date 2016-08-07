package main
import "fmt"
func PopCount(x uint64) int {
  n := 0
  for x!=0 {
    x = x & (x-1)
    n++
  }
  return n

}

func main() {
	values := []uint64{1,2,3,4,255}
	for _,value  :=range values {
		fmt.Printf("\n%d popcount is %d %[1]b\n" , value,PopCount(value))
	}
}
