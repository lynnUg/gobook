package comma

import (
	//"fmt"
	"bytes"
	"strings"
)
//add comma to decimal numbers and negative
func AddComma(s string) string {
	var buf bytes.Buffer

	if  strings.Contains(s, "-"){
		buf.WriteString("-")
		s = strings.TrimPrefix(s,"-")
	}
	
	var frac string
	if strings.Contains(s, ".") {
		parts := strings.SplitN(s,".",2)
		s ,frac = parts[0] , parts[1]
	}

	
	l := len(s)
	r := l % 3

	buf.WriteString(s[:r])
	if  r >=1 && l>3 {
	   buf.WriteString(",")

	}  

	for i :=r ; i<l ; i+=3 {
		buf.WriteString(s[i: i+3])
		if  i+3 <l {
			buf.WriteString(",")
		}
	}
	if frac !="" {	
		buf.WriteString(".")
		buf.WriteString(frac)
	}
	return buf.String()
	

}
