package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := `[{"First":"Truong Thanh","Last":"Hai","Age":25},{"First":"My","Last":"Crush","Age":23}]`
	bs := []byte(s)

	fmt.Println(s)
	fmt.Printf("%T\n", bs)

	//people:=[]person{}

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("All of data", people)
	for _, v := range people {
		fmt.Println(v.Age)
		fmt.Println(v.First)
		fmt.Println(v.Last)
	}
}

type person struct {
	First string
	Last  string
	Age   int
}
