package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "strconv"

  "observer"
  "universe"
)

var asgard = universe.NewUniverse("asgard", 100, 1.0, 1.0)

func handleStarsCommands(tokens []string) {
  switch tokens[1] {
  case "list":
    fmt.Printf("Now you have a universe called %s\n", asgard.Name())
    x, y := asgard.Coordinates()
    fmt.Printf("It is in the position of (%f, %f)\n", x, y)
    fmt.Printf("And it has %d stars!\n", asgard.Star())

  case "add":
    if len(tokens) != 3 {
      fmt.Println("Not enough parameters!")
      break
    }
    increase_count, err := strconv.ParseInt(tokens[2], 10, 32)
    if err != nil {
      fmt.Println("Increase star count error", err)
      break
    }
    asgard.Add(int(increase_count))
    fmt.Printf("And it has %d stars!\n", asgard.Star())

  case "reduce":
    if len(tokens) != 3 {
      fmt.Println("Not enough parameters!")
      break
    }
    reduce_count, err := strconv.ParseInt(tokens[2], 10, 32)
    if err != nil {
      fmt.Println("Reduce star count error", err)
      break
    }
    asgard.Reduce(int(reduce_count))
    fmt.Printf("And it has %d stars!\n", asgard.Star())

  case "offset":
    if len(tokens) == 4 { 
      x, err := strconv.ParseFloat(tokens[2], 32)
      if err != nil {
        fmt.Println("Offset x error", err)
        break
      }
      y, err := strconv.ParseFloat(tokens[3], 32)
      if err != nil {
        fmt.Println("Offset y error", err)
        break
      }
      asgard.Offset(float32(x), float32(y))
      m, n := asgard.Coordinates()
      fmt.Printf("It is in the position of (%f, %f)\n", m, n)
    } else {
      fmt.Println("You need to give me the x and y Offset")
    }

  case "observe":
    if len(tokens) == 3 {
      observer.Observe(tokens[2], asgard.Name())
    } else {
      fmt.Println("You need to tell me which tool you want to use to observe the universe")
    }

  default:
    fmt.Println("Unrecognized stars command:", tokens[1])
  }
}


func main() {
  fmt.Println(`
    Use following commands to observe the universe

    stars list -- List the universe and stars you have
    stars add <star numbers> -- Add more stars to your universe
    stars reduce <star numbers> -- Reduce stars in your universe
    stars offset <x> <y> -- Move your universe to somewhere else

    stars observe <tool> -- Observe your universe by either Hubble or Toy

    Use "q" or "e" to quit.
  `)

  r := bufio.NewReader(os.Stdin)

  for {
    fmt.Print("Enter command -> ")

    rawLine, _, _ := r.ReadLine()
    line := string(rawLine)

    if line == "q" || line == "e" {
      break
    }

    tokens := strings.Split(line, " ")

    if tokens[0] == "stars" {
      handleStarsCommands(tokens)
    } else {
      fmt.Println("Unrecognized command:", tokens[0])
    }
  }
}
