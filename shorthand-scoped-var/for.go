package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func main() {
  iterations := readNumber("Number of iterations")
  sum := 0
  for i := 1; i <= iterations; i++ {
    sum += i
  }
  fmt.Println(sum)
}

func readNumber(label string) int {
  fmt.Print(label + "? ")
  reader := bufio.NewReader(os.Stdin)
  rawInput, _ := reader.ReadString('\n')
  input := strings.TrimSuffix(rawInput, "\n")
  if result, err := strconv.Atoi(input); err == nil {
    return result
  }
  panic("Invalid input.")
}
