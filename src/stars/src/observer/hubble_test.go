package observer

import "testing"

func TestHubble(t *testing.T) {
  hubble := &HubbleObserver{}
  hubble.name = "H264"
  hubble.diameter = 1.2

  hubble.Observe()
}
