package main

import (
	"flag"
	"gopl.io/ch4/github"
	"log"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"bufio"
	"strings"
)
type Input struct {
	title string
	username string
	password string
	reponame string
	repoowner  string
	body string
	textEditor string
	num string

}
func  ReadIssue(x Input){
	result,err := github.ReadIssue(x.num, x.repoowner , x.reponame , x.username , x.password) 
	if err != nil {
		log.Fatal(err)
	}

	log.Println(result)
}
func CreateIssue(x Input) {
	if x.title == "" {
	  log.Fatal("need title")
	}

	body, err := ReadInput(x.textEditor)
	if err!= nil {
	  log.Printf("Error with reading input" , err)
	}
        x.body = body
	Item := github.Issue{Title:x.title, Body:x.body }
	result , err := github.CreateIssue(Item, x.repoowner , x.reponame , x.username , x.password) 
 
	if err!=nil {
		 log.Fatal(err)

	}

	log.Println(result)

}
func  ReadInput(textEditor string) (string ,error) {
	if textEditor == "" {
		textEditor = "vi"
	}

	tmpDir := os.TempDir()
	tmpFile , tmpFileErr := ioutil.TempFile(tmpDir,"tempFilePrefix")

	if tmpFileErr != nil  {
		return "", fmt.Errorf("Error %s creating tempFile \n",tmpFileErr)
	}

	path, err := exec.LookPath(textEditor)

	if err != nil {
		return "", fmt.Errorf("Error while looking for path %s of %s \n",path , textEditor)
	}

	cmd := exec.Command(path, tmpFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()

	if err != nil {
		return "", fmt.Errorf("Start failed %s \n" , err)
	}

	err = cmd.Wait()
	if err != nil {
		return "", fmt.Errorf("Command finished with error %s \n" ,err)
	}


	if  err != nil {
		return	"" , fmt.Errorf("Your file failed oh sorry %s ", err)
	}

	input := bufio.NewScanner(tmpFile)
	output := ""
	for input.Scan() {	
		a:= strings.TrimSpace(input.Text())
		output += a+"\n"
	}



	return output , nil
}


func main() {

	repoowner := flag.String("repoowner" ,"","repo owner")

	reponame  := flag.String("reponame","","repo name")

	username  := flag.String("username","","username")

	password  := flag.String("password", "", "password")

	title     := flag.String("title", "","title of issue")

	textEditor := flag.String("textEditor", "","text editor")

	create	   :=flag.String("create", "", "create issue")

	read       := flag.String("read", "" ,"read issue")

	num        := flag.String("num" , "","issue number")

	flag.Parse()
        
	if *create == "y" {
	createInput := Input{repoowner:*repoowner, reponame:*reponame ,username:*username,password:*password, title:*title, textEditor:*textEditor}
        CreateIssue(createInput)
} else if *read == "y" {
	readInput := Input{repoowner:*repoowner, reponame:*reponame ,username:*username,password:*password, num:*num}
	ReadIssue(readInput)
}


}
