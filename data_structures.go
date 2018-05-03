package main

import "fmt"

// A structure
type Rectangle struct {
	width  int
	length int
}

//Functions with unknown number of arguments
//Variadic function
func addNumbers(args ...int)int{
	sum := 0
	for _, n:= range args {
		sum+=n
	}
	return sum
}
func main() {

	//Arrays
	var arr [2]string
	arr[0] = "Hello"
	arr[1] = "Aanisha"
	fmt.Println(arr)

	//Slices - Method 1
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	//Slices - Method 2
	g := make([]int, 5, 10)
	s := g[2:4] //again slices using high low expression
	s[0] = 10
	s[1] = 11
	g[4] = 14
	fmt.Println(s, " capacity of s is ", cap(s))
	a := s[:cap(s)]
	//size of a will be 8 as it will take up complete capacity of s and in turn g
	fmt.Println(a)

	var pow = []int{1, 2, 4, 8, 16}
	//range keyword is used to loop through each element in the slice pow
	//n represents the index
	//p represent the value in the index n
	for n, p := range pow {
		fmt.Printf("%d indexed value is %d\n", n, p)
	}

	//creating a new slice using one created earlier
	new_slice := append(pow, 32, 64)
	fmt.Println("The new slice is ", new_slice)

	//copy one slice to another
	copied_slice := make([]int, 2, 4)
	copy(copied_slice, pow)
	fmt.Println("Copied slice is ", copied_slice)

	nil_slice1 := make([]int, 0, 0)
	nil_slice2 := make([]int, 0, 4)
	var nil_slice3 []int
	switch {
	case nil_slice1 == nil:
		fmt.Println("nil_slice1 is nill")
	case nil_slice2 == nil:
		fmt.Println("nil_slice2 is nill")
	case nil_slice3 == nil:
		fmt.Println("nil_slice3 is nill")
	default:
		fmt.Println("None of the slices are nil!!")
	}

	//Demonstration of use of structures
	r := Rectangle{10, 20}
	fmt.Println(r, "Here width: ", r.width, " and length: ", r.length)
	pr := &r
	pr.width = 15
	r.length = 30
	fmt.Println("pr: ", pr, " r: ", r)

	//Maps
	var m = map[string]Rectangle{"Rect1": {1, 2}, "Rect2": {3, 4}}
	fmt.Println(m)
	delete(m, "Rect1")
	x, ok := m["Rect1"]
	fmt.Println("Value of Rect1: ", x, " Present: ", ok)
	fmt.Println("Finally m is: ", m)

	// calling function that take n number of arguments
	fmt.Println("sum from 1 to 6 is ",addNumbers(1,2,3,4,5,6))
}
