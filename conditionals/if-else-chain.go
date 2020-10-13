package main

import (
  "os"
  "bufio"
  "fmt"
  "strings"
)

func main() {
  // Ask the user to enter a number.
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("What position did you finish at? ")
  positionInput, _ := reader.ReadString('\n')
  position := strings.TrimSuffix(positionInput, "\n")

  // Print out message for the position.
  if position == "1" {
    fmt.Println("Congrats! You finished in first place.")
  } else if position == "2" {
    fmt.Println("You were a close second.")
  } else if position == "3" {
    fmt.Println("You finished in third.")
  } else {
    fmt.Println("Better luck next time!")
  }
}
