// main.go
package main // This is the main executable package

import (
	"fmt"
	// Import our custom calculator package.
	// The path is the module path followed by the package directory.
	"calculator/calculator"
)

func main() {
    fmt.Println("--- Simple Calculator App ---")

    num1 := 10
    num2 := 5

    // Use the Add function from our calculator package
    sum := calculator.Add(num1, num2)
    fmt.Printf("%d + %d = %d\n", num1, num2, sum)

    // Use the Multiply function from our calculator package
    product := calculator.Multiply(num1, num2)
    fmt.Printf("%d * %d = %d\n", num1, num2, product)

    // Try to access the unexported function (this will cause a compile error!)
    // diff := calculator.subtract(num1, num2) // Uncommenting this line will cause an error!
    // fmt.Printf("%d - %d = %d\n", num1, num2, diff)

    fmt.Println("\nTo see an error, uncomment the line trying to call `calculator.subtract` in main.go.")
    fmt.Println("It won't compile because `subtract` is not exported (starts with lowercase).")
}