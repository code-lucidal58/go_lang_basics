# Functions
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

### Methods and Interfaces
Refer `methods_interfaces.go` in this repository for practical usages<br>
Methods are functions declared with a receiver. Receiver is similar to parameter. A receiver is a value or a pointer
of a named or struct type. See example in code for more clarification. A point to be noted, parameters are
declared before method_name in method declaration.
