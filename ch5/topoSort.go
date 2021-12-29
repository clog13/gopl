package ch5

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithm": {"data structures"},
	"calculus":  {"linear algebra"},
}

func ToSort() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var res []string
	seen := make(map[string]bool)
	var visAll func(items []string)
	visAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visAll(m[item])
				res = append(res, item)
			}
		}
	}

	var graph []string
	for k := range m {
		graph = append(graph, k)
	}

	sort.Strings(graph)
	visAll(graph)
	return res
}
