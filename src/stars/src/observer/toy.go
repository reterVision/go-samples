package observer

import "fmt"

type ToyObserver struct {
  name string
  diameter float32
}

func (t *ToyObserver)Observe(name string) {
  fmt.Printf("Observing the universe %s by Toy\n", name)
  fmt.Println("Type:", t.name)
  fmt.Println("Observation diameter:", t.diameter)
}
