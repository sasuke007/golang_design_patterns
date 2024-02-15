package bridge

import "fmt"

type Share interface {
	draw()
}

type Triangle struct {
	color Color
}

func (t Triangle) draw() {
	fmt.Println("Drawing Triangle")
}

type Circle struct{}

func (c Circle) draw() {
	fmt.Println("Drawing Circle")
}

type Color interface {
	fillColor()
}

type Red struct{}

func (r Red) fillColor() {
	fmt.Println("Filling Red Color")
}

type Green struct{}

func (g Green) fillColor() {
	fmt.Println("Filling Green Color")
}

var triangleWithRedColor = Triangle{
	color: Red{},
}
