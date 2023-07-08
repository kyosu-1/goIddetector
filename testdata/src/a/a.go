package a

func f() {
	// The pattern can be written in regular expression.
	var Id int // want "Id should be an ID"
	print(Id) // want "Id should be an ID"
}

type A struct {
	Id int // want "Id should be an ID"
}
