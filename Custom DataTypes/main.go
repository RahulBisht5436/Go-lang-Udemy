package main

// newstr is a custom (named) type whose underlying type is string.
// Defining a new type — instead of a type alias (`type newstr = string`) —
// gives us a distinct type identity, which lets us attach our own methods to it
// without modifying the built-in `string` type.
type newstr string

// AppendPrint is a method with a value receiver `(n newstr)`.
// Methods can only be defined on types declared in the same package,
// which is exactly why we created `newstr` above instead of using `string` directly.
// A value receiver works on a copy of `n`, so any mutation here would not
// affect the caller's variable (use a pointer receiver `*newstr` if mutation is needed).
func (n newstr) AppendPrint(message string) {
	println(message)
	println(n)
}

func main() {
	// `data` is explicitly typed as `newstr` so that the method set of `newstr`
	// (including AppendPrint) is available on it. A plain `string` literal
	// would not have access to AppendPrint.
	var data newstr = "This is the original data"

	// Method call uses dot syntax; Go automatically passes `data` as the receiver `n`.
	data.AppendPrint("This is appended data")
}
