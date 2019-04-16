# Lets Go
Every Go program is made up of packages. If program starts with _main_ package,
it is an executable file. Program starts executing from this package.
A package can be thought as a project or workspace. All related files are inside this
package and their first line mentions the package name. For instance, the _math/rand_
package comprises files that begin with the statement _package rand_.
Package name is the same as the last element of the import path.  
There are two types of packages:
* **executable**: spits out an executable file. It used while creating runnable
applications. It must have first line as ```package main``` and a main function.
* **reusable**: code dependencies and libraries packages. Package name can be any
thing except *main*. e.g. helper functions.

## Import statements
They give code access of the imported package to the current package.

### Factored Import statements
```go
import (
	"fmt"
	"math"
)
```
This is preferred/ convention way of including packages.  
### Non factored Import statements
```go
import "fmt"
import "math"
```
Some key packages in Go: strings, io, bytes, os, path/filepath, errors, container/list,
hash, encoding/gob. All package and their details can be found here: https://golang.org/pkg/ .
A name is exported from another package if it starts with capital letter.
Only exported names are accessible to other packages.
```go
fmt.Println(math.Pi)
```
_Pi_ is a name exported by package _math_.  
Files in the same package do not need to be imported.

## Variables
Variables in Go have the same convention as most other languages (starts with
letter, contains numbers and underscore).
* Single Variable declaration: ```var a bool```
* Double Variable declaration: ```var a,b,c bool```
Variable Naming conventions: camelCase  
Assigning value to variable:
```go
var card = "Ace of Spades" //creates new variable and assigns
a := ”Aanisha” //another method: creates new variable and assigns
a = “Aanisha” // assigns existing variable a
```
If variable is not initialised after declaration, it is initialised to 0 if int, false if bool and “” if string<br>
***NOTE***: Variables can be declared outside all functions, but cannot be initialised.
Initialisations must be done inside a function.  
**Types in Go**: uint8, uint16, uint32, uint64 (unsigned integers, memory located as per size mentioned in the name),
int8, int16...int64(signed integers). If just int is used, it becomes platform dependent. For floating point number,
float32, float64 : complex32, complex64 for storing complex numbers. Bool: true, false. String.

## Functions
Functions are defined using keyword ```func```. The parameters must have the type
specified after parameter name. Functions can return value using _return_ keyword.
The return type must also be specified in the function definition. Otherwise,
it is taken as void.
```go
E.g. func function_name(x int, y int) int {
	return x+y
}
```
The int  before curly braces states that this function is going to return an int value.
A function can return any number of results. Go's return values may be named. If so,
they are treated as variables defined at the top of the function. A return statement
without arguments returns the named return values. This is known as a _naked_ return.
They can harm readability in longer functions.  
When two or more consecutive named function parameters share a type, you can omit the
type from all but the last.
 * `(x int, y int)` can be shortened to `x, y int`.  
 * `(x int, y int) int` This part in the above function is called **signature**.
String length is found using *len* function: `len(string_variable)`
_%T_ is a format specifier which returns the type of a variable.
```go
fmt.Printf(“Type of the variable is%T”,var_name)
```

## Loops
To iterate through iterable objects, e.g. slices, `for` statements can be used.
```go
cards = []string{"Ace","Spade"}
for index, suit := range cards{
	fmt.Println(index, suit)
}
```
The variable declared in the first line of for loop, should be used in the *for* body

### Pointers
Pointers have same logic as in C++. Parameters passed to a function are by default passed by value.
`&var_name` will give the address of the variable var_name. To dereference a pointer, * is used. Suppose var_pointer
is a pointer. `*var_pointer` will give the value in the address stored by pointer var_pointer.
To declare a pointer, new  keyword is used.
```go
y : = new(int)
```
This will allocate enough memory to store an integer value and assign the memory address to y.


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
E.g. `make([]T , length, capacity)` T is data type, length is the length of slice and capacity is the length of underlying array. If capacity is not given, it is assumed to be equal to length.
A slice/ array can be sliced using high low expression. E.g. `a := s[2:4] `
2 is the lower index (inclusive) and 4 is the higher index(exclusive). Either one of the higher or lower index can be omitted. Then will be begin at 0 or go till end of the slice/array respectively. Capacity of the new slice is from lower index to the end of array. But length is 2 in this case consisting of s[2] and s[3].
Another example: `s = s[:cap(s)]` this will create a slice taking up complete capacity of former s.

_range_ keyword can be used to traverse through the length of the slice/array
If the allocated underlying array is not enough, a new array can be created using append function:
```go
slice2 := append(slice1, 11,10)  //slice1 is a slice defined earlier
```
_copy_ is used to copy data from one slice to another. Data from source slice is copied starting from index 0 and goes upto the length of the destination slice. Copying is dependent is length of destination slice and independent of capacity.

_nil_ : used to check if a slice is nil. Only this slice satisfies the condition. var slice []int

A _structure_ is a collection of data types. It is user-defined.
E.g. type struct_name struct {
    var1 int
    var2 int
}
Values can be assigned using a construct,or struct_name.variable_name or pointer_to_struct_var.variable name.

_Maps_ in go are collection of unordered elements, here keys and values. They cannot be referenced using indexes. Syntax is:
` m = map[key_type]value_type { //initializations }`
Accessing a map:  `map_name[“key_name”]`
This trues two values, value stored in the key, and a boolean if it exists or not (true is key exists, false if not)
A key:value can be deleted in map:  delete(map_name, key)

If a function to take unknown number of arguments, args  keyword is used.
```go
func function_name ( args ...data_type ){
	//statements
	}
```

Closure is a function created inside another function. It is a local function.
```go
function_name := func() int {
	//statements with return
	}
```
It can then access the local variables of that function.

Go supports use of recursion in its usual form. Closure and recursion are form of functional programing.

### Methods and Interfaces
Refer `methods_interfaces.go` in this repository for practical usages<br>
Methods are functions declared with a receiver. Receiver is similar to parameter. A receiver is a value or a pointer
of a named or struct type. See example in code for more clarification. A point to be noted, parameters are
declared before method_name in method declaration.
