package reverse

import (
	"testing"
	"bytes"
)

func TestReverse(t *testing.T) {
	cases := []struct{
	 in , want []byte
	}{
		{[]byte("foo bar"), []byte("rab oof")},
		{[]byte("Hello, 世界"), []byte("界世 ,olleH")},
	}
	for _, c := range cases {
		if Reverse(c.in); !bytes.Equal(c.in,c.want){
			t.Errorf("got %s want %s",c.in , c.want)
		}
	}
}
