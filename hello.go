package main

import "fmt"

func main()  {
	//declaring a variable and assigning it later
	var a string
	a = "Aanisha"
	//declaring a variable as well as assigning it simultaneously
	b := 6
	//printing a line by concatenating strings and integers
	fmt.Println("Hello World!",a, b)

	//declaring a constant and printing it
	const y = "y is constant"
	fmt.Println(y)
}