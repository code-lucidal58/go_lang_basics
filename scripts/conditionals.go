package main

import (
	"fmt"
	"time"
)

func main() {

	//For statements
	z := 1
	for z < 5 {
		fmt.Println(z)
		z = z + 1
	}
	// if and for combinations
	for i := 10; i > 6; i-- {
		if i%2 == 0 {
			fmt.Println(i, " is an even number")
		} else {
			continue
		}
	}

	//switch statements
	t := 4
	switch t { //one with expression
	case 3:
		fmt.Println("three")
		t++
	case 4:
		fmt.Println("four")
		t--
	}
	switch { //one without expression
	case t > 5:
		fmt.Println("big ", t)
	case t < 5:
		fmt.Println("small ", t)

	}

	t_time := time.Now().Weekday()
	switch t_time {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Its a weekday")
	default:
		fmt.Println("Its a weekend")
	}

	//defer function statements
	for i:=1; i<5;i++{
		defer fmt.Println(i)
	}
	fmt.Println("Finished counting:")

	//lets see some miracle with defer
	defer func (){
		str:= recover()
		fmt.Println(str)
	}()

	panic("Hello There")
	//panic and recover are predefined functions similar to try and catch
}
