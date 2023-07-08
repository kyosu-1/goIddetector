# goIdDetector

goIddetector is a Go static analysis tool that detects when an identifier named Id is used in the source code.

## Detection Example

```
func f() {
	// The pattern can be written in regular expression.
	var Id int // want "Id should be an ID"
	print(Id) // want "Id should be an ID"
}
```

```
type A struct {
	Id int // want "Id should be an ID"
}
```

## ref

https://github.com/golang/go/wiki/CodeReviewComments#initialisms