package main

import (
    "fmt"
)

func main() {
    var name string
    var age int

    // Ask for user's name
    fmt.Println("What is your name?")
    fmt.Scan(&name)

    // Ask for user's age
    fmt.Println("What is your age?")
    fmt.Scan(&age)

    // Create a formatted message using Sprintf and store it in newMessage
    newMessage := fmt.Sprintf("Nice to meet you, %s! You are %d years old.", name, age)

    // Print the message
    fmt.Printf(newMessage)
}
