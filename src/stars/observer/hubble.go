package observer

import "fmt"

type Observer interface {
  Observe(source string)
}
