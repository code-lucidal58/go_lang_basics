# GoLangBasics

### Introduction to the language
Go is a general purpose programming language. It is open source. Due to its concurrency feature it is ideal for 
multi-core and networked machines. Best suited for modular software designs. Compiles very quickly and performs 
runtime reflection. Developed by Google in 2007, announced in 2009 and stable version released in 2015. 
Go is a case sensitive language.

### Lets GO
Every Go program is made up of packages. If program starts with _main_ package, it is an executable file.<br>
Package name is the same as the last element of the import path. 
For instance, the _math/rand_ package comprises files that begin with the statement _package rand_. 
To run a go file:  `go run file_name.go`

### Import statements
**Factored Import statements**: 
```go
import (
	"fmt"
	"math"
)
```
This is preferred/ convention way of including packages.<br>
**Non factored Import statements**: 
```go
import "fmt"
import "math"
```
Some key packages in Go: strings, io, bytes, os, path/filepath, errors, container/list, hash, encoding/gob<br>
A name is exported from another package if it starts with capital letter.
```go
fmt.Println(math.Pi)
```
_Pi_ is a name exported by package _math_. 

### Variables
Variables in Go have the same convention as most other languages (starts with letter, contains numbers and underscore)<br>
Single Variable declaration: `var a bool`
Double Variable declaration: `var a,b,c bool`
Variable Naming conventions: camelCase <br>
Assigning value to variable: 
```go
a := ”Aanisha” //creates new variable and assigns
a = “Aanisha” // assigns existing variable a
```
If variable is not initialised after declaration, it is initialised to 0 if int, false if bool and “” if string<br>

**Types in Go**: uint8, uint16, uint32, uint64 (unsigned integers, memory located as per size mentioned in the name), 
int8, int16...int64(signed integers). If just int is used, it becomes platform dependent. For floating point number, 
float32, float64 : complex32, complex64 for storing complex numbers. Bool: true, false. 

### Pointers
Pointers have same logic as in C++. Parameters passed to a function are by default passed by value.
`&var_name` will give the address of the variable var_name. To dereference a pointer, * is used. Suppose var_pointer 
is a pointer. `*var_pointer` will give the value in the address stored by pointer var_pointer. 
To declare a pointer, new  keyword is used. 
```go
y : = new(int)
```
This will allocate enough memory to store an integer value and assign the memory address to y.

### Functions
Functions are defined using keyword func. The parameters must have the type specified alongside. 
Functions can return value using _return_ keyword. 
```go
E.g. func function_name(x int, y int) int {return x+y}
```
The int  before curly braces states that this function is going to return an int value. `(x int, y int) int`
This part in the above function is called **signature**.

String length is found using len function: `len(string_variable)`
_%T_ is a format specifier which returns the type of a variable.
```go
fmt.Printf(“Type of the variable is%T”,var_name)
``` 

### Conditionals
Refer `conditionals.go` in this repository for practical usages<br>
**For** statements are similar to _while_ statements in C++. Iterator is modified inside the _for_ block statements. 
Don’t worry, the normal version is also supported.  `for i:= 0; i<10;i++{}`. **If else** logic is same as C++, 
nested if, else if are also supported.<br>
_**NOTE**_: Parentheses are required even if _if_ block has just one statement. Same applies for other conditionals.

There are two types of switch statements:
* Type switch:  `switch <expression> {}`
* Expression switch: `switch {} //no expression`

**Defer functions**: Defer statement defers the execution of function. Function calls are stacked up and called in the 
sequence of last in first out. Just state defer before making any function call. That function will be executed 
after all statements are executed in main.

### Data Structures
Refer `data_structures.go` in this repository for practical usages<br>
**Arrays**: length of array must be specified, it can be a variable. An array type once declared cannot be changed.
```go
var arr [4][8] int
```
This is a 2D array which will hold int values _only_. Array elements are 0 indexed.

**Slices** are segment of an array. They are always associated with the underlying array. Shares saved storage with array. The length of slice is variable and found using the len function. E.g. s := []int {1,2,3,4,5}
Another way of declaring slices in using built-in function make.
E.g. make([]T , length, capacity) T is data type, length is the length of slice and capacity is the length of underlying array. If capacity is not given, it is assumed to be equal to length.
A slice/ array can be sliced using high low expression. E.g. a := s[2:4] 
2 is the lower index (inclusive) and 4 is the higher index(exclusive). Either one of the higher or lower index can be omitted. Then will be begin at 0 or go till end of the slice/array respectively. Capacity of the new slice is from lower index to the end of array. But length is 2 in this case consisting of s[2] and s[3].
Another example: s = s[:cap(s)] this will create a slice taking up complete capacity of former s.

range keyword can be used to traverse through the length of the slice/array
If the allocated underlying array is not enough, a new array can be created using append function:
slice2 := append(slice1, 11,10)  //slice1 is a slice defined earlier

copy is used to copy data from one slice to another. Data from source slice is copied starting from index 0 and goes upto the length of the destination slice. Copying is dependent is length of destination slice and independent of capacity.

nil : used to check if a slice is nil. Only this slice satisfies the condition. var slice []int

A structure is a collection of data types. It is user-defined.
E.g. type struct_name struct {
    var1 int
    var2 int
}
Values can be assigned using a construct,or struct_name.variable_name or pointer_to_struct_var.variable name.

Maps in go are collection of unordered elements, here keys and values. They cannot be referenced using indexes. Syntax is:
var m = map[key_type]value_type { //initializations }
Accessing a map:  map_name[“key_name”]
This trues two values, value stored in the key, and a boolean if it exists or not (true is key exists, false if not)
A key:value can be deleted in map:  delete(map_name, key)

If a function to take unknown number of arguments, args  keyword is used.
E.g. func function_name ( args ...data_type ){ //statements }

Closure is a function created inside another function. It is a local function.
E.g. function_name := func() int { //statements with return }
It can then access the local variables of that function. 

Go supports use of recursion in its usual form. Closure and recursion are form of functional programing.
