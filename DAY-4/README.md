---
# Go (Golang) Learning Roadmap: Day 4 - Loops (For Loop)

Welcome to Day 4 of your Go learning journey! Over the past three days, you've grasped the basics of Go setup, variables, constants, and conditional logic. Today, we're going to tackle **loops**, which are fundamental for repeating actions in your programs.
---

## Goal for Today:

- Understand Go's single **`for` loop** construct.
- Learn how to use `for` as a traditional `for` loop, a `while` loop, and an infinite loop.
- Explore the **`range` keyword** for iterating over collections.
- Practice using **`break`** and **`continue`** to control loop execution.

---

## 1. Concept: Go's Unified `for` Loop

Unlike many other languages that have `for`, `while`, and `do-while` loops, Go only has **one** looping construct: the `for` loop. This single construct is incredibly versatile and can be used to achieve all common looping patterns.

### Syntax Variations:

1.  **The Classic `for` Loop** (with initialization, condition, and post-statement):
    This is the most common form, similar to `for` loops in C, Java, or JavaScript.

    ```go
    for initialization; condition; post-statement {
        // code to execute repeatedly
    }
    ```

    - **`initialization`**: An optional statement executed once before the first iteration.
    - **`condition`**: A boolean expression evaluated before each iteration. If `true`, the loop continues; if `false`, the loop terminates.
    - **`post-statement`**: An optional statement executed after each iteration (e.g., incrementing a counter).

    **Example:**

    ```go
    for i := 0; i < 5; i++ {
        fmt.Println("Count:", i)
    }
    ```

2.  **The "While" Loop** (with only a condition):
    If you omit the initialization and post-statement, the `for` loop acts like a `while` loop, repeating as long as the condition is true.

    ```go
    for condition {
        // code to execute repeatedly
    }
    ```

    **Example:**

    ```go
    sum := 1
    for sum < 1000 {
        sum += sum
    }
    fmt.Println("Sum (while-like):", sum)
    ```

3.  **The Infinite Loop:**
    If you omit all parts of the `for` clause, you create an **infinite loop**. You'll typically use `break` inside such a loop to exit.

    ```go
    for {
        // code that runs forever unless a 'break' is encountered
    }
    ```

    **Example:**

    ```go
    counter := 0
    for {
        fmt.Println("Looping infinitely (until break):", counter)
        counter++
        if counter == 3 {
            break // Exit the loop
        }
    }
    ```

---

## 2. Concept: `break` and `continue`

- **`break`**: Terminates the innermost `for` loop (or `switch` statement) and continues execution immediately after the loop.
- **`continue`**: Skips the rest of the current iteration of the loop and proceeds to the next iteration (evaluating the condition and executing the post-statement).

**Example:**

```go
for i := 0; i < 10; i++ {
    if i%2 != 0 { // If i is odd
        continue // Skip to the next iteration
    }
    if i == 8 {
        break // Stop the loop when i is 8
    }
    fmt.Println("Even number:", i)
}
```

---

## 3. Concept: `for...range` Loop (Iterating over Collections)

The **`for...range`** loop is used to iterate over elements of collections like **arrays, slices, strings, and maps**. It returns two values for each element: the index/key and the value.

- **Over Slices/Arrays:** Returns `index`, `value`.

  ```go
  numbers := []int{10, 20, 30, 40}
  for index, value := range numbers {
      fmt.Printf("Index: %d, Value: %d\n", index, value)
  }
  ```

  If you only need the value, you can ignore the index using the **blank identifier `_`**:

  ```go
  for _, value := range numbers {
      fmt.Println("Value:", value)
  }
  ```

- **Over Strings:** Returns **byte index** (of the start of the rune) and **rune** (Unicode character).

  ```go
  greeting := "Hello, Go!"
  for index, char := range greeting {
      fmt.Printf("Byte Index: %d, Rune: %c\n", index, char)
  }
  ```

