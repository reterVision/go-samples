package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"

  "observer"
  "universe"
)


func main() {
  fmt.Println(`
    Use following commands to observe the universe

    stars create -- Create a new universe
    stars add <star numbers> -- Add more stars to your universe
    stars reduce <star numbers> -- Reduce stars in your universe
    stars offset <x> <y> -- Move your universe to somewhere else

    stars observe <tool> -- Observe your universe by either Hubble or Toy
  `)
}
