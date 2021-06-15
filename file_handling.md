# File Handling: IO utility
`io/ioutil` package implements some IO utility functions.

## Write data to a file
`WriteFile`  accepts three parameters: file name(string), data to be written(byte slice), and permissions (os.FileMode)
for the file. This functions return error.
```go
err := ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
```

## Read a file
`ReadFile` accepts the file name/path that is to be read. It returns two values: slice of byte and error.
```go
bs, err := ioutil.ReadFile(fileName)
```

## Delete a file
`Remove` function accepts the file name/path that is to be deleted.