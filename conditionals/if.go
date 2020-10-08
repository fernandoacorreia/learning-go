package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  // Initialize random number generator.
  r := rand.New(rand.NewSource(time.Now().UnixNano()))

  // Generate 2 small integer numbers.
  mine := r.Intn(10)
  yours := r.Intn(10)

  // Print out the random numbers.
  fmt.Println("Mine:", mine)
  fmt.Println("Yours:", yours)

  // Print out whether I won.
  if mine > yours {
    fmt.Println("I win.")
  }
}
