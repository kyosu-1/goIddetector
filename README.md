# goIddetector

goIddetector is a Go static analysis tool that detects when an identifier named Id is used in the source code.

## Installl

```
$ go install github.com/kyosu-1/goIddetector
```

## Usage

```
$ go vet -vettool=$(which goIddetector) path/to/your/package
```

## Example

Consider the following Go code:

```go
package main

import "fmt"

var Id = "12345"

func main() {
	Id := "67890"
	fmt.Println(Id)
}

func Id() {
	// ...
}
```

The `goIddetector` tool will detect the identifier `Id` in various places and report them as a warning. The output might look something like this:

```
main.go:5:5: Id should be an ID
main.go:8:2: Id should be an ID
main.go:12:6: Id should be an ID
```

These warnings indicate that the identifier 'Id' is being used as a variable name, a function name, and a local variable name. Following Go's convention, Id should be an identifier name called ID.

## ref

https://github.com/golang/go/wiki/CodeReviewComments#initialisms