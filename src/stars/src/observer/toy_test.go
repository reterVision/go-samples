package observer

import "testing"

func TestToy(t *testing.T) {
  toy := &ToyObserver{}
  toy.name = "Optimus"
  toy.diameter = 2.4

  toy.Observe("asgard")
}
