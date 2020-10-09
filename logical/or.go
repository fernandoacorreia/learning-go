package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main() {
  reader := bufio.NewReader(os.Stdin)

  fmt.Print("What day of the week is it? ")
  dowInput, _ := reader.ReadString('\n')
  dow := strings.TrimSuffix(dowInput, "\n")

  if dow == "Saturday" || dow == "Sunday" {
    fmt.Println("Stay home.")
  } else {
    fmt.Println("Go to work.")
  }
}
