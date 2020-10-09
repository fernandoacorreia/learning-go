package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main() {
  reader := bufio.NewReader(os.Stdin)

  fmt.Print("Are you hungry (y/n)? ")
  hungryInput, _ := reader.ReadString('\n')
  hungryCleaned := strings.TrimSuffix(hungryInput, "\n")
  hungry := hungryCleaned == "y"
  notHungry := !hungry

  fmt.Println("You are hungry: ", hungry)
  fmt.Println("You are satisfied: ", notHungry)
}
