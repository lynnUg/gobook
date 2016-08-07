package  countbits

import (
//	"crypto/sha256"
//	"fmt"
)
func CountBits(a, b [32]byte) int {
//	fmt.Println(a,b)
//	fmt.Printf("%x\n%x\n%T\n",a,b,a)
	var n int
	for  i:=0 ; i<32 ; i++ {
		x := a[i] ^ b[i]
		for x!=0 {
			x &= (x-1)
			n++
		}
	}
	return n
}
