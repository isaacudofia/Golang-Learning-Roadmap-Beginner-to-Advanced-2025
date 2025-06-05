---
# Go (Golang) Learning Roadmap: Day 2 - Variables, Constants, and Basic Data Types

Welcome to Day 2 of your Go learning journey! Yesterday, you successfully set up your environment and ran your first "Hello, Go!" program. Today, we'll build on that foundation by diving into how Go handles data using variables, constants, and fundamental data types.
---

## Goal for Today:

- Understand how to declare and initialize variables in Go.
- Learn about Go's basic data types (integers, floats, booleans, strings).
- Explore type inference and zero values.
- Understand the concept and usage of constants.
- Perform basic arithmetic operations.

---

## 1. Concept: Variables

Variables are named storage locations that hold data. Go is a statically typed language, meaning a variable's type is known at compile time. However, Go's powerful type inference often allows you to write less verbose code.

### Declaring Variables:

Go provides several ways to declare variables:

1.  **Full Declaration:**

    ```go
    var name string = "Alice"
    var age int = 30
    var pi float64 = 3.14159
    var isGoFun bool = true
    ```

    This explicitly states the variable name, its type, and its initial value.

2.  **Type Inference (Initial Value Required):** Go can often infer the type from the initial value.

    ```go
    var name = "Alice" // Go infers 'string'
    var age = 30       // Go infers 'int'
    var pi = 3.14159   // Go infers 'float64'
    var isGoFun = true // Go infers 'bool'
    ```

    This is commonly used when you provide an initial value.

3.  **Short Variable Declaration (`:=`):** This is the **most common** way to declare and initialize variables in Go, especially within functions. It combines declaration and assignment, and Go infers the type.
    ```go
    name := "Alice"
    age := 30
    pi := 3.14159
    isGoFun := true
    ```
    **Important:** The `:=` operator can only be used _inside_ functions. For package-level variable declarations, `var` is required.

### Declaring without Initializing (Zero Values):

If you declare a variable without an initial value, Go automatically assigns a "zero value" to it:

- `0` for numeric types (`int`, `float`, etc.)
- `false` for booleans
- `""` (empty string) for strings
- `nil` for pointers, functions, interfaces, slices, channels, and maps

```go
var defaultInt int      // defaultInt will be 0
var defaultBool bool    // defaultBool will be false
var defaultString string // defaultString will be ""
```

---

## 2. Concept: Basic Data Types

Go has several built-in basic data types:

- **Numeric Types:**

  - **Integers:** `int` (platform-dependent), `int8`, `int16`, `int32`, `int64` (signed); `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr` (unsigned). `byte` is an alias for `uint8`, and `rune` is an alias for `int32`.
  - **Floating-Point Numbers:** `float32` (single-precision), `float64` (double-precision, more common).
  - **Complex Numbers:** `complex64`, `complex128`.

- **Boolean Type:**

  - `bool`: Represents `true` or `false`.

- **String Type:**
  - `string`: A sequence of bytes, typically representing UTF-8 encoded Unicode text. Strings in Go are **immutable**.

---

## 3. Concept: Constants

Constants are like variables, but their values are fixed at compile time and cannot be changed during program execution. They are declared using the `const` keyword.

### Declaration:

```go
const Pi = 3.14159
const Greeting = "Hello from a constant!"
const MaxAttempts = 3
```

Constants can also be declared without an explicit type, allowing Go to infer it.

### `iota`:

`iota` is a special identifier that represents successive untyped integer constants. It's often used to define enumerated constants, automatically incrementing its value within a `const` block.

```go
const (
    _ = iota // 0 (blank identifier to discard the first value)
    KB = 1 << (10 * iota) // 1 << 10 = 1024
    MB = 1 << (10 * iota) // 1 << 20 = 1048576
    GB = 1 << (10 * iota) // 1 << 30
)
```

---

## 4. Project for Today: Variable and Constant Playground

We'll write a program that declares variables and constants of various types, assigns values, performs some basic arithmetic, and prints them to the console.

### Instructions:

1.  **Create a new directory** for your Day 2 project:

    ```bash
    cd ~/go/src # or wherever you keep your Go projects
    mkdir day2_data_types
    cd day2_data_types
    ```

2.  **Create a new file** named `main.go` inside the `day2_data_types` directory.

3.  **Open `main.go`** in your code editor and type the following code:

    ```go
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
        fmt.Printf("Name: %s\n", studentName)     // %s for string
        fmt.Printf("Age: %d\n", studentAge)         // %d for decimal integer
        fmt.Printf("Grade: %c\n", studentGrade)     // %c for character (rune)
        fmt.Printf("Enrolled: %t\n", isEnrolled)   // %t for boolean
        fmt.Printf("Course: %s (Code: %s), Credits: %.1f\n", courseName, courseCode, credits) // %.1f for float with 1 decimal place

        // --- 2. Zero Values ---
        fmt.Println("\n--- Zero Values ---")
        var defaultInt int
        var defaultFloat float64
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
    ```

4.  **Save** the `main.go` file.

5.  **Run the program:** Open your terminal or command prompt, navigate to the `day2_data_types` directory, and run:
    ```bash
    go run main.go
    ```
    Observe the output and how different data types are handled and printed. Pay close attention to integer division vs. float division.

### Self-Reflection & Experimentation:

- **Try changing variable values:** What happens if you try to assign a string to an `int` variable? (Go will give a compile-time error).
- **Experiment with different data types:**
  - Declare an `int8` and see its range (try assigning a value larger than 127).
  - Declare a `float32` instead of `float64`.
- **Remove an `import "fmt"` statement:** What happens if you remove an import that is being used? (Go will give a compile-time error: `imported and not used: "fmt"`).
- **Remove a variable declaration that is not used:** What happens if you declare a variable but never use it? (Go will give a compile-time error: `variable declared and not used`). This is a helpful Go feature that prevents dead code.
- **Play with `Printf` format verbs:** The `fmt` package has many "verbs" for formatting different types (`%v` for default value, `%T` for type, etc.). Look up `fmt.Printf` documentation if you're curious!

You've made great progress today by understanding how Go handles data! This is fundamental to all programming.

## Get ready for Day 3, where we'll explore control flow with If/Else and Switch statements!
