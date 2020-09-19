package main

import (
  "fmt"
)

func main() {
  fmt.Print("What is your name? ")
  var response string
  fmt.Scan(&response)
  fmt.Printf("Hello, %v!\n", response)
}
