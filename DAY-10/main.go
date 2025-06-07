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