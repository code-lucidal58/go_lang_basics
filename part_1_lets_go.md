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
This is preferred/convention way of including packages.  
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
* Multiple Variable declaration: ```var a,b,c bool```
Variable Naming conventions: camelCase  
Assigning value to variable:
```go
var card = "Ace of Spades" //creates new variable and assigns
//mentioning the type in the above statement is not necessary
a := ”Aanisha” //another method: creates new variable and assigns
a = “Aanisha” // assigns existing variable a
```
If variable is not initialised after declaration, it is initialised to 0 if int, false
if bool and “” if string.  
***NOTE***: Variables can be declared outside all functions, but cannot be initialised.
Initialisations must be done inside a function.  
### Data Types in Go
uint8, uint16, uint32, uint64 (unsigned integers, memory located as per size mentioned in
the name), int8, int16...int64(signed integers). If just int is used, it becomes
platform dependent. For floating point number, float32, float64 : complex32, complex64
for storing complex numbers. Bool: true, false. String.
