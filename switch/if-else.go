package main

import (
  "fmt"
)

func main() {
  x := 0
  y := 0
  for command := ""; command != "end"; {
    fmt.Print("Command? ")
    fmt.Scan(&command)
    if command == "up" {
      y -= 1
    } else if command == "down" {
      y += 1
    } else if command == "left" {
      x -= 1
    } else if command == "right" {
      x += 1
    } else if command == "end" {
      fmt.Println("Exiting.")
    } else {
      fmt.Printf("Invalid command: %v\n", command)
    }
    fmt.Printf("x=%v, y=%v\n", x, y)
  }
}
