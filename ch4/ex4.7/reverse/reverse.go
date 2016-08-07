package reverse

import (
	"unicode/utf8"
)

func Reverse( input []byte)  {
	var p , q rune
	var pz , qz int
	for i, j := 0 , len(input) ; i < j-1 ; i, j= i+qz, j-pz {
		p , pz = utf8.DecodeRune(input[i:])
		q , qz = utf8.DecodeLastRune(input[:j])
		
	
	       if qz > pz {
				copy(input[i+qz:], input[i+pz:j-qz])
	       }

	       copy(input[i:],[]byte(string(q)))
	       copy(input[j-pz:],[]byte(string(p)))
	}
	
}
