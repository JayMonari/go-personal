package functions

import "fmt"

// privateFunc is an example function, that is not exported. It is not always
// necessary to document unexported or private functions, because they won't
// show up in any documentation. It is only a good idea when the function does
// something that you wouldn't expect and should explain it.
func privateFunc() {
	fmt.Println("This function can only be called from within this package.")
}

// FuncPublic is an example function, that is exported. It is always a good
// idea to document your exported functions and variables, so that other
// developers can know how to use your code! Use `go doc --all .` in your
// terminal in this package and see what you get!
func FuncPublic() {
	fmt.Println("This function is exported and can be called anywhere.")
}

// FuncWithParams is an example function, that shows you how to pass in
// multiple arguments to a function and use them.
func FuncWithParams(name string, value int, emoji rune) {
	fmt.Printf("%s looks like %s and is a %d/10", name, string(emoji), value)
}

// FuncWithReturn is an example function on how to specify what type you want a
// function to return.
func FuncWithReturn() string {
	return "It's just this easy to return a type"
}

// FuncWithMultipleReturn is an example function that will return two types at
// the same time.
func FuncWithMultipleReturn() ([]int, bool) {
	return []int{1, 2, 3, 4, 5}, true
}

// FuncWithNamedReturns is an example function that shows how you can name all
// of your parameters and all of your return types if you want to. You will
// notice we don't have to specify the type over and over if they are the same
// type. i.e. (email string, url string) == (email, url string)
func FuncWithNamedReturns(name, scheme, host, path, query string) (email, url string) {
	// Notice we don't use `:=` for email and url. The function already makes
	// them for us when we named them up above.
	email = name + "@" + host
	url = scheme + host + path + query
	return email, url
}

// FuncVariadic is an example function. It takes in an arbitrary amount of
// `int`s and allows you to use all of them, the way you see fit. This can be
// seen as a more powerful version of `[]int`, and it works for all types.
func FuncVariadic(varargsNums ...int) (sum int) {
	for _, n := range varargsNums {
		sum += n
	}
	return sum
}
