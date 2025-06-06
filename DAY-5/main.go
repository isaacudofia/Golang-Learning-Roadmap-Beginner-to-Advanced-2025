package main

import (
	"errors" // Import the errors package for creating custom errors
	"fmt"
)

// Function 1: No parameters, no return values
func sayHello() {
    fmt.Println("Hello from sayHello function!")
}

// Function 2: Two integer parameters, one integer return value
func add(a, b int) int {
    return a + b
}

// Function 3: Two float64 parameters, two float64 return values (area and circumference)
// Uses named return values
func calculateCircleMetrics(radius float64) (area float64, circumference float64) {
    const pi = 3.14159
    area = pi * radius * radius
    circumference = 2 * pi * radius
    return // Naked return (returns current values of area and circumference)
}

// Function 4: Division with multiple return values (result and an error)
func divide(numerator, denominator float64) (float64, error) {
    if denominator == 0 {
        // Return zero value for float64 and a new error
        return 0, errors.New("cannot divide by zero")
    }
    // Return the result and nil (meaning no error)
    return numerator / denominator, nil
}

// Function 5: Demonstrating 'defer' for cleanup
func simulateFileProcessing(filename string) {
    fmt.Printf("Attempting to open file: %s\n", filename)

    // Simulate opening a file. In a real scenario, this would be os.Open(filename)
    // For demonstration, let's just print a message.
    // Assume file opens successfully.
    fmt.Println("File opened successfully.")

    // Defer this call to ensure it runs just before the function returns
    defer fmt.Println("File closed.")
    defer fmt.Println("Performing final cleanup...") // This will run before "File closed." (LIFO)

    // Simulate some work with the file
    fmt.Println("Reading data from file...")
    // ... some complex operations ...
    fmt.Println("Data processing complete.")

    // No explicit close call here; defer handles it.
    // What if an error occurs here? defer still runs!
    // fmt.Println(1 / 0) // Uncommenting this would cause a panic, but defer would still run!
}


func main() {
    fmt.Println("--- Function Calls ---")

    // Calling Function 1
    sayHello()

    // Calling Function 2
    sumResult := add(10, 5)
    fmt.Printf("10 + 5 = %d\n", sumResult)

    // Calling Function 3
    circleArea, circleCircumference := calculateCircleMetrics(7.0)
    fmt.Printf("Circle with radius 7.0: Area = %.2f, Circumference = %.2f\n", circleArea, circleCircumference)

    // Calling Function 4: Handling multiple return values, including error
    fmt.Println("\n--- Division Examples ---")
    result1, err1 := divide(10.0, 2.0)
    if err1 != nil {
        fmt.Println("Error:", err1)
    } else {
        fmt.Printf("10.0 / 2.0 = %.2f\n", result1)
    }

    result2, err2 := divide(10.0, 0.0) // This will cause an error
    if err2 != nil {
        fmt.Println("Error:", err2)
    } else {
        fmt.Printf("10.0 / 0.0 = %.2f\n", result2)
    }

    // Calling Function 5: Demonstrating defer
    fmt.Println("\n--- Defer Example (Simulated File Processing) ---")
    simulateFileProcessing("my_document.txt")
    fmt.Println("Program continues after simulateFileProcessing.")
}