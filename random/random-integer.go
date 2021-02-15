package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  for i := 0; i < 3; i++ {
    r := rand.Int()
    fmt.Println(r)
  }
}
