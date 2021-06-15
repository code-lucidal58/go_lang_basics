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
