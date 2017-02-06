# go-world

[![Go Report Card](https://goreportcard.com/badge/github.com/jnst/go-world)](https://goreportcard.com/report/github.com/jnst/go-world)

### Run

```
$ go run factorial.go -n 2
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
