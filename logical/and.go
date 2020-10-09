package main

import (
  "os"
  "bufio"
  "fmt"
  "strings"
)

func main() {
  reader := bufio.NewReader(os.Stdin)

  fmt.Print("Are the lights on (y/n)? ")
  lightsOnInput, _ := reader.ReadString('\n')
  lightsOnCleaned := strings.TrimSuffix(lightsOnInput, "\n")
  lightsOn := lightsOnCleaned == "y"

  fmt.Print("Is the door open (y/n)? ")
  doorOpenInput, _ := reader.ReadString('\n')
  doorOpenCleaned := strings.TrimSuffix(doorOpenInput, "\n")
  doorOpen := doorOpenCleaned == "y"

  if lightsOn && doorOpen {
    fmt.Println("Go in.")
  } else {
    fmt.Println("Come back later.")
  }
}
