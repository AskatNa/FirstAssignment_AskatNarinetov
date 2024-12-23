package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Length
}
func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Length)
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Square struct {
	Length float64
}

func (s *Square) Area() float64 {
	return s.Length * s.Length
}
func (s *Square) Perimeter() float64 {
	return s.Length * 4
}

type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

func (t *Triangle) Area() float64 {
	heron := (t.SideA + t.SideB + t.SideC) / 2
	return math.Sqrt(heron * (heron - t.SideA) * (heron - t.SideB) * (heron - t.SideC))
}
func (t *Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}
func PrintShapeDetails(s Shape) {
	fmt.Printf("Shape Detailes: \n")
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter: ", s.Perimeter())
}

func main() {
	shapes := []Shape{
		&Rectangle{Length: 5, Width: 10},
		&Circle{Radius: 7},
		&Square{Length: 4},
		&Triangle{SideA: 5, SideB: 6, SideC: 8},
	}
	for _, shape := range shapes {
		PrintShapeDetails(shape)
	}
}
