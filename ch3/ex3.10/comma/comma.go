package comma

import (
	//"fmt"
	"bytes"
)

func AddComma(x string) string {
	if len(x) <=3 {
		return x
	}
	
	var buf bytes.Buffer
	l := len(x)
	r := l % 3

	if  r >=1 {
	   buf.WriteString(x[:r])
	   buf.WriteString(",")

	}  

	for i :=r ; i<l ; i+=3 {
		buf.WriteString(x[i: i+3])
		if  i+3 <l {
			buf.WriteString(",")
		}
	}

	return buf.String()


}
