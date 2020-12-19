package main

import (
  "fmt"
)

func main() {
  x := 0
  y := 0
  endCommand := "end"
  for command := ""; command != endCommand; {
    fmt.Print("Command? ")
    fmt.Scan(&command)
    switch command {
    case "up":
      y -= 1
    case "down":
      y += 1
    case "left":
      x -= 1
    case "right":
      x += 1
    case endCommand:
      fmt.Println("Exiting.")
    default:
      fmt.Printf("Invalid command: %v\n", command)
    }
    fmt.Printf("x=%v, y=%v\n", x, y)
  }
}
