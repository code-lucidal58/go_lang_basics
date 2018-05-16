# GoLangBasics
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
This is preferred/ convention way of including packages.
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
Functions in go can take two or more arguements.
