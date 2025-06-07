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

    referredMap["C"] = 3 // Add through referredMap
    delete(originalMap, "A") // Delete through originalMap

    fmt.Printf("Original Map: %v\n", originalMap) // Changes are reflected
    fmt.Printf("Referred Map: %v\n", referredMap)   // Changes are reflected
}