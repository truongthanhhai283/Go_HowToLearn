package main

import (
	"fmt"
	"sort"
)

type Node struct {
	key, value int
}

type ByKey []Node

func (s ByKey) Len() int { return len(s) }

func (s ByKey) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s ByKey) Less(i, j int) bool { return s[i].key < s[j].key }

func main() {

	nodes := []Node{
		{key: 1, value: 100},
		{key: 2, value: 200},
		{key: 3, value: 50},
	}

	sort.Sort(ByKey(nodes))
	fmt.Println(nodes)
}
