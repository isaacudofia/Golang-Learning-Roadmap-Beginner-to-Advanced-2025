---
# Go (Golang) Learning Roadmap: Day 9 - Maps

Alright, welcome to Day 9 of your Go learning journey! You've done a fantastic job understanding **arrays** and **slices**, which are sequence-based data structures. Today, we're moving on to **Maps**, a powerful data structure for storing key-value pairs.
---

## Goal for Today:

- Understand what maps are in Go: unordered collections of key-value pairs.
- Learn how to create maps using `make` and map literals.
- Practice adding, retrieving, updating, and deleting elements from a map.
- Learn how to check for the existence of a key in a map.
- Understand iterating over maps using `for...range`.

---

## 1. Concept: What are Maps?

A **map** (also known as a hash map, hash table, or dictionary in other languages) is an **unordered collection of key-value pairs**. Each key in a map must be unique, and it maps to exactly one value.

- **Key-Value Pairs:** Maps store data as pairs, where a unique key is associated with a value.
- **Unordered:** The order in which elements are stored or retrieved from a map is not guaranteed.
- **Unique Keys:** Each key must be unique within a single map. If you try to add a new value with an existing key, the old value will be overwritten.
- **Homogeneous Types:** All keys in a map must be of the same type, and all values must be of the same type.
- **Reference Type:** Like slices, maps are **reference types**. When you assign one map to another, both variables refer to the same underlying map data structure. Changes made through one variable will be visible through the other.

### Valid Key Types:

Map keys can be of any type that is **comparable**, meaning they can be compared using `==` and `!=`. This includes:

- Booleans
- Numeric types (integers, floats, complex numbers)
- Strings
- Arrays (since arrays are comparable)
- Structs (if all their fields are comparable)
- Pointers
- Interfaces (if the dynamic value is comparable)

### Invalid Key Types:

Types that are not comparable cannot be used as map keys. This includes:

- Slices
- Maps
- Functions

---

## 2. Creating Maps

- **Using `make`:** This is the common way to create an empty map. You can optionally provide an initial capacity hint.

  ```go
  // map[KeyType]ValueType
  var ages map[string]int         // Declares a map variable, but it's nil (not initialized)
  ages = make(map[string]int)     // Initializes an empty map

  // Or in one line:
  grades := make(map[string]float64

  // With an initial capacity hint (for performance, not strict size limit)
  users := make(map[int]string, 100) // Creates a map that can ideally hold 100 entries
  ```

- **Using Map Literals (Composite Literals):** To create and initialize a map with values.

  ```go
  colors := map[string]string{
      "red":   "#FF0000",
      "green": "#00FF00",
      "blue":  "#0000FF",
  }
  // You can omit the type if the elements are all of the same type as inferred
  // studentScores := map[string]int{"Alice": 90, "Bob": 85}
  ```

---

## 3. Map Operations

- **Adding/Updating Elements:**
  Use the key in square brackets to assign a value. If the key doesn't exist, it's added. If it exists, its value is updated.

  ```go
  myMap := make(map[string]int)
  myMap["apple"] = 10     // Add
  myMap["banana"] = 20    // Add
  myMap["apple"] = 15     // Update
  fmt.Println(myMap)      // Output: map[apple:15 banana:20] (order might vary)
  ```

- **Retrieving Elements:**
  Use the key in square brackets. If the key exists, you get its value. If it doesn't exist, you get the **zero value** for the map's value type.

  ```go
  count := myMap["banana"] // count will be 20
  invalidKeyCount := myMap["orange"] // invalidKeyCount will be 0 (zero value for int)
  ```

- **Checking for Existence (Comma Ok Idiom):**
  This is crucial! When retrieving a value, you can get a second boolean return value that indicates whether the key actually existed in the map. This is known as the "`comma ok`" idiom.

  ```go
  value, ok := myMap["banana"]
  if ok {
      fmt.Printf("Banana count: %d\n", value)
  } else {
      fmt.Println("Banana not found.")
  }

  value, ok = myMap["orange"]
  if ok {
      fmt.Printf("Orange count: %d\n", value)
  } else {
      fmt.Println("Orange not found.")
  }
  ```

