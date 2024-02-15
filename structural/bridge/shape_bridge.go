package bridge

import "fmt"

type Share interface {
	draw()
}

type Triangle struct {
	color Color
}

func (Triangle) draw() {
	fmt.Println("Drawing Triangle")
}

type Circle struct{}

func (Circle) draw() {
	fmt.Println("Drawing Circle")
}

type Color interface {
	fillColor()
}

type Red struct{}

func (Red) fillColor() {
	fmt.Println("Filling Red Color")
}

type Green struct{}

func (Green) fillColor() {
	fmt.Println("Filling Green Color")
}

var triangleWithRedColor = Triangle{
	color: Red{},
}
