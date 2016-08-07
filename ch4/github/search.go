// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"bytes"
	
)


func ReadIssue(num string, repoowner string, reponame string, USERNAME string, PASSWORD string) (bool, error){
	q := fmt.Sprintf(CRUDIssueURL,repoowner,reponame)
	req , err := http.NewRequest("GET" , q+"/"+num , nil)
	if err!=nil {
		return false, err
	}

	req.SetBasicAuth(USERNAME,PASSWORD)
	resp, err := http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return  false ,fmt.Errorf("read issue query failed: %s" , resp.Status)
	}
	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return false, err
	}
	resp.Body.Close()
	
	fmt.Println(result)

	return true , nil
}
func CreateIssue(Item Issue, repoowner string, reponame string, USERNAME string, PASSWORD string) (bool, error){
	q := fmt.Sprintf(CRUDIssueURL,repoowner,reponame)
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(Item)
	req , err := http.NewRequest("POST" , q , b)
	if err!=nil {
		return false, err
	}

	req.SetBasicAuth(USERNAME,PASSWORD)
	resp, err := http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusCreated {
		resp.Body.Close()
		return  false ,fmt.Errorf("create issue query failed: %s" , resp.Status)
	}

	return true , nil
}
// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	//!-
	// For long-term stability, instead of http.Get, use the
	// variant below which adds an HTTP request header indicating
	// that only version 3 of the GitHub API is acceptable.
	//
	//   req, err := http.NewRequest("GET", IssuesURL+"?q="+q, nil)
	//   if err != nil {
	//       return nil, err
	//   }
	//   req.Header.Set(
	//       "Accept", "application/vnd.github.v3.text-match+json")
	//   resp, err := http.DefaultClient.Do(req)
	//!+

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

//!-
