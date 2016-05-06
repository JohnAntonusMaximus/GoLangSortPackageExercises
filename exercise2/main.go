package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	printPeople(s)
}

func printPeople(p []string) {
	for _, person := range p {
		fmt.Printf("%v \n", person)
	}
}
