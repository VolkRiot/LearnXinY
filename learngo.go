// Single line comment
/* Multi line comment
... see */

// Package clause starts soure file
package main

import (
  "fmt" // Go standard lib package
  "io/ioutil" // Implements some IO functionality
  m "math" // Math library with m as local alias
  "net/http" // Web Server!!!!
  "os" // OS functions like working with file system
  "strconv" // String conversions library
)

// A function definition. Main is special. It is the entry point for the
// executable program. Love it or hate it, Go uses brace brackets.

func main() {
  // Println outputs a line to stdout.
  // Qualify it with the package name, fmt.
  fmt.Println("Hello World!")

  // Call another function in main
  beyondHello()
}

// Functions have parameters in parentheses.
// If there are no parameters, empty parentheses are still required.
func beyondHello() {
  var x int // Var declaration. Var must be declared before it is used
  x = 3  // Vr assingment

  // "Short" declarations use := to infer the type, declare, and assign.
  y := 4

  sum, prod := learnMultiple(x, y)
  fmt.Println("sum:", sum, "prod:", prod) // Simple output, auto deliniated
  learnTypes() // < y minutes, learn more!
}

/* <- multiline comment
Functions can have parameters and (multiple!) return values.
Here `x`, `y` are the arguments and `sum`, `prod` is the signature (what's returned).
Note that `x` and `sum` receive the type `int`.
*/

func learnMultiple(x, y int) (sum, prod int) {
  return x + y, x * y
}

func learnTypes() {
  str := "Learn go!"

  s2 := `A "raw" string literal can include line breaks`
  g := 'Σ' // rune type, an alias for int32, holds a unicode code point.

  f := 3.14195 // Float type
  c := 3 + 4i // complex128, represented internally with two float64's.

  // Var syntax with initializers
  var u uint = 7 // Unsigned int
  var pi float32 = 22. / 7

  // Conversion syntax with a short declaration
  n := byte('\n')

  // Arrays have size fixed at compile time.
  var a4 = [4]int // an array of 4 integers intialized to all 0
  a5 := [...]int{3, 1, 5, 10, 100} // An array initialized with a fixed size of five
    // elements, with values 3, 1, 5, 10, and 100.

  // Slices have dynamic size. Arrays and slices each have advantages
  // but use cases for slices are much more common.

  s3 := []int{4, 5, 9} // Compare to a5. No ellipsis here.
  s4 := make([]int, 4) // Allocates slice of 4 ints, initialized to all 0.
  var d2 [][]float64 // Declaration only, nothing allocated here.
  bs := []byte("a slice") // Type conversion syntax


}
