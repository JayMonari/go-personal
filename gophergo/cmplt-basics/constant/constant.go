package constant

import (
	"fmt"
	"math"
)

// Stuck is an untyped string
const Stuck = "This variable can never be reassigned."

// Stuck = "This won't work, Stuck is constant!"
// NOTE(jay): cannot assign to Stuck (untyped string constant "This variable
// can never be reassigned.")

// HeartEyes is an untyped rune
const HeartEyes = '😍'

// Arithmetic is an untyped float
const Arithmetic = 600 / 3.421

// AlwaysTrue is an untyped bool
const AlwaysTrue = true

// MaxByteValue is a constant byte value. To let the compiler know we wanted a
// byte instead of an int, we can specifically tell it we want that value.
const MaxByteValue byte = 255

// Create a grouping of const values, not needed to type the `const` keyword
// over and over again. This also works for `var`!
const (
	IsConst                   = true
	IsInGrouping              = true
	OneAndQuarterAsUntypedInt = 5 / 4
)

const piThousand = 3.141592653589793238462643383279502884197169399375105820974944592307816406286208998628034825342117067982148086513282306647093844609550582231725359408128481117450284102701938521105559644622948954930381964428810975665933446128475648233786783165271201909145648566923460348610454326648213393607260249141273724587006606315588174881520920962829254091715364367892590360011330530548820466521384146951941511609433057270365759591953092186117381932611793105118548074462379962749567351885752724891227938183011949129833673362440656643086021394946395224737190702179860943702770539217176293176752384674818467669405132000568127145263560827785771342757789609173637178721468440901224953430146549585371050792279689258923542019956112129021960864034418159813629774771309960518707211349999998372978049951059731732816096318595024459455346908302642522308253344685035261931188171010003137838752886587533208381420617177669147303598253490428755468731159562863882353787593751957781857780532171226806613001927876611195909216420198

// However you **cannot** declare arrays, slices, maps, or structs constant.
// NOTE(jay): (value of type [2]string) is not constant
// const myArray = [2]string{"won't", "work"}
// const mySlice = []string{"still", "doesn't", "work"}
// const myMap = map[string]int{}
// const me = struct{ name string }{name: "Jay"}

// UntypedConst shows that constants can have values that will be automatically
// converted to the necessary type that the function needs at runtime.
func UntypedConst() {
	// const values do not have a type and therefore are very useful when you
	// don't want to have to do explicit casting.
	const untyped = 42
	// We don't care what this function does, we only care what it looks like
	// math.IsInf(float64, int)
	fmt.Println(math.IsInf(untyped, untyped))
	// If we try this with typed int we have to cast it.
	var typed int = 42 // or typed := 42
	fmt.Println(math.IsInf(float64(typed), typed))
}

// BigPreciseConstants shows that the compiler won't complain if you make a
// value that goes for 1000s of digits if it's a constant. Just when it comes
// time to use that constant it must be casted to a type, such as float64.
func BigPreciseConstants() { fmt.Printf("%.64f", float64(piThousand)) }

// UntypedString shows us that it is flexible to be **any** type that has
// `string` as the underlying type.
const UntypedString = "I fit wherever the underlying type of something is a string!"

// myString has an underlying type of `string` **but** it's `type` is myString
type myString string

// Print takes in a myString type and prints it to standard out.
func Print(s myString) { fmt.Println(s) }
