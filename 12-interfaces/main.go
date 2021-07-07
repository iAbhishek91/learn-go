// interfaces defines methods signatures
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	height, width float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{0, 0, 5}
	rectangle := Rectangle{10, 12}

	fmt.Println(circle.area())
	fmt.Println(getArea(circle))
	fmt.Println(rectangle.area())
	fmt.Println(getArea(rectangle))
}
