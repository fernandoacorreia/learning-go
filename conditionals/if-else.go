package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  // Initialize random number generator.
  r := rand.New(rand.NewSource(time.Now().UnixNano()))

  // Generate 2 integer numbers.
  limit := r.Intn(10)
  value := r.Intn(20)

  // Print out the random numbers.
  fmt.Println("Limit:", limit)
  fmt.Println("Value:", value)

  // Print out result of comparison.
  if value > limit {
    fmt.Println("Above limit.")
  } else {
    fmt.Println("Within limit.")
  }
}
