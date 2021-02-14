package main

import (
  "fmt"
)

func main() {
  fruits := map[string]string{"a": "apple", "b": "banana"}
  for key, value := range fruits {
    fmt.Printf("%s is for %s.\n", key, value)
  }
}
