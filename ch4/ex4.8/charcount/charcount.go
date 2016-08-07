package charcount

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"

)

func CharCount(in *bufio.Reader) map[string]int {
	counts := make(map[string]int)
	

	

	for {
		r, n , err := in.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr,"charcount: %v \n" ,err)
			os.Exit(1)
		}
		 
		if r == unicode.ReplacementChar && n == 1 {
	
		        counts["invalid"]++
			continue
		}


		switch {
		case unicode.IsControl(r):
			counts["control"]++
		case unicode.IsLetter(r):
			counts["letter"]++
		case unicode.IsMark(r):
			counts["mark"]++
		case unicode.IsNumber(r):
			counts["number"]++
		case unicode.IsPunct(r):
			counts["punct"]++
		case unicode.IsSpace(r):
			counts["space"]++
		case unicode.IsSymbol(r):
			counts["symbol"]++
		}
		
	}

    
    return counts
}
