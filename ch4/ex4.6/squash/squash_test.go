package squash

import (
	"testing"
	"bytes"
)

func TestSquash(t *testing.T) {
  cases := []struct {
	  in , want []byte
  } {
   {[]byte("x\t\t\t世\t\t\ty \v\v\n  z"), []byte("x 世 y z") },
  }

  for _, c := range cases {
	//got := Squash(c.in)
	//fmt.Println(got)
	if got := Squash(c.in) ; !bytes.Equal(got,c.want) {
		t.Errorf(" got %s expected %s",got, c.want)
	}
  }
}
