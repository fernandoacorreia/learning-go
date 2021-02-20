package main

import (
  "encoding/binary"
  "fmt"
  crypto_rand "crypto/rand"
  math_rand "math/rand"
)

func main() {
  seedRandomizer()
  for i := 0; i < 3; i++ {
    r := math_rand.Int()
    fmt.Println(r)
  }
}

func seedRandomizer() {
    var b [8]byte
    _, err := crypto_rand.Read(b[:])
    if err != nil {
        panic("cannot seed math/rand package with cryptographically secure random number generator")
    }
    math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}
