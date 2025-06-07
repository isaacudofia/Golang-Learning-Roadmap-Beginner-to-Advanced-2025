Alright, welcome to Day 10 of your Go learning journey! You've successfully navigated the world of maps for key-value storage. Today, we're moving on to **Structs**, which are fundamental for defining custom data types in Go, allowing you to group related fields together.

---

### Golang Learning Roadmap: Day 10 - Structs

**Goal for Today:**

- Understand what structs are and why they are used to define custom data types.
- Learn how to declare struct types.
- Practice creating instances (objects) of structs.
- Learn how to access and modify struct fields.
- Explore anonymous structs and embedded structs for composition.

---

**1. Concept: What are Structs?**

A **struct** (short for "structure") is a composite data type that allows you to group together values of different types under a single name. It's similar to classes in object-oriented languages, but Go structs are much simpler; they don't have methods directly within their declaration (methods are added separately, which we'll cover on Day 13).

- **Custom Data Types:** Structs enable you to define your own types that precisely model real-world entities. For example, a `User` struct could have `Name` (string), `Age` (int), and `Email` (string) fields.
- **Grouped Fields:** All fields within a struct are related and represent attributes of the entity the struct describes.
- **Value Type:** Like arrays, structs in Go are **value types**. When you assign one struct to another, or pass a struct to a function, a _copy_ of the entire struct is made. If you need to modify the original struct inside a function, you'll need to pass a pointer to it (we'll cover pointers on Day 14).

**Syntax for Declaring a Struct Type:**

```go
type StructName struct {
    FieldName1 Type1
    FieldName2 Type2
    FieldName3 Type3
    // ...
}
```

- `type`: Keyword to declare a new type.
- `StructName`: The name of your new struct type. Follow Go's naming conventions (start with an uppercase letter if it's exported, lowercase if it's private to the package).
- `struct`: Keyword indicating it's a struct.
- `FieldName Type`: Each line defines a field with its name and data type. Field names also follow export rules (uppercase for exported, lowercase for unexported).

---

**2. Creating Struct Instances (Objects)**

Once you define a struct type, you can create variables of that type, which are called **instances** or **objects**.

- **Zero-Value Initialization:**

  ```go
  var p Person // p will have Name="", Age=0, IsStudent=false
  ```

  All fields are initialized to their respective zero values.

- **Using Composite Literals (Recommended):**
  You can initialize fields directly when creating an instance.

  ```go
  // Order matters (less readable, but works)
  user1 := User{"Alice", 30, "alice@example.com"}

  // Field names explicitly (recommended for readability, order doesn't matter)
  user2 := User{
      Name:    "Bob",
      Age:     25,
      Email:   "bob@example.com",
  }

  // You can omit fields; they'll get their zero values
  user3 := User{Name: "Charlie"} // Age=0, Email=""
  ```

- **Using `new()` (Returns a Pointer):**
  The `new()` built-in function allocates memory for a new variable of a given type and returns a pointer to it. The variable is initialized to its zero value.

  ```go
  p := new(Person) // p is now a pointer to a Person struct: *Person
  fmt.Println(p.Name) // Access fields using dot notation, Go automatically dereferences the pointer for you.
  ```

  We'll cover pointers in more detail on Day 14. For now, focus on composite literals.

---

**3. Accessing and Modifying Struct Fields**

You access individual fields of a struct instance using the **dot operator (`.`)**.

```go
type Book struct {
    Title  string
    Author string
    Pages  int
}

myBook := Book{Title: "Go Programming", Author: "John Doe", Pages: 300}

fmt.Println(myBook.Title) // Access
myBook.Pages = 320        // Modify
fmt.Println(myBook.Pages)
```

---

**4. Anonymous Structs**

You can define structs without explicitly giving them a type name. These are called **anonymous structs** and are often used for temporary, one-off data structures, especially when encoding/decoding JSON.

```go
employee := struct {
    Name    string
    ID      int
    IsActive bool
}{
    Name:    "Jane Doe",
    ID:      101,
    IsActive: true,
}
fmt.Println(employee.Name)
```

---

**5. Embedded Structs (Composition)**

Go doesn't have traditional inheritance. Instead, it uses **composition** through **embedded structs**. When you embed a struct, its fields and methods (we'll cover methods soon) are promoted to the outer struct, allowing for code reuse and building more complex types.

```go
type Person struct {
    Name string
    Age  int
}

type Employee struct {
    Person // Embedded Person struct (no field name, just type)
    EmployeeID string
    Department string
}

emp := Employee{
    Person: Person{Name: "Alice", Age: 30}, // Initialize the embedded struct
    EmployeeID: "E123",
    Department: "Engineering",
}
// Access promoted fields directly:
fmt.Println(emp.Name) // Alice
fmt.Println(emp.Age)  // 30
fmt.Println(emp.EmployeeID) // E123
```

This is a powerful feature for modeling "has-a" relationships (e.g., an `Employee` _has a_ `Person`).

---

**6. Project for Today: User Management System**

We'll define a `User` struct and perform various operations: creating users, accessing their details, updating information, and demonstrating embedding a `Address` struct.

**Instructions:**

1.  **Create a new directory** for your Day 10 project.

    ```bash
    cd ~/go/src # or wherever you keep your Go projects
    mkdir day10_structs
    cd day10_structs
    ```

