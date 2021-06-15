# Testing in Go

Create a new file with ending as `_test.go`. Run all tests in a package using `go test`.
`testing` package is used for testing. All test functions should start with capital letter. These functions should
have `t *testing.T` as its first parameter. Error formatiing, test logging etc. should be done using the testing
variable `T`.

The testing framework will not take care of the setup cleanup. This is to be handled by the developer. Hence, any breakage
in the test functions should be handled properly.
