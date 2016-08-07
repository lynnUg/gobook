package main

import (
//	"fmt"
	"log"
	"os"
	"gopl.io/ch4/github"
//	"gopl.io/ch4/github/Issue"
	"time"
)


func main() {
	results, err := github.SearchIssues(os.Args[1:])
        //        log.Printf("%T", results)
	if err != nil {
	  log.Fatal(err)
	}
	categories := make(map[string][]*github.Issue)
	log.Printf("Total count of results %d %d", results.TotalCount, len(results.Items))
        monthHours := 30.0 * 24
	yearHours := monthHours * 12
	today := time.Now()
	for _,item := range results.Items {
		diff := today.Sub(item.CreatedAt)
		if diff.Hours() < monthHours {
		  categories["<month"]= append(categories["<month"],item)
		} else if diff.Hours() < yearHours {
		categories["<year"]= append(categories["<year"], item)
		} else  {
		categories["year>"]= append(categories["year>"], item)
		}
	}

	for key, value := range categories {
	     log.Printf("%s has a count %d",key , len(value))	
      }
}
