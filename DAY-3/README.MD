---
# Go (Golang) Learning Roadmap: Day 3 - Control Flow (If/Else and Switch Statements)

Welcome to Day 3 of your Go learning journey! Yesterday, you mastered **variables**, **constants**, and **basic data types**. Today, we'll dive into **control flow**, which allows your programs to make decisions and execute different blocks of code based on specific conditions.
---

## Goal for Today:

- Understand how to use `if`, `else if`, and `else` statements for conditional execution.
- Learn about `switch` statements for multi-way branching.
- Explore `switch` statement variations, including using no condition (tagless switch) and `fallthrough`.
- Practice writing simple decision-making logic.

---

## 1. Concept: `if`, `else if`, and `else` Statements

The `if` statement is used to execute a block of code only if a specified condition is true. You can extend this with `else if` for additional conditions and `else` for a fallback block if none of the preceding conditions are met.

### Syntax:

```go
if condition1 {
    // code to execute if condition1 is true
} else if condition2 {
    // code to execute if condition1 is false AND condition2 is true
} else {
    // code to execute if all preceding conditions are false
}
```

### Key Points in Go:

- **No Parentheses:** Unlike many other languages (like C++, Java, JavaScript), the condition in Go's `if` statement does **not** require parentheses.
- **Braces are Mandatory:** The curly braces `{}` around the code block are always required, even if there's only one statement.
- **Optional `else if` and `else`:** You can have an `if` statement by itself, or with just an `else`, or with multiple `else if` clauses.
- **Short Statement Form (Initialization Statement):** Go allows you to include a short statement before the condition. This is often used for variable declarations whose scope is limited _only_ to the `if`/`else if`/`else` block. This is a very common Go idiom.

  ```go
  if num := 10; num > 0 { // num is declared and initialized here
      fmt.Println("Number is positive:", num)
  } else {
      fmt.Println("Number is non-positive:", num)
  }
  // fmt.Println(num) // This would cause a compile-time error: undefined: num
  ```

---

## 2. Concept: `switch` Statements

The `switch` statement provides a concise way to handle multiple conditional branches based on the value of an expression. It's often cleaner than a long chain of `if-else if` statements.

### Syntax:

```go
switch expression {
case value1:
    // code to execute if expression == value1
case value2, value3: // Multiple values can be in one case
    // code to execute if expression == value2 OR expression == value3
default:
    // code to execute if no other case matches (optional)
}
```

### Key Points in Go:

- **Implicit `break`:** Unlike C/C++ or Java, Go's `switch` statements implicitly `break` after a `case` block. You **don't** need a `break` keyword. Execution stops after the first matching case.
- **`fallthrough`:** If you _do_ want execution to continue into the next case, you can explicitly use the `fallthrough` keyword. This is rarely needed but available.
- **No Expression `switch` (Tagless Switch):** A `switch` statement can be used without an expression. In this case, each `case` statement is a boolean expression, and the `switch` behaves like a more flexible `if-else if` chain. The first `case` that evaluates to true is executed.

  ```go
  score := 85
  switch { // No expression here
  case score >= 90:
      fmt.Println("Excellent!")
  case score >= 70:
      fmt.Println("Good!")
  default:
      fmt.Println("Needs improvement.")
  }
  ```

- **Short Statement Form:** Like `if` statements, `switch` statements can also include an initialization statement.

  ```go
  switch day := "Sunday"; day {
  case "Saturday", "Sunday":
      fmt.Println("It's the weekend!")
  default:
      fmt.Println("It's a weekday.")
  }
  ```

---

## 3. Project for Today: Grade Calculator & Day of the Week

We'll create two programs that demonstrate `if/else` and `switch` statements:

1.  A simple grade calculator based on a score.
2.  A program that determines the type of day (weekday/weekend) based on a given day number.

### Instructions:

1.  **Create a new directory** for your Day 3 project:
    ```bash
    cd ~/go/src # or wherever you keep your Go projects
    mkdir day3_control_flow
    cd day3_control_flow
    ```
