---
# Go (Golang) Learning Roadmap: Day 5 - Functions

Great! Welcome to Day 5 of your Go learning journey. So far, you've covered the absolute essentials: setting up, handling data with variables and constants, making decisions with `if`/`switch`, and repeating actions with `for` loops. Today, we're going to dive into **functions**, the building blocks that allow you to organize your code into reusable and manageable chunks.
---

## Goal for Today:

- Understand what functions are and why they are important.
- Learn how to define functions with parameters and return values (single and multiple).
- Explore named return values.
- Understand the concept of `defer` statements for cleanup.
- Practice writing and calling your own functions.

---

## 1. Concept: What are Functions?

A function is a block of organized, reusable code that is used to perform a single, related action. Functions provide:

- **Modularity:** Breaking down complex problems into smaller, manageable pieces.
- **Reusability:** Writing code once and using it multiple times.
- **Readability:** Making your code easier to understand and maintain.
- **Abstraction:** Hiding complex implementations behind a simple interface.

You've already used a function: `main()` is the special function where your program starts executing, and `fmt.Println()` is a function from the `fmt` package.

### Syntax for Defining a Function:

```go
func functionName(parameter1 type1, parameter2 type2) (returnValue1 typeA, returnValue2 typeB) {
    // Function body
    // ...
    return returnValue1, returnValue2 // Optional return statement
}
```

- **`func`**: The keyword to declare a function.
- **`functionName`**: The name of your function. Follow Go's naming conventions (start with a lowercase letter for unexported/private functions, uppercase for exported/public functions).
- **(parameter1 type1, parameter2 type2)**: The list of input parameters (arguments) the function accepts, along with their types. Parameters are optional. If multiple parameters share the same type, you can group them: `(a, b int, c string)`.
- **(returnValue1 typeA, returnValue2 typeB)**: The list of return values the function produces, along with their types. Return values are optional. If there's only one return value, the parentheses are often omitted (e.g., `(int)` becomes `int`).
- **`{}`**: The function body, containing the code to be executed.

---

## 2. Concept: Parameters and Return Values

- **Parameters:** Data passed into the function when it's called. They act as local variables within the function's scope.
- **Return Values:** Data sent back from the function after its execution. Go functions can return **multiple values**, which is a distinctive and powerful feature, often used for returning a result and an error (e.g., `value, err := someFunction()`).

### Examples:

```go
// No parameters, no return values
func sayHello() {
    fmt.Println("Hello there!")
}

// One parameter, one return value
func add(a int, b int) int { // can also be func add(a, b int) int
    return a + b
}

// Multiple parameters, multiple return values
func swap(x, y string) (string, string) {
    return y, x
}
```

---

## 3. Concept: Named Return Values

Go allows you to name your return values. This makes the code more readable, especially for functions with multiple return values, and you can use a "naked" `return` statement at the end, which automatically returns the current values of the named return variables.

```go
func calculateMetrics(radius float64) (area float64, circumference float64) {
    const pi = 3.14159
    area = pi * radius * radius
    circumference = 2 * pi * radius
    // A "naked" return returns the current values of area and circumference
    return
}
```

While convenient, excessive use of naked returns can sometimes reduce readability for very long functions.

---

## 4. Concept: `defer` Statement

The `defer` statement schedules a function call to be executed just before the surrounding function returns. This is incredibly useful for cleanup operations like closing files, unlocking mutexes, or flushing buffers, ensuring they happen regardless of how the function exits (e.g., normal return, panic).

Multiple `defer` statements are executed in **Last-In, First-Out (LIFO)** order.

**Example:**

```go
func processFile(filename string) {
    file, err := os.Open(filename) // os package would be imported for real file ops
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    // This defer will ensure file.Close() is called when processFile returns
    defer file.Close()
    fmt.Println("File opened successfully:", filename)

    // Simulate some work with the file
    // ...
    fmt.Println("Finished processing file.")
    // file.Close() will be called here automatically upon return
}
```

---

## 5. Project for Today: Simple Calculator & File Operation Simulation

We'll build a program that uses functions to:

- Perform basic arithmetic operations (addition, subtraction, multiplication, division).
- Demonstrate multiple return values for division (result and error).
- Showcase the use of `defer` in a simulated file operation.

### Instructions:

1.  **Create a new directory** for your Day 5 project:
    ```bash
    cd ~/go/src # or wherever you keep your Go projects
    mkdir day5_functions
    cd day5_functions
    ```
2.  **Create a new file** named `main.go` inside the `day5_functions` directory.
3.  **Open `main.go`** in your code editor and type the following code:

    ```go
    package main

    import (
        "fmt"
        "errors" // Import the errors package for creating custom errors
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
    ```

4.  **Save** the `main.go` file.
5.  **Run the program:** Open your terminal or command prompt, navigate to the `day5_functions` directory, and run:
    ```bash
    go run main.go
    ```
    Observe the output, especially:
    - How multiple values are returned and assigned from the `divide` function.
    - The order of messages in the `simulateFileProcessing` function, particularly how `defer` calls execute after the main function body but before the function exits. Notice the LIFO order for multiple `defer` calls.

### Self-Reflection & Experimentation:

- **Create your own function:** Write a new function that takes two numbers and returns their product. Call it from `main`.
- **Modify parameters/returns:** Change `add` to take `float64`s and return `float64`.
- **Experiment with `defer`:**
  - Add more `defer` statements in `simulateFileProcessing` with different print messages to confirm the LIFO order.
  - Try uncommenting `fmt.Println(1 / 0)` inside `simulateFileProcessing` to see that `defer` functions still run even if a `panic` occurs. This is why `defer` is so crucial for cleanup.
- **Named vs. Unnamed Returns:** If you remove the names from `calculateCircleMetrics`'s return values (e.g., `(float64, float64)`), you'd have to explicitly `return area, circumference`. Try it out.

Functions are truly the backbone of well-structured programs. Mastering them is a huge step forward in your Go journey!

---

Get ready for Day 6, where we'll learn about organizing our Go code into **packages and modules**!
