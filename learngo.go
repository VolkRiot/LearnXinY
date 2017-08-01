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
  //learnTypes() // < y minutes, learn more!
}

/* <- multiline comment
Functions can have parameters and (multiple!) return values.
Here `x`, `y` are the arguments and `sum`, `prod` is the signature (what's returned).
Note that `x` and `sum` receive the type `int`.
*/

func learnMultiple(x, y int) (sum, prod int) {
  return x + y, x * y
}
