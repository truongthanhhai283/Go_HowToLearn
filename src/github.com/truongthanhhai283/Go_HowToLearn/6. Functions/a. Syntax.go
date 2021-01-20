package main

import "fmt"

func main() {
	foo()

	//1 returns
	bar("Hai")
	x := woo("Truong")
	fmt.Println(x)

	//multiple returns
	x, y := demo("Hai", "Truong")
	fmt.Println(x, y)
}

//func (r receiver) identifier(parameters) (return(s)) {....}
func foo() {
	fmt.Println("Foo func")
}

//everything in go is pass by value
func bar(s string) {
	fmt.Println("Hello,", s)
}

//1 return
func woo(s string) string {
	return fmt.Sprintf("Hello from woo %s", s)
}

//multiple returns
func demo(fn string, ln string) (string, bool) {
	a := fmt.Sprint(ln, " ", fn, " ", "say", " ", "Hello")
	b := false
	return a, b
}
