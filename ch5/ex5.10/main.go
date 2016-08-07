// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import "fmt"

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//!-table

//!+main
func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
func contains(s []string, i string) bool {
	for _, a := range s {
		if a == i {
			return true
		}
	}

	return false
}
func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string][]string)

	visitAll = func(items map[string][]string) {

		for key, val := range items {
			if !seen[key] {
				order = append(order, key)
				visit := val
				for len(visit) > 0 {
					p := visit[len(visit)-1]
					visit = visit[:len(visit)-1]
					if !seen[p] {
						seen[p] = true
						order = append(order, p)
						visit = append(visit, m[p]...)
					}

				}

			}
		}

	}

	visitAll(m)
	return order
}

//!-main
