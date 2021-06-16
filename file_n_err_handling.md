# File and Error Handling

## File Handling: IO utility

`io/ioutil` package implements some IO utility functions.

### Write data to a file

`WriteFile`  accepts three parameters: file name(string), data to be written(byte slice), and permissions (os.FileMode)
for the file. This functions return error.

```go
err := ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
```

### Read a file

`ReadFile` accepts the file name/path that is to be read. It returns two values: slice of byte and error.

```go
bs, err := ioutil.ReadFile(fileName)
```

### Delete a file

`Remove` function accepts the file name/path that is to be deleted.

## Error Handling

Error is an important part of the language. Many functions return error and allows the user to decide the flow of the
code in such situation. 