- **Deleting Elements:**
  Use the `delete()` built-in function. Deleting a non-existent key is a no-op (doesn't cause an error).

  ```go
  delete(myMap, "banana")
  fmt.Println(myMap) // Output: map[apple:15]
  ```

- **Length of a Map:**
  Use `len()` to get the number of key-value pairs in a map.

  ```go
  fmt.Println("Map length:", len(myMap))
  ```

- **Iterating over Maps (`for...range`):**
  You can use `for...range` to iterate over all key-value pairs in a map. The order of iteration is not guaranteed.

  ```go
  for key, value := range colors {
      fmt.Printf("Key: %s, Value: %s\n", key, value)
  }

  // If you only need keys:
  for key := range colors {
      fmt.Println("Key:", key)
  }

  // If you only need values:
  for _, value := range colors { // Use blank identifier for key
      fmt.Println("Value:", value)
  }
  ```

---

## 4. Project for Today: Simple Dictionary/Inventory

We'll build a program that simulates a basic dictionary or inventory system using maps. You'll perform add, retrieve, update, and delete operations.

### Instructions:

1.  **Create a new directory** for your Day 9 project:
    ```bash
    cd ~/go/src # or wherever you keep your Go projects
    mkdir day9_maps
    cd day9_maps
    ```
2.  **Create a new file** named `main.go` inside the `day9_maps` directory.
3.  **Open `main.go`** in your code editor and type the following code:

    ```go
    package main

    import "fmt"

    func main() {
        fmt.Println("--- Go Maps Demonstration: Simple Inventory ---")

        // 1. Create a map for inventory (item name -> quantity)
        inventory := make(map[string]int)
        fmt.Printf("1. Initial inventory: %v (len: %d)\n", inventory, len(inventory))

        // 2. Add items to inventory
        inventory["Laptop"] = 5
        inventory["Mouse"] = 12
        inventory["Keyboard"] = 8
        fmt.Printf("2. Inventory after adding items: %v (len: %d)\n", inventory, len(inventory))

        // 3. Retrieve item quantity
        fmt.Println("\n--- Retrieving Items ---")
        laptopQty := inventory["Laptop"]
        fmt.Printf("Quantity of Laptop: %d\n", laptopQty)

        // Trying to retrieve a non-existent item (will get zero value)
        monitorQty := inventory["Monitor"]
        fmt.Printf("Quantity of Monitor (non-existent): %d\n", monitorQty) // Output: 0

        // 4. Check for item existence (comma ok idiom)
        fmt.Println("\n--- Checking for Existence ---")
        keyboardQty, ok := inventory["Keyboard"]
        if ok {
            fmt.Printf("Keyboard exists. Quantity: %d\n", keyboardQty)
        } else {
            fmt.Println("Keyboard does not exist.")
        }

        speakerQty, ok := inventory["Speakers"]
        if ok {
            fmt.Printf("Speakers exist. Quantity: %d\n", speakerQty)
        } else {
            fmt.Println("Speakers does not exist.")
        }

        // 5. Update item quantity
        fmt.Println("\n--- Updating Items ---")
        inventory["Mouse"] = 15 // Update quantity of Mouse
        fmt.Printf("Inventory after updating Mouse: %v\n", inventory)

        // 6. Delete an item
        fmt.Println("\n--- Deleting Items ---")
        delete(inventory, "Keyboard")
        fmt.Printf("Inventory after deleting Keyboard: %v (len: %d)\n", inventory, len(inventory))

        // Try deleting a non-existent item (no error)
        delete(inventory, "Headphones")
        fmt.Printf("Inventory after deleting non-existent Headphones: %v (len: %d)\n", inventory, len(inventory))

        // 7. Iterate over the map using for-range
        fmt.Println("\n--- Iterating Over Inventory ---")
        fmt.Println("Current Stock:")
        for item, qty := range inventory {
            fmt.Printf("  %s: %d units\n", item, qty)
        }

        // 8. Map as a Reference Type
        fmt.Println("\n--- Map as Reference Type ---")
        originalMap := map[string]int{"A": 1, "B": 2}
        referredMap := originalMap // Both now refer to the same underlying map

        referredMap["C"] = 3    // Add through referredMap
        delete(originalMap, "A") // Delete through originalMap

        fmt.Printf("Original Map: %v\n", originalMap) // Changes are reflected
        fmt.Printf("Referred Map: %v\n", referredMap)   // Changes are reflected
    }
    ```

4.  **Save** the `main.go` file.
5.  **Run the program:** Open your terminal or command prompt, navigate to the `day9_maps` directory, and run:
    ```bash
    go run main.go
    ```
    Observe the output and how each map operation affects the data. Pay special attention to the "comma ok" idiom for checking existence and the reference-type behavior.

### Self-Reflection & Experimentation:

- **Change Key/Value Types:** Create a map where keys are `int` and values are `string` (e.g., `employeeIDs := make(map[int]string)`).
- **Map of Slices/Maps:** Try to create a map where values are slices or even other maps (e.g., `userPreferences := make(map[string][]string)`). This is very common.
- **Invalid Key Types:** Try to use a slice as a map key (e.g., `invalidMap := map[[]int]string{}`). What compile-time error do you get?
- **Iterate only keys or values:** Modify the `for...range` loop to print only the item names or only the `qty` values.
- **Nil Map:** What happens if you try to add elements to a `nil` map (e.g., `var myMap map[string]int; myMap["test"] = 1`)? (It will cause a `panic`). This is why you must always `make` or initialize a map before using it.

Maps are incredibly useful for quick lookups and managing collections where items are identified by unique keys. You'll use them constantly in Go programming.

---

Get ready for Day 10, where we'll delve into a powerful concept for organizing custom data: **Structs**!
