package main

import (
	"fmt"
	"time"
)

// We'll use this for a short pause in the countdown

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