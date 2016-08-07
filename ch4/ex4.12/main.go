package main

import (
//	"fmt"
	"gopl.io/ch4/xkcd"
	"flag"
	"log"
)
func DownloadJson() {
	for i:=1 ; i<570 ; i++ {
	 xkcd.AddPages(i)
	}
}
func main() {
	SearchTerm := flag.String("s","","search term")
	flag.Parse()
	if  *SearchTerm == "" {
		log.Fatal("search term cant be empty string")
	}
	pages := xkcd.GetPages()
	xkcd.Search(pages , *SearchTerm)

}
