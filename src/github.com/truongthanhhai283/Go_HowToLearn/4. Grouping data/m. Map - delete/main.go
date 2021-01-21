package main

import "fmt"

func main() {
	x := map[string]int{
		"Hai":  30,
		"Bob":  40,
		"Jack": 50,
	}

	x["todd"] = 33

	if _, ok := x["todd"]; ok {
		delete(x, "lisa")
		fmt.Println(x)
	}else {
		fmt.Println(false)
	}


}
