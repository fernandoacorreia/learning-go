package main

import (
  "fmt"
)

func main() {
  fmt.Print("What is your first name? ")
  var response string
  fmt.Scanln(&response)
  fmt.Printf("Hello, %v!\n", response)
}
