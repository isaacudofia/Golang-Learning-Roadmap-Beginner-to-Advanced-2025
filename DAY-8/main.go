package main

import "fmt"

func main() {
    fmt.Println("--- Go Slices Demonstration: Dynamic List ---")

    // 1. Create a slice using make (length 0, capacity 5)
    todoList := make([]string, 0, 5)
    fmt.Printf("1. Initial todoList: %v (len: %d, cap: %d)\n", todoList, len(todoList), cap(todoList))

    // 2. Append items to the list
    todoList = append(todoList, "Learn Go slices")
    todoList = append(todoList, "Practice slice operations")
    todoList = append(todoList, "Understand underlying arrays")
    fmt.Printf("2. todoList after appending 3 items: %v (len: %d, cap: %d)\n", todoList, len(todoList), cap(todoList))

    // 3. Append more items (exceeding initial capacity)
    todoList = append(todoList, "Review slice examples")
    todoList = append(todoList, "Experiment with append")
    todoList = append(todoList, "Prepare for Day 9") // This will likely trigger a reallocation
    fmt.Printf("3. todoList after appending more (reallocation?): %v (len: %d, cap: %d)\n", todoList, len(todoList), cap(todoList))

    // 4. Accessing and modifying elements
    fmt.Println("\n--- Accessing and Modifying ---")
    fmt.Printf("First task: %s\n", todoList[0])
    todoList[1] = "Master slice operations" // Modify an element
    fmt.Printf("Modified second task: %s\n", todoList[1])
    fmt.Printf("Current todoList: %v\n", todoList)

    // 5. Slicing the slice
    fmt.Println("\n--- Slicing Examples ---")
    firstTwoTasks := todoList[:2] // Get the first two tasks
    fmt.Printf("First two tasks: %v (len: %d, cap: %d)\n", firstTwoTasks, len(firstTwoTasks), cap(firstTwoTasks))

    remainingTasks := todoList[2:] // Get the remaining tasks
    fmt.Printf("Remaining tasks: %v (len: %d, cap: %d)\n", remainingTasks, len(remainingTasks), cap(remainingTasks))

    // 6. Demonstrate the reference-type behavior
    fmt.Println("\n--- Slice as Reference Type ---")
    originalList := []string{"a", "b", "c"}
    copiedList := originalList // copiedList now points to the same underlying array

    copiedList[0] = "x" // Modify through copiedList
    fmt.Printf("originalList: %v\n", originalList) // originalList is also changed!
    fmt.Printf("copiedList: %v\n", copiedList)

    // 7. Appending to a slice that shares an underlying array
    sharedList := []string{"p", "q", "r"}
    subList1 := sharedList[:2] // ["p", "q"]
    subList2 := sharedList[1:] // ["q", "r"]

    subList1 = append(subList1, "s") // ["p", "q", "s"]
    // If there's enough capacity in the underlying array,
    // this can affect subList2 and sharedList

    fmt.Printf("sharedList after append to subList1: %v\n", sharedList) // May be ["p", "q", "s"] or ["p", "q", "r"]
    fmt.Printf("subList1: %v\n", subList1)                             // ["p", "q", "s"]
    fmt.Printf("subList2: %v\n", subList2)                             // May be ["q", "s"] or ["q", "r"]

    subList1 = append(subList1, "t", "u") // Appending beyond capacity *will* allocate a new underlying array
    fmt.Printf("\nsharedList after append beyond capacity to subList1: %v\n", sharedList) // Will be ["p", "q", "r"]
    fmt.Printf("subList1: %v\n", subList1) // ["p", "q", "s", "t", "u"]
    fmt.Printf("subList2: %v\n", subList2) // ["q", "r"]

}