2.  **Create a new file** named `main.go` inside the `day10_structs` directory.

3.  **Open `main.go` in your code editor** and type the following code:

    ```go
    package main

    import "fmt"

    // Define a struct type for Address
    type Address struct {
        Street  string
        City    string
        ZipCode string
        Country string
    }

    // Define a struct type for User
    type User struct {
        ID        int
        Name      string
        Email     string
        IsActive  bool
        Address   Address // Field of type Address (composition)
        // You could also embed it directly:
        // Address
    }

    // Define a struct type for Product (demonstrating value type behavior later)
    type Product struct {
        Name  string
        Price float64
        Stock int
    }


    func main() {
        fmt.Println("--- Go Structs Demonstration: User Management ---")

        // 1. Create a User instance with zero values
        var user1 User
        fmt.Printf("1. User1 (zero-value): %+v\n", user1) // %+v prints field names and values

        // 2. Create a User instance using composite literal (explicit field names, recommended)
        user2 := User{
            ID:    1,
            Name:  "Alice Wonderland",
            Email: "alice@example.com",
            IsActive: true,
            Address: Address{
                Street:  "123 Main St",
                City:    "Wonderland",
                ZipCode: "90210",
                Country: "Fantasy",
            },
        }
        fmt.Printf("2. User2 (initialized): %+v\n", user2)

        // 3. Accessing struct fields
        fmt.Println("\n--- Accessing User2 Details ---")
        fmt.Printf("User ID: %d\n", user2.ID)
        fmt.Printf("User Name: %s\n", user2.Name)
        fmt.Printf("User Email: %s\n", user2.Email)
        fmt.Printf("User City: %s\n", user2.Address.City) // Accessing embedded struct field

        // 4. Modifying struct fields
        fmt.Println("\n--- Modifying User2 Details ---")
        user2.Email = "alice.new@example.com"
        user2.IsActive = false
        user2.Address.ZipCode = "90211" // Modify embedded struct field
        fmt.Printf("User2 (modified): %+v\n", user2)


        // 5. Create an anonymous struct
        fmt.Println("\n--- Anonymous Struct ---")
        tempProduct := struct {
            PName  string
            PPrice float64
        }{
            PName:  "Temporary Gadget",
            PPrice: 99.99,
        }
        fmt.Printf("Anonymous product: %+v\n", tempProduct)
        fmt.Printf("Anonymous product name: %s\n", tempProduct.PName)


        // 6. Structs are Value Types
        fmt.Println("\n--- Structs as Value Types ---")
        product1 := Product{Name: "Laptop", Price: 1200.0, Stock: 5}
        product2 := product1 // A copy of the struct is made

        fmt.Printf("Product1: %+v\n", product1)
        fmt.Printf("Product2 (copy): %+v\n", product2)

        product2.Stock = 3 // Modify product2's stock
        product2.Name = "Desktop" // Modify product2's name

        fmt.Printf("Product1 after Product2 modification: %+v\n", product1) // Product1 remains unchanged
        fmt.Printf("Product2 after modification: %+v\n", product2)

        // Example of passing a struct by value to a function
        fmt.Println("\n--- Passing Struct by Value to a Function ---")
        printProductDetails(product1) // Pass a copy of product1
        fmt.Printf("Product1 after printProductDetails (in main): %+v\n", product1) // Still unchanged
    }

    // Function to demonstrate passing struct by value
    func printProductDetails(p Product) {
        fmt.Printf("  Inside function - Product Name: %s, Price: %.2f\n", p.Name, p.Price)
        p.Stock = 0 // This change only affects the copy 'p' within this function
        fmt.Println("  Inside function - Modified stock (local copy):", p.Stock)
    }
    ```

4.  **Save the `main.go` file.**

5.  **Run the program:** Open your terminal or command prompt, navigate to the `day10_structs` directory, and run:

    ```bash
    go run main.go
    ```

    Observe the output carefully. Pay attention to:

    - How fields are initialized, accessed, and modified.
    - The output of `%+v` for structs.
    - The behavior of anonymous structs.
    - The crucial "Structs as Value Types" section, showing that `product1` is unaffected when `product2` (its copy) is modified, and when `product1` is passed to `printProductDetails`.

---

**Self-Reflection & Experimentation:**

- **Embed the `Address` struct directly:** Change `Address Address` to just `Address` in the `User` struct. Re-run and confirm that you can still access `user2.Street`, `user2.City`, etc., directly (due to field promotion).
- **Add more fields:** Add a `PhoneNumber` (string) or `DateOfBirth` (e.g., `time.Time` - you'd need `import "time"`) field to the `User` struct.
- **Nested Structs:** Define a `Company` struct that contains a `Name` (string) and a `CEO` (which is a `User` struct).
- **Struct Pointers:** (Preview for Day 14) Instead of `product2 := product1`, try `product2 := &product1` (now `product2` is a pointer). Then modify `product2.Stock` using `product2.Stock = 3` (Go handles dereferencing for you). See how `product1` _does_ change in this case, illustrating reference behavior when using pointers.

Structs are incredibly versatile and will be the backbone of most of your custom data modeling in Go. Understanding them well is key to writing expressive and organized Go code.

Get ready for Day 11, where we'll tie structs and maps together to create more complex data structures!
