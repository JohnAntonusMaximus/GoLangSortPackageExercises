package main

import (
	"fmt"
	"sort"
)

type People []string

func (p People) Len() int {
	return len(p)
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type ByName struct {
	People
}

func (b ByName) Less(i, j int) bool {
	return b.People[i] < b.People[j]
}

func main() {
	studyGroup := People{
		"Zeno",
		"John",
		"Al",
		"Jenny",
	}
	sort.Sort(sort.StringSlice(studyGroup))
	printPeople(studyGroup)
}

func printPeople(p []string) {
	for _, person := range p {
		fmt.Printf("%v \n", person)
	}
}
