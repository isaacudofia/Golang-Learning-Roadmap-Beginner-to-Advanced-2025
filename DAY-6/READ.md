---
# Go (Golang) Learning Roadmap: Day 6 - Packages and Modules

Alright, welcome to Day 6 of your Go learning journey! You've done a great job grasping **functions** on Day 5, which are key to code organization. Today, we're taking that concept further by exploring **packages and modules**, the fundamental ways Go organizes and manages code at a higher level.
---

## Goal for Today:

- Understand what Go **packages** are and their role in code organization.
- Learn how to define and **export** elements (functions, variables) from a package.
- Understand the `main` package and the `import` statement.
- Grasp the concept of **Go Modules** for dependency management.
- Practice creating a multi-file, multi-package project with a Go module.

---

## 1. Concept: Go Packages

A **package** is a way to organize Go code. It's a collection of source files in the same directory that are compiled together. All functions, types, variables, and constants declared in one source file of a package are visible to all other source files in the same package.

- **Naming Convention:**

  - The package name is typically the last element of its import path (e.g., `fmt` from `import "fmt"`).
  - It should be all lowercase.

- **Visibility (Exporting):**

  - An identifier (function name, variable name, type name) is **exported** (publicly visible outside the package) if its first letter is **uppercase**.
  - An identifier is **unexported** (private, only visible within the package) if its first letter is **lowercase**.

  **Example:** If you have a package named `calculator`:

  - `func Add(a, b int) int` would be exported and usable by other packages.
  - `func subtract(a, b int) int` would be unexported and only usable within the `calculator` package.

- **The `main` Package:**
  The `main` package is special. It defines a standalone executable program. It must contain a `main` function, which is the entry point of the program.

- **The `import` Statement:**
  The `import` statement brings packages into the current file's scope, allowing you to use their exported identifiers.

  ```go
  import "fmt"       // Imports the fmt standard library package
  import "math"      // Imports another standard library package

  // Multiple imports can be grouped using parentheses
  import (
      "fmt"
      "net/http" // Example of a sub-package
      "time"
  )
  ```

  You refer to exported identifiers using the package name followed by a dot (e.g., `fmt.Println`, `math.Pi`).

---

## 2. Concept: Go Modules

Before Go modules (introduced in Go 1.11), dependency management was often handled by tools like `dep` or by relying on the `GOPATH`. **Go Modules** are the standard and recommended way to manage dependencies in Go projects.

- **What is a Module?** A module is a collection of related Go packages. It's the unit of source code versioning.
- **`go.mod` file:** Every module has a `go.mod` file at its root, which defines the module's path, Go version, and lists its direct and indirect dependencies with their specific versions.
- **`go.sum` file:** This file contains cryptographic checksums of the module's dependencies to ensure the integrity and authenticity of downloaded modules.
- **Initialization:** You initialize a new module in a project directory using `go mod init <module_path>`. The `module_path` is usually your repository path (e.g., `github.com/yourusername/yourproject`).
- **Automatic Management:** When you `go run`, `go build`, or `go test`, Go automatically downloads and manages the necessary dependencies listed in `go.mod`.
- **Vendoring (Optional):** `go mod vendor` copies all direct and indirect module dependencies into a local `vendor/` directory. This is useful for environments with restricted internet access.

---

## 3. Project for Today: A Simple Calculator with Packages and Modules

We'll create a Go module that contains two packages:

- `main`: The entry point of our application.
- `calculator`: A custom package that provides arithmetic functions.

This will demonstrate how to structure a multi-package application within a module.

### Instructions:

1.  **Create a main project directory:**

    ```bash
    cd ~/go/src # or wherever you keep your Go projects
    mkdir calculator_app
    cd calculator_app
    ```

2.  **Initialize the Go module:**
    This command creates the `go.mod` file in your `calculator_app` directory. The module path is crucial – it's how other projects would import your packages. For local practice, a simple name like `example.com/calculator_app` is fine. In a real project, this would typically be your Git repository path (e.g., `github.com/yourusername/calculator_app`).

    ```bash
    go mod init example.com/calculator_app
    ```

    You should see a `go.mod` file created. Its content will look something like:

    ```go
    module example.com/calculator_app

    go 1.22 // or your current Go version
    ```

3.  **Create the `calculator` package directory:**

    ```bash
    mkdir calculator
    ```

4.  **Create `calculator/add.go`:**
    Open `calculator/add.go` and add the following code:

    ```go
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
    ```

5.  **Create `calculator/multiply.go`:**
    Open `calculator/multiply.go` and add the following code:

    ```go
    // calculator/multiply.go
    package calculator // This file also belongs to the 'calculator' package

    // Multiply exports the multiplication function.
    func Multiply(a, b int) int {
        return a * b
    }
    ```

    **Note:** Both `add.go` and `multiply.go` declare `package calculator`. This means they are part of the same `calculator` package.

6.  **Create `main.go` in the root `calculator_app` directory:**
    Open `main.go` and add the following code:

    ```go
    // main.go
    package main // This is the main executable package

    import (
        "fmt"
        // Import our custom calculator package.
        // The path is the module path followed by the package directory.
        "example.com/calculator_app/calculator"
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
    ```

7.  **Save all files.**

8.  **Run the program:**
    Make sure you are in the root `calculator_app` directory in your terminal.

    ```bash
    go run .
    ```

    (The `.` tells Go to run the main package in the current directory, discovering its dependencies from `go.mod` and other files in the module).

    You should see output similar to this:

    ```
    --- Simple Calculator App ---
    10 + 5 = 15
    10 * 5 = 50

    To see an error, uncomment the line trying to call `calculator.subtract` in main.go.
    It won't compile because `subtract` is not exported (starts with lowercase).
    ```

### Experiment with the unexported function:

- Uncomment the line `// diff := calculator.subtract(num1, num2)` in `main.go`.
- Try to run `go run .` again.
- You should now get a compile-time error like: `calculator.subtract undefined (cannot refer to unexported name calculator.subtract)`. This demonstrates the visibility rules!

### Self-Reflection & Experimentation:

- **Create another function in `calculator`:** Add a `Divide` function to `calculator/multiply.go` (or a new file `calculator/divide.go`). Make sure it's exported!
- **Import the `errors` package in `calculator`:** Modify your `Divide` function to return an error if the denominator is zero, similar to what you did on Day 5.
- **Add a new sub-package:** Create a `utils` directory inside `calculator_app`, declare `package utils` in a new file, and add a simple exported function like `PrintGreeting(name string)`. Then import `example.com/calculator_app/utils` into `main.go` and use it.
- **Explore `go mod tidy`:** If you add a new dependency (e.g., `github.com/google/uuid`), Go will automatically add it to `go.mod` when you run code that uses it. `go mod tidy` will clean up unused dependencies.

You've now moved beyond single-file programs and understand how to structure your Go applications into reusable packages managed by modules. This is a crucial step towards building larger, more organized projects!

---

Get ready for Day 7, where we'll delve into the foundational data structure: **Arrays**!
