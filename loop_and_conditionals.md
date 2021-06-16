# Loops and Conditionals
Loops are used to access elements of a iterable data structure. Conditional statements are used to make decisions. They 
are also used to skip a check of code dependent on input or other parameters. 

Refer `conditionals.go` in this repository for practical usages.  

## Conditional statements
 **If else** statements have the following syntax
```txt
if (condition1) {
    block 1
} else if (condition2) {
    block 2
} else {
    block 3
}
```

**NOTE**: Parentheses are required even if `if` block has just one statement. Same applies for other conditionals.

There are two types of switch statements:
* Type switch:  `switch <expression> {}`
* Expression switch: `switch {} //no expression`

## Loops
`for` statements iterate over iterable data structures. One of the ways of using the keyword: `for i:= 0; i<10;i++{}`.
To loop over slices, array, maps or other iterable objects, `for` is useful. It can be couple with `range`. 
```go
cards = []string{"Ace","Spade"}
for index, suit := range cards{
	fmt.Println(index, suit)
}
```
`:=` is used for variables `index` and `suit` because these variables are discarded and re-created after each iteration.