- **Over Maps** (we'll cover maps in more detail later): Returns `key`, `value`.
  ```go
  // This is just a preview; we'll cover maps on Day 11
  ages := map[string]int{"Alice": 30, "Bob": 25}
  for name, age := range ages {
      fmt.Printf("%s is %d years old.\n", name, age)
  }
  ```

---

## 4. Project for Today: Summation, Countdown, and List Processing

We'll create a single program that uses different forms of the `for` loop to:

- Calculate the sum of numbers in a range.
- Implement a countdown timer.
- Process elements from a list (slice).

### Instructions:

1.  **Create a new directory** for your Day 4 project:
    ```bash
    cd ~/go/src # or wherever you keep your Go projects
    mkdir day4_loops
    cd day4_loops
    ```
2.  **Create a new file** named `main.go` inside the `day4_loops` directory.
3.  **Open `main.go`** in your code editor and type the following code:

    ```go
    package main

    import (
        "fmt"
        "time" // We'll use this for a short pause in the countdown
    )

    func main() {
        fmt.Println("--- 1. Classic For Loop: Sum of Numbers ---")
        sum := 0
        for i := 1; i <= 10; i++ {
            sum += i // sum = sum + i
        }
        fmt.Printf("Sum of numbers from 1 to 10: %d\n", sum)

        // --- 2. For Loop as a "While" Loop: Countdown ---
        fmt.Println("\n--- 2. For Loop as a 'While' Loop: Countdown ---")
        count := 5
        for count > 0 {
            fmt.Printf("Countdown: %d...\n", count)
            time.Sleep(time.Second) // Pause for 1 second (Requires "time" import)
            count--
        }
        fmt.Println("Blast off!")

        // --- 3. Infinite For Loop with Break: Password Attempt Simulation ---
        fmt.Println("\n--- 3. Infinite For Loop with Break: Password Attempts ---")
        attempts := 0
        correctPassword := "gosecret"
        enteredPassword := "" // Simulate user input

        for {
            attempts++
            fmt.Printf("Attempt %d: Enter password: ", attempts)
            // In a real app, you'd get input here, e.g., fmt.Scanln(&enteredPassword)
            // For this example, let's just predefine attempts
            if attempts == 1 {
                enteredPassword = "wrong"
            } else if attempts == 2 {
                enteredPassword = "gosecret" // Correct password on 2nd attempt
            } else {
                enteredPassword = "another_wrong"
            }


            if enteredPassword == correctPassword {
                fmt.Println("Access Granted!")
                break // Exit the loop on success
            } else {
                fmt.Println("Incorrect password.")
            }

            if attempts >= 3 {
                fmt.Println("Too many failed attempts. Account locked.")
                break // Exit loop after 3 attempts
            }
        }

        // --- 4. For-Range Loop: Iterate over a Slice ---
        fmt.Println("\n--- 4. For-Range Loop: Processing Items ---")
        fruits := []string{"Apple", "Banana", "Cherry", "Date"}

        fmt.Println("Listing fruits with index:")
        for i, fruit := range fruits {
            fmt.Printf("Fruit #%d: %s\n", i+1, fruit) // i is 0-indexed, so add 1 for display
        }

        fmt.Println("\nListing fruits (value only):")
        for _, fruit := range fruits { // Using blank identifier for index
            fmt.Println("->", fruit)
        }

        // --- 5. For-Range Loop: Iterate over a String (Runes) ---
        fmt.Println("\n--- 5. For-Range Loop: Iterating over a String ---")
        goMessage := "你好 Go!" // Contains Unicode characters
        for i, r := range goMessage {
            fmt.Printf("Char at byte index %d: '%c' (Unicode: %U)\n", i, r, r)
        }
    }
    ```

4.  **Save** the `main.go` file.
5.  **Run the program:** Open your terminal or command prompt, navigate to the `day4_loops` directory, and run:
    ```bash
    go run main.go
    ```
    Observe the output from each section of the program. Notice how the countdown pauses and how the password attempt simulation works. Pay attention to the `for-range` output for strings, especially with Unicode characters.

### Self-Reflection & Experimentation:

- **Modify Loop Conditions:** Try changing the conditions in the `for` loops (e.g., sum from 1 to 100, countdown from 10).
- **Experiment with `break` and `continue`:**
  - In the "Sum of Numbers" loop, add a `continue` statement to skip adding even numbers.
  - In the "Password Attempts" loop, change the `break` condition or remove it to see what happens.
- **Modify the `fruits` slice:** Add or remove elements from the `fruits` slice and observe how the `for-range` loop adapts.
- **Try iterating over an empty slice:** What happens when you use `for-range` on an empty slice? (It simply won't execute any iterations).

You've now got a solid understanding of how to control repetition in your Go programs, a fundamental skill for building dynamic applications!

---

Get ready for Day 5, where we'll explore **functions** in detail!
