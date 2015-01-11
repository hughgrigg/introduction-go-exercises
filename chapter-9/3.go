/*
Add a new method to the Shape interface called perimeter which calculates the
perimeter of a shape. Implement the method for Circle and Rectangle.
*/

package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return math.Pi * 2 * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) width() float64 {
	return math.Abs(r.x1 - r.x2)
}

func (r *Rectangle) height() float64 {
	return math.Abs(r.y1 - r.y2)
}

func (r *Rectangle) area() float64 {
	return r.width() * r.height()
}

func (r *Rectangle) perimeter() float64 {
	return (r.width() + r.height()) * 2
}

func main() {
	circle := Circle{x: 0, y: 0, r: 5}
	rectangle := Rectangle{x1: 0, y1: 0, x2: 10, y2: 5}
	fmt.Printf("Circle circumference: %.3f\n", circle.perimeter())
	fmt.Printf("Rectangle perimeter: %.3f\n", rectangle.perimeter())
}
