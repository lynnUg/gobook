package xkcd

import (
	"strings"
	"io/ioutil"
	"log"
	"encoding/json"
        "fmt"
	"net/http"
)

func Search ( p []Page, x string) []Page{
	var output []Page
	for _,item := range p {
		if strings.Contains(item.Transcript, x) {
			output = append(output,item)
		}
	}

	return output
}

func GetPages() []Page {
	input , err := ioutil.ReadFile(db)

	if err != nil {
		log.Fatal(err)
	}

	var c []Page
	json.Unmarshal(input, &c)

	return c
}

func AddPages(i int) (bool,error) {
	pages := GetPages()

        req, err := http.NewRequest("GET", fmt.Sprintf(XkcdURL,i),nil)

	if err != nil {
	  return false , err
        }
        resp, err := http.DefaultClient.Do(req)
        if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return false , fmt.Errorf("Adding pages failed:  %s" ,resp.Status)
	}

        var result Page
         if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resp.Body.Close()
			return false , err
	 }

         resp.Body.Close()
         pages = append(pages , result)
         b, err := json.Marshal(pages)
	 if err != nil {
		log.Println(err)
          }

	 if err := ioutil.WriteFile(db, b, 0644) ; err != nil {
		return false , err
         }
        return true , nil
}
