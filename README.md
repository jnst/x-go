# x-go

[![Go Report Card](https://goreportcard.com/badge/github.com/jnst/x-go)](https://goreportcard.com/report/github.com/jnst/x-go)

### Run

```
$ go run main.go factorial
```

### Test

```
$ go test .
```
or
```
$ go test -run TestFactorial
```


### Generate protobuf

```
$ protoc --go_out=. internal/*.proto
```
