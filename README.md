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
This is preferred/ convention way of including packages.<br>
**Non factored Import statements**: 
```go
import "fmt"
import "math"
```
