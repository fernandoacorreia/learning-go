package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main() {
  fmt.Print("Which direction? ")
  switch direction := readString(); direction {
  case "north":
    fmt.Printf("Far up %v you see a mountain range.\n", direction)
  case "south":
    fmt.Printf("Looking %v you see sand all the way to the horizon.\n", direction)
  case "east":
    fmt.Printf("A short distance %v you see a house by a tree.\n", direction)
  case "west":
    fmt.Printf("Looking %v you see the shore.\n", direction)
  default:
    fmt.Printf("%v is not a valid direction.\n", direction)
  }
}

func readString() string {
  reader := bufio.NewReader(os.Stdin)
  rawInput, _ := reader.ReadString('\n')
  input := strings.TrimSuffix(rawInput, "\n")
  return input
}
