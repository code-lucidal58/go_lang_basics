# HTTP client and server
HTTP client and server implementations are available in `net/http` package. This is a huge package.
The package documentation can be found [here](https://golang.org/pkg/net/http/)

Response struct has Status StatusCode, Body (`io.ReadCloser`), etc.
`io.ReadCloser` is an interface that has two interface inside it i.e. `Reader` and `Closer` interface.
These two interfaces have `Read` and `Close` names functions in it respectively.

In `Reader`, the `Read` function's signature takes byte slice as a paramter. This means user will provide the byte slice
the implementation of `Read` will modify the slice (remember slice is passed by reference). The function should return
an `int` i.e. the number of bytes read and `error` i.e. error if encountered any.

```go
resp, err := http.Get("http://google.com")
if err != nil {
    fmt.Println("Error:", err)
    os.Exit(1)
}
// Read will take place till the byte slice is full or object to read is complete
var bs = make([]byte, 99999)
resp.Body.Read(bs)
fmt.Println(string(bs))
```
`Writer` interface can be used to print a []byte data. Ih has the function `Write` which receives a byte slice and returns
an integer and error.

To implement both of these (Reader and Writer) simultaneously, `io.Copy` is used. It takes a `Writer` and `Reader`. Reads
data from somewhere and writes it to somewhere. 
```go
_, err := io.Copy(os.Stdout, resp.Body)
```

Here is an example to a custom struct that implements `Writer`.
```go
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

lw := logWriter{}
io.Copy(lw, resp.Body)
```