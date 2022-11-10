package primitives
import "math"

//// Player stats

// streams basics
func NewBornStreams_BendHeraldFromClass(class float64) (float64, float64) { return math.Cbrt(7.5-class), math.Cbrt(0.5+class) }
func NewBornStreams_LenFromStream(stream Stream) float64 { return stream.Creation*1024 }
func NewBornStreams_WidFromStream(stream Stream) float64 { return 32*stream.Alteration/Vector(stream.Alteration, stream.Creation) }
func NewBornStreams_PowFromStream(stream Stream) float64 { return 10*stream.Destruction/Vector(stream.Destruction, stream.Alteration, stream.Creation) }

// elemental state
func InnerAffinization_ResistanceFromState(state Stream) float64 { return Vector(state.Destruction,state.Destruction,state.Creation) }

// pool
func ExtendPool_MaxVolFromStreams(streams []Stream) float64 { buffer := 0.0 ; for _, each := range streams { buffer += math.Sqrt(NewBornStreams_LenFromStream(each)) } ; return buffer }

// dots
func Pool_RegenerateFullTimeOut() float64 { return 4000 }
func CrackStream_DotWeightFromState(state Stream) float64 { return Log1479( state.Destruction ) * (RNF() + RNF()) / 2 }
func RegenerateDots_PortionFromPool(max float64, current int) int { return int( math.Sqrt(max-float64(current)) ) }
func EmitDot_DotWeightAndTimeoutFromStreamAndMaxVol(stream Stream, maxvol float64) (float64, float64, float64) {
  weight  := Log1479( stream.Alteration ) * (1 + RNF()) / 2
  timeout := 1000*Log1479(maxvol)/math.Sqrt(maxvol)*math.Sqrt(weight)
  health  := 0.0
  return weight, timeout, health
}
func DotTransferIn_DotWeightAndTimeoutFromState(state Stream) (float64, float64) {
  weight := Log1479( state.Alteration ) * (1 + RNF()) / 2
  step   := 32*math.Sqrt(1+math.Abs(state.Creation))
  pause  := 1000*math.Log2(step)/math.Sqrt(step)*math.Sqrt(weight)
  return weight, pause
}
func DotTransferOut_TimeoutFromWeightAndState(weight float64, state Stream) (float64) {
  step   := 32*math.Sqrt(1+math.Abs(state.Destruction))
  pause  := 1000*math.Log2(step)/math.Sqrt(step)*math.Sqrt(weight)
  return pause
}
// func Transference_TotalCooldownFromDemand(demand [9]int) float64 {
//   sum := 0.0
//   for _, i := range demand { sum += math.Abs(float64(i))*500 }
//   return sum
// }
func Transference_DotCountDemandAndTotalCooldownFromStates(estate [9]Stream) (float64, [9]int) {
  demand := [9]int{}
  count := 0.0
  cooldown := 0.0
  for i, source := range estate {
    // count := primitives.Transference_DotCountFromState(source)
    if source.Creation < 0 { count = math.Sqrt(1+math.Abs(source.Destruction)) * (1 + RNF()) / -2 }
    if source.Creation > 0 { count = math.Sqrt(1+math.Abs(source.Creation)) * (1 + RNF()) / 2 }
    if i == 0 { count = 0 }
    demand[i] = ChancedRound(count * Sign(source.Creation))
    cooldown = math.Max(math.Abs(count) * 500, cooldown)
  }
  if cooldown == 0 { cooldown = Pool_RegenerateFullTimeOut() }
  // cooldown := primitives.Transference_TotalCooldownFromDemand(demand)
  return cooldown, demand
}
