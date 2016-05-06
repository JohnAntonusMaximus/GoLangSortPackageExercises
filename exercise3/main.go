package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	printNumbers(n)
}

func printNumbers(n []int) {
	for _, number := range n {
		fmt.Printf("%v \n", number)
	}
}
