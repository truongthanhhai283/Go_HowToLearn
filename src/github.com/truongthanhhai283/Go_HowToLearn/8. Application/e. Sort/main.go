package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{23, 35, 12, 6, 123, 4, 13}
	si2:=[]string{
		"sadsad",
		"ấdwdsd",
		"hấd",
		"sad",
		"214xf",
	}

	sort.Ints(xi)

	sort.Strings(si2)

	sort.Sort(sort.Reverse(sort.IntSlice(xi)))
	fmt.Println(xi)
	fmt.Println(si2)
}
