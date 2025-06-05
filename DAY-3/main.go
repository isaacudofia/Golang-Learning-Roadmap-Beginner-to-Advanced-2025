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