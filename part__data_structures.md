# Data Structures
Refer
[data_structures.go](https://github.com/code-lucidal58/go_lang_basics/blob/master/single_page_scripts/data_structures.go)
practical usages.
## Arrays
Length of array must be specified. It can be a variable. An array type once declared cannot be changed.
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

### Pointers
Pointers have same logic as in C++. Parameters passed to a function are by default passed by value.
`&var_name` will give the address of the variable var_name. To dereference a pointer, * is used. Suppose var_pointer
is a pointer. `*var_pointer` will give the value in the address stored by pointer var_pointer.
To declare a pointer, new  keyword is used.
```go
y : = new(int)
```
This will allocate enough memory to store an integer value and assign the memory address to y.




**Defer functions**: Defer statement defers the execution of function. Function calls are stacked up and called in the
sequence of last in first out. Just state defer before making any function call. That function will be executed
after all statements are executed in main.
