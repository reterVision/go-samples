package universe

type Universe struct {
  name string
  star int
  latitude float32
  longtitude float32
}

func NewUniverse() *Universe {
  return &Universe{"Asgard", 100, 1.0, 1.0}
}

func (u *Universe) Add(star int) *Universe {
  u.star += star
  return u
}

func (u *Universe) Reduce(star int) *Universe {
  u.star -= star
  return u
}

func (u *Universe) Offset(x float32, y float32) *Universe {
  u.latitude += x
  u.longtitude += y
  return u
}
