package observer

import "fmt"

type HubbleObserver struct {
  name string
  diameter float32
}

func (h *HubbleObserver)Observe(name string) {
  fmt.Printf("Observing the universe %s by Hubble\n", name)
  fmt.Println("Type:", h.name)
  fmt.Println("Observation diameter:", h.diameter)
}
