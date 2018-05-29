# GoLangBasics

### Introduction to the language
Go is a general purpose programming language. It is open source. Due to its concurrency feature it is ideal for 
multi-core and networked machines. Best suited for modular software designs. Compiles very quickly and performs 
runtime reflection. Developed by Google in 2007, announced in 2009 and stable version released in 2015. 
Go is a case sensitive language.

### Lets GO
Every Go program is made up of packages. Program starts running in _main_ package.<br>
Package name is the same as the last element of the import path. For instance, the _math/rand_ package comprises files that begin with the statement _package rand_. 

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

A name is exported from another package if it starts with capital letter.
```go
fmt.Println(math.Pi)
```
_Pi_ is a name exported by package _math_. 

### Functions
Functions in go can take two or more arguments.
