Go is not an Object Oriented Programming Language. Hence types are created to replicate
the functionalities obtained from creating classes.

# Object Oriented Programming in Go
`type` keyword is used to create an alias of slice.
```go
type deck []string
```
In the application, any string slice can now be replaced with the word 'deck'.
Functionalities will still remain the same. Methods can be attached to types,
called ***receiver function***. Here is an example:
```go
func (d deck) print(){
  for i,cards := range d {
    fmt.Println(i, cards)
  }
}
```
`d deck` is the receiver variable of type `deck`. It is a reference to the
actual variable that is calling this function. It is a convention of assigning
single letter name to receiver variable, mainly the initials of the type name.
This receiver variable can be thought as `self` or `this` in other OOP languages.

The usage of type and receiver functions can be found in the project
[cards](https://github.com/code-lucidal58/go_lang_basics/tree/master/cards)
