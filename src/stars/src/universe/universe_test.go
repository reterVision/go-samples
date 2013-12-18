package universe

import "testing"

func TestOps(t *testing.T) {
  uu := NewUniverse("asgard", 100, 1.0, 1.0)
  if uu == nil {
    t.Error("NewUniverse failed.")
  }

  uu.name = "Asgard"
  uu.star = 100
  uu.latitude = 1.0
  uu.longtitude = 1.0

  uu.Add(100)
  if uu.star != 200 {
    t.Error("Universe.Add() failed.")
  }

  uu.Reduce(100)
  if uu.star != 100 {
    t.Error("Universe.Reduce() failed.")
  }

  uu.Offset(1.0, 1.0)
  if uu.latitude != 2.0 || uu.longtitude != 2.0 {
    t.Error("Universe.Offset() failed.")
  }
}
