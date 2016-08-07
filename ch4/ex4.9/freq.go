package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
    count := make(map[string]int)
    file, err := os.Open("text.txt")

   if err != nil {
	fmt.Fprintf(os.Stderr,"the error %v\n", err)
	os.Exit(1)
  }
  defer file.Close()
  
  input := bufio.NewScanner(os.Stdin)
  input.Split(bufio.ScanWords)
  
  for input.Scan(){
	word := input.Text()
	count[word]++
  }

  if  err := input.Err(); err != nil {
	fmt.Fprintf(os.Stderr,"%v\n",err)
  }
   
  for key , value := range count {
	fmt.Printf("%s\t%d\n",key ,value)
  }
}
