package main

import (
	"math"
	"fmt"
)

// Structure
type Rect struct{
	 length, breadth float64
}

// Method
func (r *Rect) diagonal() float64{
	return math.Sqrt(r.length * r.length + r.breadth * r.breadth)
}

func main(){
	r := Rect{10,20}
	fmt.Println(r.diagonal())
}