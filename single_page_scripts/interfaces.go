package main

import (
	"fmt"
	"io"
	"os"
)

type triange struct {
	base   float64
	height float64
}

type square struct {
	side float64
}

type shape interface {
	getArea() float64
}

func (t triange) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func printArea(s shape) {
	fmt.Println("Area:", s.getArea())
}

func main() {
	var t = triange{base: 10, height: 5}
	var s = square{side: 7}
	printArea(t)
	printArea(s)

	//Read from text file and write to terminal
	fileArg := os.Args[1]
	file, _ := os.Open(fileArg)
	_, err := io.Copy(os.Stdout, file)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
