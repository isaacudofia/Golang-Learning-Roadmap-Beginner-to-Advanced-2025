// calculator/add.go
package calculator // This file belongs to the 'calculator' package

// Add exports the addition function. Its name starts with an uppercase letter.
func Add(a, b int) int {
    return a + b
}

// subtract is an unexported function, only visible within the calculator package.
func subtract(a, b int) int {
    return a - b
}