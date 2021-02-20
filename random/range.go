package main

import (
  "encoding/binary"
  "fmt"
  crypto_rand "crypto/rand"
  math_rand "math/rand"
)

func main() {
  seedRandomizer()
  for i := 0; i < 30; i++ {
    fmt.Printf("%v ", randomInt(1, 10))
  }
  fmt.Printf("\n")
}

func randomInt(min, max int) int {
  return min + math_rand.Intn(max - min + 1)
}

// Based on John Leidegren's code at https://stackoverflow.com/a/54491783/376366
func seedRandomizer() {
    var b [8]byte
    _, err := crypto_rand.Read(b[:])
    if err != nil {
        panic("cannot seed math/rand package with cryptographically secure random number generator")
    }
    math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}
