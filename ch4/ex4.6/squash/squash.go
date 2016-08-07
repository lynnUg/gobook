package squash

import (
	"unicode"
	"unicode/utf8"
)


func Squash(s []byte)[]byte {
	i := 0
	isLastSpace := false
	m := string(s)
	
	for _, c := range m  {
		if unicode.IsSpace(c) || c == ' '{
			if !isLastSpace {
				c = ' '
				isLastSpace = true  
				
			} else {
				continue
			}

		} else {
			isLastSpace = false
		}

		if utf8.RuneLen(c) > 1{
			u := make([]byte, 3)
			n := utf8.EncodeRune(u , c)
			j := i + n
			copy(s[i:j],u)
			i = j
		} else {
			s[i] = byte(c)
			i++
	       }
	}

	return s[:i]

}
