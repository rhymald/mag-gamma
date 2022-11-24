package player
import "rhymald/mag-gamma/primitives"

type Flock struct {
  Streams []int
  HeatInfo []float64
  // C, A, D, Ca, Cd, Ac, Ad, Dc, Da
}

func ComposeFlock(streams []primitives.Stream, estate ElementalState) Flock {
  buffer := Flock{}
  for i:=0; i<len(streams); i++ { buffer.Streams = append(buffer.Streams, i) }
  // for _, index := range buffer.Streams { index = 0 }
  return buffer
}
