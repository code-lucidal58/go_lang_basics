# Functions

`func` keyword defines a function. The parameters must have the type specified after parameter name. Functions can
return value using `return` keyword. The return type must also be specified in the function definition. Otherwise, it is
taken as void.

```
func function_name(x int, y int) int {
	return x+y
}
```
Arguments in the function call are by default passed by value. 
There is an exception to few, like slice, where it is passed by reference (not exactly).
Since slice is actually a structure with (pointerToHead, capacity and length), passing a slice as an argument, will
pass a copy of this structure into the function. PointerToHead is still the same address. Hence, no copy of original elements
is made. Other data structure that behave this way are maps, channels, pointers, functions. Pointers are not required
for these types.

Curly braces states that this function is going to return an `int` value. A function can return any
number of results. Go's return values may be named. If so, they are treated as variables defined at the top of the
function. A return statement without arguments returns the named return values. This is known as a _naked_ return. They
can harm readability in longer functions.  
When two or more consecutive named function parameters share a type, you can omit the type from all but the last.

* `(x int, y int)` can be shortened to `x, y int`.
* `(x int, y int) int` This part in the above function is called **signature**. String length is found using *len*
  function: `len(string_variable)`
  _%T_ is a format specifier which returns the type of a variable.

```
fmt.Printf(“Type of the variable is%T”,var_name)
```

A function can also have a receiver. A receiver adds a rule that only an object of the receiver can call the function.
In the below example, a user-defined type `deck` is the receiver. The function `print` can be called only by an object
of `deck`.

```go
func (d deck) print(){
// function body
}
```

By convention, the variable in the receiver is a single letter, usually the first letter of the receiver type. The receiver
object is sent by value i.e. a copy of the object calling the function is created and that is sent to the function. 

## Structs

Structs is a data structure that is a collection of data related to each other. `struct` keyword is used to declare a
structure. Properties in the struct can be accessed using the dot notation.

```go
type person struct {
  firstName string
  lastName string
}

// Go uses the order of property declaration in the struct definition
p1 := person{"Alex", "Anderson"}

// mention both key and value
p2 := person{firstName: "Alex", lastName: "Anderson"}

// more elaborated assigning of values
var p3 person
p3.firstName = "Alex"
p3.lastName = "Anderson"
// This is the same way properties can be updated
```

One important thing
to note while writing multi-line struct object, each line should be separated by a comma(,), even the last one.
```go
p4 := person {
	firstName: "Alex",
	lastName: "Anderson",
}
```

The type of a property inside a struct can be any primitive, non-primitive or user-defined datatype. There are two ways a 
struct can be embedded inside another struct.
```go
type child struct {
	p1 string
	p2 int
}
```
The struct can be added to the parent struct just like any other property i.e., with a field name and the data type.
```go
type parent struct {
	c1 child
	c2 string
	c3 boolean
}
```
Or, the property name can be omitted. In this case, it is assummed that the property name is same as its datatype.
```go
type parent struct {
  child // equivalent to `child child`
  c2 string
  c3 boolean
}
parent_object := parent{c2: "aa", c3: true, child : child{p1: "bb", p2: 10}}
```

Receivers can be added to functions to add methods to struct.
```go
func (p *person) updateLastName(name string) {
	p.lastName = name
}
```

Make sure to use pointer in the receiver section if any property value will be updated in the function body. The above
function can be called as `p3.updateLastName("Miller")`. Notice that a pointer of `p3` was not required to call the function.
Go is smart enough to figure that out. If receiver has a pointer, Go will pass the object by reference, even if not explicitly 
mentioned. Explicit mention of the pointer by adding `&` infront of `p3` is also acceptable.

### Methods and Interfaces

Refer `methods_interfaces.go` in this repository for practical usages<br>
Methods are functions declared with a receiver. Receiver is similar to parameter. A receiver is a value or a pointer of
a named or struct type. See example in code for more clarification. A point to be noted, parameters are declared before
method_name in method declaration.
