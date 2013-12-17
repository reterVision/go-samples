package observer

import "fmt"

type Observer interface {
  Observe()
}

func Observe(otype string) {
  var o Observer

  switch otype {
  case "Hubble":
    o = &HubbleObserver{}
  case "Toy":
    o = &ToyObserver{}
  default:
    fmt.Println("Unsupported oberserver type", otype)
    return
  }

  o.Observe()
}
