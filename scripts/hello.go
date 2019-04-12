package main

import "fmt"

func changeByValue(y string) {
	y = "This was passed by value"
}

func changeByRef(y *string) {
	*y = "This was passed by reference"
}

// multiple return
func multipleReturnFunction(x string) (string, int) {
	return "It was good", len(x)
}

//naked return
func nakedReturnFunction(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	//declaring a variable and assigning it later
	var a string
	a = "Aanisha"
	//declaring a variable as well as assigning it simultaneously
	b := 6
	//printing a line by concatenating strings and integers
	fmt.Println("Hello World!", a, b)

	//declaring a constant and printing it
	const y = "y is constant"
	fmt.Println(y)

	xx, yy := nakedReturnFunction(45)
	fmt.Println("Naked return function executed: ", xx,yy)

	//Introduction to pointers
	s := "Lets check what pointer is"
	changeByValue(s)
	fmt.Println("Changed by value: ",s)

	changeByRef(&s)
	fmt.Println("Changed by reference: ",s)

	//Pointer declaration
	x := new(string)
	changeByRef(x)
	fmt.Println(*x)

	fmt.Println(len(*x))

	var boolean bool
	fmt.Printf("The type of variable boolean is %T\n", boolean)

	//multiple return from functions
	str := "How was your day?"
	sx, sy := multipleReturnFunction(str)
	fmt.Println(sx, sy)
}
