package primitives
import "math"

//// Player stats

// streams basics
func NewBornStreams_BendHeraldFromClass(class float64) (float64, float64) { return math.Sqrt(7.5-class), math.Sqrt(0.5+class) }

// elemental state
func InnerAffinization_ResistanceFromState(state Stream) float64 { return Vector(state.Destruction,state.Destruction,state.Creation) }

// pool
func ExtendPool_MaxVolFromStreams(streams []Stream) float64 { buffer := 0.0 ; for _, each := range streams { buffer += math.Sqrt(1024*each.Creation) } ; return buffer }

// dots
func Pool_RegenerateFullTimeOut() float64 { return 4096 }
func CrackStream_DotWeightFromStream(stream Stream) float64 { return math.Log2(Vector(stream.Destruction+1,stream.Creation,stream.Alteration)/3.5+1) * (1 + RNF()) / 2 }
func RegenerateDots_PortionFromPool(max float64, current int) int { return int( math.Sqrt(max-float64(current)) ) }
func EmitDot_DotWeightAndTimeoutFromStreamAndMaxVol(stream Stream, maxvol float64) (float64, float64, float64) {
  weight  := math.Log2(Vector(stream.Destruction,stream.Creation+1,stream.Alteration)/3.5+1) * (1 + RNF()) / 2
  timeout := 1024*math.Log2(maxvol)/math.Sqrt(maxvol)*math.Sqrt(weight)
  health  := 0.0
  return weight, timeout, health
}
func DotTransferIn_DotWeightAndTimeoutFromState(state Stream) (float64, float64) {
  weight := math.Log2(Vector(state.Destruction,state.Creation,state.Alteration+1)/3.5+1) * (1 + RNF()) / 2
  step   := 32*math.Sqrt(math.Abs(state.Creation))
  pause  := 1024*math.Log2(step)/math.Sqrt(step)*math.Sqrt(weight)
  return weight, pause
}
func DotTransferOut_TimeoutFromWeightAndState(weight float64, state Stream) (float64) {
  step   := 32*math.Sqrt(math.Abs(state.Destruction))
  pause  := 1024*math.Log2(step)/math.Sqrt(step)*math.Sqrt(weight)
  return pause
}
func Transference_DotCountDemandAndTotalCooldownFromStates(estate [9]Stream, istate [9]Stream) (float64, [9]int) {
  demand := [9]int{}
  count := 0.0
  cooldown := 0.0
  counter := 0
  for i, source := range estate {
    if source.Creation < 0 { count = math.Sqrt(math.Sqrt(1024*math.Abs(source.Destruction)))*Sign(istate[i].Creation) ; counter++ }
    if source.Creation > 0 { count = math.Sqrt(math.Sqrt(1024*math.Abs(source.Creation)))*Sign(istate[i].Creation) ; counter++ }
    if i == 0 { count = 0 }
    demand[i] = ChancedRound(count * Sign(source.Creation))
    cooldown += math.Abs(count) * 2048
  }
  if counter != 0 { cooldown = cooldown / float64(counter) }
  if cooldown == 0 { cooldown = Pool_RegenerateFullTimeOut() }
  return cooldown, demand
}

// heat
func GenerateHeat_FromStreamAndDot(stream Stream, dot Dot) float64 {
  heat := 1 + (1+dot.Weight) * math.Log2(2+stream.Destruction*stream.Alteration/stream.Creation)
  if stream.Element == "Common" { return heat*math.Log2(2+heat) } // does not care dot element
  if dot.Element == "Common" { heat = heat/math.Phi }
  // + poisoned prey eats pred -- less heat
  // + same -- less heat
  // TBD
  return heat
}
func NewBorn_HeatThresholdFromStream(stream Stream, sum Stream) float64 {
  limits := Vector(1+sum.Creation,sum.Destruction,1+sum.Alteration)
  limits = Vector( Vector(1+stream.Creation,stream.Destruction,1+stream.Alteration), limits*math.Log10(10-float64(ElemToInt(stream.Element))) )
  return limits
}
func GenerateHeat_EffectFromCurrentAndThresholds(current float64, threshold float64) (float64, float64) {
  stabilityChance := math.Sqrt(1/(current/threshold+1-math.Sqrt2/2))
  selfDestruction := 1-1/(current/threshold/math.Sqrt(math.Sqrt2))
  return stabilityChance, selfDestruction
}
