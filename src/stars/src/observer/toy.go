package observer

import "fmt"

type ToyObserver struct {
  name string
  diameter float32
}

func (t *ToyObserver)Observe() {
  fmt.Println("Observing the universe by Toy")
  fmt.Println("Type:", t.name)
  fmt.Println("Observation diameter:", t.diameter)
}
