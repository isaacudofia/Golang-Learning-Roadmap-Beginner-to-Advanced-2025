package main

import "fmt"

func main() {
    fmt.Println("--- Go Arrays Demonstration ---")

    // 1. Declare an array of 5 integers (initialized to zero values)
    var numbers [5]int
    fmt.Printf("1. Default initialized array (numbers): %v (Length: %d)\n", numbers, len(numbers))

    // 2. Initialize elements individually
    numbers[0] = 10
    numbers[1] = 20
    numbers[2] = 30
    // numbers[3] and numbers[4] remain 0
    fmt.Printf("2. After individual assignment (numbers): %v\n", numbers)

    // 3. Declare and initialize an array using a composite literal
    // Compiler infers the length using "..."
    grades := [...]int{95, 88, 72, 91, 85}
    fmt.Printf("3. Grades array: %v (Length: %d)\n", grades, len(grades))

    // 4. Declare and initialize an array of strings
    weekdays := [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
    fmt.Printf("4. Weekdays array: %v (Length: %d)\n", weekdays, len(weekdays))

    // 5. Accessing and Modifying Elements
    fmt.Println("\n--- Accessing and Modifying Elements ---")
    fmt.Printf("First grade: %d\n", grades[0]) // Access first element
    fmt.Printf("Last weekday: %s\n", weekdays[len(weekdays)-1]) // Access last element

    weekdays[1] = "New Monday" // Modify an element
    fmt.Printf("Modified Monday: %s\n", weekdays[1])
    fmt.Printf("Weekdays array after modification: %v\n", weekdays)

    // Attempting to access an out-of-bounds index will cause a runtime panic
    // fmt.Println(numbers[5]) // Uncommenting this will cause: runtime error: index out of range [5] with length 5

    // 6. Iterating over an array using a traditional for loop
    fmt.Println("\n--- Iterating with Traditional For Loop ---")
    fmt.Println("Grades:")
    for i := 0; i < len(grades); i++ {
        fmt.Printf("  Grade at index %d: %d\n", i, grades[i])
    }

    // 7. Iterating over an array using a for-range loop (recommended)
    fmt.Println("\n--- Iterating with For-Range Loop ---")
    fmt.Println("Numbers:")
    totalSum := 0
    for index, value := range numbers {
        fmt.Printf("  Element at index %d: %d\n", index, value)
        totalSum += value
    }
    fmt.Printf("Total sum of 'numbers' array: %d\n", totalSum)

    // If you only need the value (and not the index)
    fmt.Println("\n--- Iterating (Value Only) with For-Range Loop ---")
    maxGrade := grades[0] // Assume first element is max
    for _, grade := range grades {
        if grade > maxGrade {
            maxGrade = grade
        }
    }
    fmt.Printf("Highest grade: %d\n", maxGrade)

    // 8. Arrays are Value Types
    fmt.Println("\n--- Arrays as Value Types ---")
    originalArray := [3]int{100, 200, 300}
    copiedArray := originalArray // A copy is made
    fmt.Printf("Original: %v, Copied: %v\n", originalArray, copiedArray)

    copiedArray[0] = 999 // Modify the copied array
    fmt.Printf("Original after copy modification: %v\n", originalArray) // Original remains unchanged
    fmt.Printf("Copied after copy modification: %v\n", copiedArray)
}