package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func main() {
  maxPurchase := 1000.0
  quantity := readNumber("Quantity")
  unitPrice := readNumber("Unit price")
  taxes := readNumber("Taxes")
  if total := quantity * unitPrice * (1 + taxes); total > maxPurchase {
    fmt.Printf("Purchase denied: $%.2f is above $%.2f.\n", total, maxPurchase)
  } else {
    fmt.Printf("Purchase approved: $%.2f\n", total)
  }
}

func readNumber(label string) float64 {
  fmt.Print(label + "? ")
  reader := bufio.NewReader(os.Stdin)
  rawInput, _ := reader.ReadString('\n')
  input := strings.TrimSuffix(rawInput, "\n")
  if result, err := strconv.ParseFloat(input, 64); err == nil {
    return result
  }
  panic("Invalid input.")
}
