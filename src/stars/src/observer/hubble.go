package observer

import "fmt"

type HubbleObserver struct {
  name string
  diameter float32
}

func (h *HubbleObserver)Observe() {
  fmt.Println("Observing the universe by Hubble")
  fmt.Println("Type:", h.name)
  fmt.Println("Observation diameter:", h.diameter)
}
