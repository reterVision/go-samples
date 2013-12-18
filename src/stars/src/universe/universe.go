package universe

type Universe struct {
  name string
  star int
  latitude float32
  longtitude float32
}

func NewUniverse(name string, star int, latitude float32, longtitude float32) *Universe {
  return &Universe{name, star, latitude, longtitude}
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

func (u *Universe) Name() string {
  return u.name
}

func (u* Universe) Coordinates() (float32, float32) {
  return u.latitude, u.longtitude
}

func (u* Universe) Star() int {
  return u.star
}
