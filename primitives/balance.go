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
func CrackStream_DotWeightFromStream(stream Stream) float64 { return Log1479( stream.Destruction ) * (RNF() + RNF()) / 2 }
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

// heat
func GenerateHeat_FromStreamAndDot(stream Stream, dot Dot) float64 { return Vector(Log1479(stream.Destruction),Log1479(dot.Weight)) }
func GenerateHeat_ComposeHeat(heat [9]float64) [9]float64 {
  resume := [9]float64{}
  for i, h := range heat {
    if i == 0 {
      resume[i] = 0
    } else if i == 5 {
      resume[1] += - math.Sqrt(h)
      resume[2] += - math.Sqrt(h)
      resume[3] += - math.Sqrt(h)
      resume[4] += - math.Sqrt(h)
      resume[6] += - math.Sqrt(h)
      resume[7] += - math.Sqrt(h)
      resume[8] += - h
      resume[5] += h
    } else if i == 8 {
      resume[i] += resume[i] + h
      for j:=0; j<8; j++ { resume[j] *= (RNF() + RNF()) }
    } else {
      resume[i] += h
    }
  }
  return resume
}
func GenerateHeat_CompareHeat(heat [9]float64, flock []Stream) [9]float64 {
  limits := [9]float64{}
  // // cres, alts, dess := 0.0, 0.0, 0.0
  // for _, each := range flock {
  //   // alts += each.Alteration
  //   // dess += each.Destruction
  //   // cres += each.Creation
  //   limits[ElemToInt(each.Element)] += 1
  // }
  // limits[0] = Vector(alts+1,dess+1,dess+1,cres+1)
  // for i:=1; i<9; i++ { if limits[i]!=0 { limits[i] =  heat[i] /limits[i] / limits[0] }}
  return limits
}
