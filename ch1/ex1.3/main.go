
package main
import (
	"fmt"
	"time"
	"os"
	"strings"
)
func echo1(input []string){
	s, sep := "" , ""
	for _, arg := range input {
		s +=sep + arg
		sep = " "
	}
	fmt.Println(s)
}
func echo3(input []string) {
	fmt.Println(strings.Join(input," "))
}

func main() {
	t1 := time.Now()
        echo1(os.Args[1:])	
	t2 := time.Now()
        fmt.Printf("it took echo 1 this long %v \n", t2.Sub(t1))
        t3 := time.Now()
	echo3(os.Args[1:])
	t4 := time.Now()
	fmt.Printf("it took echo 3 this long %v \n" , t4.Sub(t3))
}
