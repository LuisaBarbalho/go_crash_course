package main

import (
	"fmt"
	"math"
)

// Shape interface
type Shape interface {
	area() float64
	perimeter() float64
}

// Circle struct
type Circle struct {
	x, y, radius float64
}

// Rectangle struct
type Rectangle struct {
	wiidht, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r Rectangle) area() float64 {
	return r.wiidht * r.height
}

func (r Rectangle) perimeter() float64 {
	return (2 * r.wiidht) + (2 * r.height)
}

func getArea(s Shape) float64 {
	return s.area()
}

func getPerimeter(s Shape) float64 {
	return s.perimeter()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{wiidht: 10, height: 5}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Circle Perimeter: %f\n", getPerimeter(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))
	fmt.Printf("Rectangle Perimeter: %f\n", getPerimeter(rectangle))

}
