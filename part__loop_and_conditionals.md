### Conditionals
Refer `conditionals.go` in this repository for practical usages<br>
**For** statements are similar to _while_ statements in C++. Iterator is modified inside the _for_ block statements.
Don’t worry, the normal version is also supported.  `for i:= 0; i<10;i++{}`. **If else** logic is same as C++,
nested if, else if are also supported.<br>
_**NOTE**_: Parentheses are required even if _if_ block has just one statement. Same applies for other conditionals.

There are two types of switch statements:
* Type switch:  `switch <expression> {}`
* Expression switch: `switch {} //no expression`

## Loops
To iterate through iterable objects, e.g. slices, `for` statements can be used.
```go
cards = []string{"Ace","Spade"}
for index, suit := range cards{
	fmt.Println(index, suit)
}
```
The variable declared in the first line of for loop, should be used in the *for* body
