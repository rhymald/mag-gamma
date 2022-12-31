package player
import "rhymald/mag-gamma/primitives"

func DefaultFlockFromAllStreams(streams []primitives.Stream) []int {
  buffer := []int{}
  for i:=0; i<len(streams); i++ { buffer = append(buffer, i) }
  return buffer
}
