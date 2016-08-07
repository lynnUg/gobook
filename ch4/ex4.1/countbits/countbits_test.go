package countbits

import (
	"testing"
	"crypto/sha256"
)

func TestCountBits(t *testing.T) {
	cases := []struct {
		in1,in2 string
		want int
	}{
		{"x","X",125},
	}

	for _,c :=range cases {
		a := sha256.Sum256([]byte(c.in1))
		b := sha256.Sum256([]byte(c.in2))
		got := CountBits(a, b)
		if got != c.want {
			t.Errorf("CountBit(%s,%s) == %d but want %d" , c.in1, c.in2, got , c.want )
		}
	}
}
