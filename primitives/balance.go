package primitives
import "math"

//// Player stats

// streams basics
func BendHeraldFromClass(class float64) (float64, float64) { return math.Cbrt(7.5-class), math.Cbrt(0.5+class) }
func LenFromStream(stream Stream) float64 { return stream.Creation*1024 }
func WidFromStream(stream Stream) float64 { return 32*stream.Alteration/Vector(stream.Alteration, stream.Creation) }
func PowFromStream(stream Stream) float64 { return 10*stream.Destruction/Vector(stream.Destruction, stream.Alteration, stream.Creation) }

// elemental state
func ResistanceFromState(state Stream) float64 { return Vector(state.Destruction,state.Destruction,state.Creation) }

// pool
func MaxVolFromStreams(streams []Stream) float64 { buffer := 0.0 ; for _, each := range streams { buffer += math.Sqrt(LenFromStream(each)) } ; return buffer }

// dots
func RegenerateFullTimeOut() float64 { return 4000 }
func RegenerationPortionFromPool(max float64, current int) int { return int( math.Sqrt(max-float64(current)) ) }
func DotWeightAndTimeoutForRegenerationFromStreamAndMaxVol(stream Stream, maxvol float64) (float64, float64, float64) {
  weight  := Log1479( stream.Alteration ) * (1 + RNF()) / 2
  timeout := 1000*Log1479(maxvol)/math.Sqrt(maxvol)*math.Sqrt(weight)
  health  := 0.0
  return weight, timeout, health
}
