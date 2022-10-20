package environment

import "rhymald/mag-gamma/primitives"

type Location struct {
  Name string
  Description string
  Wells []PowerState
  Curses []PowerState
}

type PowerState struct {
  Description string
  Nature []primitives.Stream
  Area float64
  XYZs [][3]float64
  Concentrated bool
}

// +spawn

var (
  Environment Location
)
