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
    fmt.Println(randomString(20))
  }
}

func randomInt(min, max int) int {
  return min + math_rand.Intn(max - min + 1)
}

// Based on Flavio Copes's code at https://flaviocopes.com/go-random/
func randomString(length int) string {
  bytes := make([]byte, length)
  for i := 0; i < length; i++ {
    bytes[i] = byte(randomInt(97, 122))
  }
  return string(bytes)
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
