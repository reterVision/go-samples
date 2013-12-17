package observer

import "fmt"

type Observer interface {
  Observe(source string)
}

func Observe(source, otype string) {
  var o Observer

  switch otype {
  case "Hubble":
    o = &HubbleObserver()
  case "Toy":
    o = &ToyObeserver()
  default:
    fmt.Println("Unsupported oberserver type", otype)
    return
  }

  o.Observe(source)
}