2.  **Create a new file** named `main.go` inside the `day3_control_flow` directory.
3.  **Open `main.go`** in your code editor and type the following code:

    ```go
    package main

    import "fmt"

    func main() {
        // --- Part 1: Grade Calculator using if/else if/else ---
        fmt.Println("--- Grade Calculator ---")

        score := 78 // Try changing this score

        if score >= 90 {
            fmt.Printf("Score: %d, Grade: A (Excellent!)\n", score)
        } else if score >= 80 {
            fmt.Printf("Score: %d, Grade: B (Very Good)\n", score)
        } else if score >= 70 {
            fmt.Printf("Score: %d, Grade: C (Good)\n", score)
        } else if score >= 60 {
            fmt.Printf("Score: %d, Grade: D (Pass)\n", score)
        } else {
            fmt.Printf("Score: %d, Grade: F (Fail)\n", score)
        }

        fmt.Println("\n--- Number Classification (using short statement form) ---")
        if num := 10; num > 0 { // num is scoped to this if-else block
            fmt.Printf("%d is a positive number.\n", num)
        } else if num < 0 {
            fmt.Printf("%d is a negative number.\n", num)
        } else {
            fmt.Printf("%d is zero.\n", num)
        }

        // --- Part 2: Day of the Week using switch ---
        fmt.Println("\n--- Day of the Week ---")

        dayOfWeek := 3 // 1=Sunday, 2=Monday, ..., 7=Saturday (Try changing this number)

        switch dayOfWeek {
        case 1:
            fmt.Println("It's Sunday - A weekend day!")
        case 2:
            fmt.Println("It's Monday - Start of the week!")
        case 3:
            fmt.Println("It's Tuesday.")
        case 4:
            fmt.Println("It's Wednesday.")
        case 5:
            fmt.Println("It's Thursday.")
        case 6:
            fmt.Println("It's Friday - Almost weekend!")
        case 7:
            fmt.Println("It's Saturday - A weekend day!")
        default:
            fmt.Println("Invalid day number. Please enter a number between 1 and 7.")
        }

        // --- Part 3: Tagless Switch (like an if-else if chain) ---
        fmt.Println("\n--- Activity Suggestion (Tagless Switch) ---")
        temperature := 25 // Try changing this temperature

        switch { // No expression here, cases are boolean expressions
        case temperature < 10:
            fmt.Printf("Temperature is %d°C. It's cold, stay indoors!\n", temperature)
        case temperature >= 10 && temperature < 20:
            fmt.Printf("Temperature is %d°C. Nice weather for a walk.\n", temperature)
        case temperature >= 20 && temperature < 30:
            fmt.Printf("Temperature is %d°C. Perfect for outdoor activities!\n", temperature)
        default:
            fmt.Printf("Temperature is %d°C. It's very hot, find some shade!\n", temperature)
        }

        // --- Part 4: Using fallthrough (use sparingly, typically not Go idiomatic) ---
        fmt.Println("\n--- Fallthrough Example ---")
        city := "London" // Try changing this to "Paris" or "New York"

        switch city {
        case "New York":
            fmt.Println("Big Apple!")
            fallthrough // This will cause the next case to execute
        case "London":
            fmt.Println("Capital of the UK!")
            fallthrough // This will cause the next case to execute
        case "Paris":
            fmt.Println("City of Love!")
        default:
            fmt.Println("Unknown city.")
        }
    }
    ```

4.  **Save** the `main.go` file.
5.  **Run the program:** Open your terminal or command prompt, navigate to the `day3_control_flow` directory, and run:
    ```bash
    go run main.go
    ```
    Observe the output and how different conditions lead to different execution paths.

### Self-Reflection & Experimentation:

- **Change values:** Modify the `score`, `dayOfWeek`, `temperature`, and `city` variables and run the program again to see how the output changes.
- **Experiment with the `switch` expression:**
  - What happens if you use a string as the `switch` expression and `case` values? (It works!)
  - What if you try to put a value in `case` that doesn't match the `switch` expression's type? (Compile-time error).
- **Remove `else if` or `else`:** See how the logic changes if you remove parts of the `if-else if-else` chain.
- **Modify `fallthrough`:** Comment out `fallthrough` in the example to see the default `switch` behavior. Add another `fallthrough` if you want to see it continue further.

You've now learned how to add decision-making logic to your Go programs, a crucial skill for any programmer!

---

Get ready for Day 4, where we'll focus on loops and repetition!
