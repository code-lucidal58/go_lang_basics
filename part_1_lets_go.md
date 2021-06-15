# Lets Go

Every Go program is made up of packages. If program has **main** package, it is an executable program, and execution
start from this package. A **package** can be thought as a project or workspace. All related files are inside this
package, and their first line (always) mentions the package name. For instance, the _math/rand_
package comprises files that begin with the statement _package rand_. Package name is the same as the last element of
the import path.  
There are two types of packages:

* **executable**: spits out an executable file. It is used while creating runnable applications. It must
  have ```package main``` as first line and should have a main function.
* **reusable**: code dependencies and libraries packages. Package name can be anything except *main*.

## Import statements

These statements give code access of the imported package to the current package. List of inbuilt
packages: https://golang.org/pkg/

### Factored Import statements

```go
import (
"fmt"
"math"
)
```

This is preferred/convention way of including packages.

### Non factored Import statements

```go
import "fmt"
import "math"
```

Some key packages in Go: strings, io/ioutils, bytes, os, path/filepath, errors, container/list, hash, encoding/gob, fmt.

A name is exported from another package if it starts with capital letter. Only exported names are accessible to other
packages.

```go
fmt.Println(math.Pi)
```

_Pi_ is a name exported by package _math_.  
Files in the same package do not need to be imported.

## Variables

Variables in Go starts with letter, contains numbers and underscore. It makes use of `var` keyword.

* Single Variable declaration: `var a bool`
* Multiple Variable declaration: `var a,b,c bool`
  Variable Naming conventions: camelCase

Just underscore (_) can also be used as a replacement to variable when that value is to be ignored and will not be used
ahead in the code.

```go
var card = "Ace of Spades" //creates new variable and assigns
//mentioning the type in the above statement is not necessary
a := "Aanisha" //another method: creates new variable and assigns
a = "Aanisha"        // assigns existing variable a
_, card := newCard() // newCard returns two values. Ignore the first one.
```

If a variable is not initialised after declaration, it takes in falsey values: 0 if int, false if bool and "" if string.
Go complains if a declared/initialised variable is not used.

***NOTE***: Variables can be declared outside all functions, but cannot be initialised. Initialisations must be done
inside a function.

### Data Types in Go

uint8, uint16, uint32, uint64 (unsigned integers, memory located as per size mentioned in the name), int8,
int16...int64(signed integers). If just int is used, it becomes platform dependent. For floating point number, float32,
float64 : complex32, complex64 for storing complex numbers. Bool: true, false. String.

### Strings vs Characters

A single character uses single quotes for representation. This is also called a `rune`. In Go, `rune` is actually an
alias of int32. A string uses double quotes for representation. This is the reason all imported packages are mentioned
in double quotes.

String can also be represented as a byte slice, where each character in the string is stored as its ASCII value in a
slice. This is a computer-friendly represented of a string. String can be explicitly converted to a byte slice using
type conversion.

### Arrays vs Slices

Array is primitive data structure to hold records. It has a fixed length that is mentioned beforehand. Slices is a fancy
array. It can grow or shrink at will. New element can be added using `append` function. This function does not modify
the original slice. It creates a new slices and returns it. All elements should be of identical data types.

```go
cards := []string{"a", "b", "cd"} // slice
cards = append(cards, "df")       //add new element
fmt.Println(cards[2]) // cd
```

To create a slice out of a slice, `arr[start_index:end_index]` syntax can be used. Start index is included and end index
is not included. If start index if left empty, that means start from the beginning. Similarly, if end index is missing,
it means go till the end.

```go
a := []string{"a", "b", "c", "d"}
fmt.Println(a[1:3]) // b c
```

### Type aliasing

An alias of a data type can be created to attached more features to the datatype. In the below example, the type `cards`
is an alias for a string of array. `deck` will have all properties of a slice of strings. Custom methods can also be
attached to this new datatype.

```go
type deck []string
func (d *deck) findAce(){
//function block
}
```

The function `findAce` can be called on objects of `deck` but not on other variable that are of type `[]string`.

## Type conversion

The syntax for type conversion is `TypeToBeConverted(VariableToBeConverted)`. For example,
`string(23)` or `[]byte("Hello World")`.

## Random

I have added random functionality in the first notes-sheet because I think it is an important and easy part in any
language. And, it does not get much credit. `math/rand` is used to generate pseudo random numbers. There are number of
functions for this purpose inside the package. One of them is `Intn`. It takes an integer n as parameter, and returns a
pseudo-random integer between 0 and <n. It panics if n <= 0.

`rand.seed` is used to set the seed value. This will randomise the number generation on multiple runs. Another
interesting function is `rand.Shuffle`. It is used to shuffle elements in an object.

```go
rand.Seed(time.Now().UnixNano())
d := []int{1, 2, 3, 7, 9, 10, 34, 12}
rand.Shuffle(len(d), func (i, j int) {
    d[i], d[j] = d[j], d[i]
})
```