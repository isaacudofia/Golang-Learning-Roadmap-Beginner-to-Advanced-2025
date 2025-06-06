---
# Go (Golang) Learning Roadmap: Day 7 - Arrays

Alright, welcome to Day 7 of your Go learning journey! You've just mastered organizing your code with **packages and modules**, a significant leap towards building larger applications. Today, we're going to dive into the first of Go's built-in data structures: **Arrays**.

---

## Goal for Today:

- Understand what arrays are in Go: fixed-size, contiguous data structures.
- Learn how to declare and initialize arrays.
- Practice accessing and modifying array elements.
- Understand the concept of array length.
- Perform basic operations like summing elements or finding min/max.

---

## 1. Concept: What are Arrays?

In Go, an array is a **fixed-size** sequence of elements of the **same type**, stored contiguously in memory.

- **Fixed Size:** Once an array is declared with a certain size, its size cannot be changed. This is a crucial characteristic.
- **Homogeneous:** All elements in an array must be of the same data type (e.g., all `int`s, all `string`s).
- **Zero-Indexed:** Array elements are accessed using an index, starting from `0` for the first element up to `length - 1` for the last.
- **Value Type:** Arrays in Go are **value types**. This means when you assign one array to another, or pass an array to a function, a **copy of the entire array is made**. This is an important distinction compared to many other languages where arrays are reference types.

### Syntax for Declaring and Initializing Arrays:

- **Declaration with Zero Values:**
  ```go
  var a [5]int // Declares an array 'a' of 5 integers, all initialized to 0 (zero value for int)
  ```
- **Declaration and Initialization:**
  ```go
  var b [3]string = [3]string{"apple", "banana", "cherry"}
  ```
- **Short Variable Declaration with Initialization:**
  ```go
  c := [4]float64{1.1, 2.2, 3.3, 4.4}
  ```
- **Compiler Inferring Length (`...`):** You can let the compiler count the number of elements by using `...` instead of a specific length. This is very common.
  ```go
  d := [...]int{10, 20, 30} // d will be an array of length 3
  ```
- **Initializing Specific Elements:**
  ```go
  e := [5]int{1: 10, 3: 30} // e will be [0, 10, 0, 30, 0]
  ```

### Accessing Elements:

You access individual elements using square brackets `[]` with their index.

```go
myArray := [3]string{"Go", "is", "fun"}
fmt.Println(myArray[0]) // Output: Go
myArray[1] = "was"      // Modify an element
fmt.Println(myArray[1]) // Output: was
```

### Array Length:

You can get the length of an array using the built-in `len()` function.

```go
myArray := [...]int{1, 2, 3, 4, 5}
fmt.Println(len(myArray)) // Output: 5
```

---

## 2. Project for Today: Array Operations

We'll create a program that demonstrates array declaration, initialization, access, modification, and basic operations like summing elements and finding the maximum value.

### Instructions:

1.  **Create a new directory** for your Day 7 project:
    ```bash
    cd ~/go/src # or wherever you keep your Go projects
    mkdir day7_arrays
    cd day7_arrays
    ```
2.  **Create a new file** named `main.go` inside the `day7_arrays` directory.
3.  **Open `main.go`** in your code editor and type the following code:

    ```go
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
    ```

4.  **Save** the `main.go` file.
5.  **Run the program:** Open your terminal or command prompt, navigate to the `day7_arrays` directory, and run:
    ```bash
    go run main.go
    ```
    Carefully observe the output for each section, paying attention to how array elements are accessed, modified, and how the `len()` function works. Also, note the clear distinction of arrays being value types demonstrated in section 8.

### Self-Reflection & Experimentation:

- **Change Array Sizes and Types:** Declare arrays of different sizes and types (e.g., `[10]bool`, `[2]float32`).
- **Experiment with Initialization:** Try initializing an array using specific indices like `data := [5]int{0: 1, 4: 5}`. What are the intermediate values?
- **Out-of-Bounds Access:** Uncomment the line `// fmt.Println(numbers[5])` and run the program. Observe the runtime panic message. This teaches you about preventing such errors.
- **Pass array to a function:**
  - Create a function `func printArray(arr [5]int)` that takes an array as input and prints it.
  - Call this function from `main`.
  - Then, inside `printArray`, try to modify an element (e.g., `arr[0] = 99`).
  - After the function call, print the original array in `main` again. Does it change? (It shouldn't, because arrays are value types and a copy was passed).

Arrays are a foundational data structure, but their fixed size and value-type nature often lead to the use of **slices** in Go for more dynamic collections. This is exactly what we'll cover tomorrow!

---

Get ready for Day 8, where we'll dive into the much more flexible and commonly used **slices**!
