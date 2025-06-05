package main

import "fmt"

func main() {
    // --- 1. Variable Declarations and Initialization ---

    // Full declaration with initial value
    var studentName string = "Alice Johnson"
    var studentAge int = 22

    // Type inference (Go figures out the type)
    var studentGrade = 'A' // Go infers rune (int32) for single quotes
    var isEnrolled = true

    // Short variable declaration (most common inside functions)
    courseName := "Go Programming"
    courseCode := "GOLANG101"
    credits := 3.5

    fmt.Println("--- Student Information ---")
    fmt.Printf("Name: %s\n", studentName) // %s for string
    fmt.Printf("Age: %d\n", studentAge)     // %d for decimal integer
    fmt.Printf("Grade: %c\n", studentGrade) // %c for character (rune)
    fmt.Printf("Enrolled: %t\n", isEnrolled) // %t for boolean
    fmt.Printf("Course: %s (Code: %s), Credits: %.1f\n", courseName, courseCode, credits) // %.1f for float with 1 decimal place

    // --- 2. Zero Values ---
    fmt.Println("\n--- Zero Values ---")
    var defaultInt int
    var defaultFloat float32 //float32 only has about 6 significant digits. If you need more precision, use float64.
    var defaultBool bool
    var defaultString string

    fmt.Printf("Default int: %d\n", defaultInt)
    fmt.Printf("Default float: %.2f\n", defaultFloat)
    fmt.Printf("Default bool: %t\n", defaultBool)
    fmt.Printf("Default string: '%s'\n", defaultString) // Notice the empty string

    // --- 3. Constants ---
    fmt.Println("\n--- Constants ---")
    const Pi = 3.14159
    const CompanyName = "GoTech Solutions"
    const MaxStudents = 50

    fmt.Printf("Pi: %f\n", Pi)
    fmt.Printf("Company: %s\n", CompanyName)
    fmt.Printf("Maximum Students: %d\n", MaxStudents)

    // Constants can be used in expressions
    const CircleRadius = 10.0
    circleArea := Pi * CircleRadius * CircleRadius
    fmt.Printf("Area of a circle with radius %.2f: %.2f\n", CircleRadius, circleArea)

    // --- 4. Basic Arithmetic Operations ---
    fmt.Println("\n--- Arithmetic Operations ---")
    num1 := 15
    num2 := 4
    floatNum1 := 15.0
    floatNum2 := 4.0

    sum := num1 + num2
    difference := num1 - num2
    product := num1 * num2
    quotientInt := num1 / num2 // Integer division
    quotientFloat := floatNum1 / floatNum2 // Floating-point division
    remainder := num1 % num2

    fmt.Printf("%d + %d = %d\n", num1, num2, sum)
    fmt.Printf("%d - %d = %d\n", num1, num2, difference)
    fmt.Printf("%d * %d = %d\n", num1, num2, product)
    fmt.Printf("%d / %d (integer division) = %d\n", num1, num2, quotientInt)
    fmt.Printf("%.1f / %.1f (float division) = %.2f\n", floatNum1, floatNum2, quotientFloat)
    fmt.Printf("%d %% %d = %d\n", num1, num2, remainder)

    // Type conversion is explicit in Go
    // For example, converting int to float64 for calculation
    resultMixed := float64(num1) / floatNum2
    fmt.Printf("int(%d) / float(%.1f) = %.2f\n", num1, floatNum2, resultMixed)

    // Using iota for enums
    const (
        // iota is reset to 0 for each const block
        StatusPending = iota // 0
        StatusApproved       // 1 (iota increments automatically)
        StatusRejected       // 2
    )
    fmt.Println("\n--- iota Example ---")
    fmt.Printf("StatusPending: %d\n", StatusPending)
    fmt.Printf("StatusApproved: %d\n", StatusApproved)
    fmt.Printf("StatusRejected: %d\n", StatusRejected)
}