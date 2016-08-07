package main
import "fmt"
var pc [256]byte

func init() {
  for i := range pc {
    pc[i] = pc[i/2] + byte(1&i)
  }
}

func PopCount(x uint64) int {
  var n byte
  fmt.Println(x)
  for i:=uint(0) ; i < 8 ; i++ {
    n += pc[byte(x>>i*8)]
  }
  return int(n)
}

func main() {
	for _, x := range []uint64{0,1,2,3,255,0x1234567890ABCDEF} {
		fmt.Printf("\n%8d has a popcount of %d : %8[1]b\n",x,PopCount(x))
	}
}


