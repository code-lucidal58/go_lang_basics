# GoLangBasics

Golang Notes from multiple sources.

Pre-requisite to understand the content
- Knows any other programming language basics

Using Golang 1.16

## Introduction to the language
Go is a general purpose programming language. It is open source. It is NOT an object-oriented programming language. 
Due to its concurrency feature it is ideal for multi-core and networked machines. Best suited for modular
software designs. Compiles very quickly and performs runtime reflection. Developed
by Google in 2007, announced in 2009 and stable version released in 2015. Go is a
case sensitive language.  
Go is not an Object Oriented Programming Language.

## Installing Go
Go can be downloaded from https://golang.org/dl/ . The installer will append *PATH*
variable of the system with the directory path to go/bin. By default, the working directory
is taken as $HOME/go. It set it something different from the convention, an environment
variable *GOPATH* needs to be created. For detailed steps for installation instructions,
follow this link https://golang.org/doc/install .

## Go CLI commands
Building a go file, creates an executable file, with the same name as the go file,
in the same location. Suppose the script is ```main.go```. The executable can be
executed in mac by writing ```./main``` and in windows by ```main.exe```.
Some of the frequently used commands are:
* **go build**: compiles go sources files.
* **go run**: compiles and executes one or two files.
* **go fmt**: formats all files in current directory.
* **go install**: compiles and install a packages.
* **go get**: downloads the raw source code of third party packages.
* **go test**: runs any tests associated with current project.

## Contents of the repository
* [Introduction](./lets_go.md) contains introduction to packages, variables, data types, type aliasing, conversion and pointers.
* [Loops and Conditions](./loop_and_conditionals.md) contains notes on `if` and `for` statements.
* [Functions](./functions.md)  

[hello](./single_page_scripts/hello.go), [data_structure](./single_page_scripts/data_structures.go),
and [conditionals](./single_page_scripts/conditionals.go) contain scripts based on the concepts learned.

* [Test](tests.md) contains notes on go unit tests
* [File and Error handling](file_n_err_handling.md) contains a brief peek into the `file` and `os` packages.
* [HTTP calls](http_calls.md) is about make HTTP calls and read the data received.

[http_interface](./single_page_scripts/http_interface.go), [interface](./single_page_scripts/interfaces.go) 
contain scripts based on the concepts learned.

There are some more scripts in the `single_page_scripts` directory. The file names are kept intuitive.  
[Cards](./cards) is a simple project for dealing a set of cards.  
[Knowledge Check](./single_page_scripts/knowledge_check.go) contains some logics that can challenge the reader's understanding.