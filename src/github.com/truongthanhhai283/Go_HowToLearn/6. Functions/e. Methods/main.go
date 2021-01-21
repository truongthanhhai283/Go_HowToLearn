package main

import "fmt"

//func (r receiver) indentity(params) (return(s)) {
//	.....code
//}

type vehicle struct {
	wheel int
	color string
}

type car struct {
	vehicle
	windows int
}

//method
func (c car) show() {
	fmt.Println("This's", c.color, "car,", c.wheel, "wheels")
}

func main() {
	c := car{
		vehicle: vehicle{
			wheel: 4,
			color: "red",
		},
		windows: 2,
	}

	c.show()

	fmt.Println(c)
}
