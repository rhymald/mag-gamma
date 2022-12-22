package primitives

import "math"

type Stream struct {
  Creation float64
  Alteration float64
  Destruction float64
  Element string
  Heat struct {
    Threshold float64
    Current float64
  }
}

func StreamSum(element string, list []Stream) Stream {
  buffer := Stream{Element: element}
  dummy := Stream{Element: element}
  for _, each := range list {
    if each.Element == element || element == "Common" {
      buffer.Creation += each.Creation
      buffer.Alteration += each.Alteration
      buffer.Destruction += each.Destruction
    }
  }
  if buffer.Creation <= 0 {return dummy}
  return buffer
}

func StreamMean(element string, list []Stream) Stream {
  buffer, counter := Stream{Element: element}, 0
  dummy := Stream{Element: element}
  for _, each := range list {
    if each.Element == element || element == "Common" {
      buffer.Creation += 1 / each.Creation
      buffer.Alteration += 1 / each.Alteration
      buffer.Destruction += 1 / each.Destruction
      counter++
    }
  }
  if buffer.Creation <= 0 {return dummy}
  buffer.Creation = float64(counter) / buffer.Creation
  buffer.Alteration = float64(counter) / buffer.Alteration
  buffer.Destruction = float64(counter) / buffer.Destruction
  return buffer
}

func StatsFromStream(stream Stream) map[string]float64 {
  coefficiaent := 1.1479
  // lowthres := 1/math.Sqrt(coefficiaent)
  buffer := make(map[string]float64)
  // volume := Vector(stream.Creation+1,stream.Destruction+1,stream.Alteration+1)
  // Destruction = damage, power
  buffer["D-Power"] = Log1479(stream.Destruction)
  if StreamStructure2(stream.Destruction,stream.Creation,stream.Alteration,coefficiaent) {
    // Antibarrier = +AddDamage, +ticks, - if D>C close to each other
    buffer["Dc-Sharpening"] = StreamAffinity2(stream.Destruction,stream.Creation,coefficiaent)
  }
  if StreamStructure2(stream.Destruction,stream.Alteration,stream.Creation,coefficiaent) {
    // Permanent debuff (hard to clean) = +Speed, +effectiveness, - if D>A close to each other
    buffer["Da-Barrier"] = StreamAffinity2(stream.Destruction,stream.Alteration,coefficiaent)
  }
  if StreamStructure3(stream.Destruction,stream.Alteration,stream.Creation,coefficiaent) {
    // Pulsing damage = +efectiveness, +damage, +speed, - if D>(A=C) when ac close to each other
    buffer["Dac-Disaster"] = StreamAffinity3(stream.Destruction,stream.Alteration,stream.Creation,coefficiaent)
  }
  // Alteration = luck, dexterity
  buffer["A-Concentration"] = math.Sqrt(Log1479(stream.Alteration)+1)-1
  if StreamStructure2(stream.Alteration,stream.Destruction,stream.Creation,coefficiaent) {
    // Smooth damaging conditions (easy to clean) = +time, +damage : A>D
    buffer["Ad-Condition"] = StreamAffinity2(stream.Alteration,stream.Destruction,coefficiaent)
  }
  if StreamStructure2(stream.Alteration,stream.Creation,stream.Destruction,coefficiaent) {
    // Smooth buff (easy to rip-off) = +time, +edfectiveness : A>C
    buffer["Ac-Boon"] = StreamAffinity2(stream.Alteration,stream.Creation,coefficiaent)
  }
  if StreamStructure3(stream.Alteration,stream.Creation,stream.Destruction,coefficiaent) {
    // Permanent buff trigger = +effectiveness, +chance, +speed : A>(D=C)
    buffer["Adc-Transformation"] = StreamAffinity3(stream.Alteration,stream.Destruction,stream.Creation,coefficiaent)
  }
  // Creation = give, intelligence
  buffer["C-Toughness"] = math.Sqrt(1+stream.Creation)-1
  if StreamStructure2(stream.Creation,stream.Destruction,stream.Creation,coefficiaent) {
    // Shield = +amount, +time : C>D
    buffer["Cd-Decay"] = StreamAffinity2(stream.Creation,stream.Destruction,coefficiaent)
  }
  if StreamStructure2(stream.Creation,stream.Alteration,stream.Destruction,coefficiaent) {
    // Heal recovery = +efectiveness, +speed : C>A
    buffer["Ca-Restoration"] = StreamAffinity2(stream.Creation,stream.Alteration,coefficiaent)
  }
  if StreamStructure3(stream.Creation,stream.Destruction,stream.Alteration,coefficiaent) {
    // Conjuration local shadows, wells = +volume, +activity, +efectiveness : C>(A=D)
    buffer["Cad-Summon"] = StreamAffinity3(stream.Creation,stream.Alteration,stream.Destruction,coefficiaent)
  }
  // Main meta
  buffer["M-Quickness"] = 1000/Vector(math.Log2(2+stream.Destruction),math.Log2(1+stream.Alteration))
  buffer["M-Fuel"] = math.Log2(1024*Vector(stream.Destruction,stream.Creation))
  // buffer["M-Heat"]   = math.Log2(1024*Vector(stream.Destruction,stream.Creation))
  return buffer
}

func StreamStructure2(a float64, b float64, c float64, t float64) bool { if a > b && b*math.Sqrt(t) > c && a/b > 1 && a/b < t { return true } ; return false }
func StreamStructure3(a float64, b float64, c float64, t float64) bool { if ( StreamStructure2(a,b,c,t) || StreamStructure2(a,c,b,t) ) && math.Max(math.Max(a/b,a/c),b/c)<math.Cbrt(t)*math.Cbrt(t) && math.Max(b/c,c/b) < math.Sqrt(t) { return true } ; return false }
func StreamAffinity2(a float64, b float64, t float64) float64 { return math.Pow(math.Log2(t/(a/b))/math.Log2(t), 2) }
func StreamAffinity3(a float64, b float64, c float64, t float64) float64 {
  ab, ca := math.Max(a,b)/math.Min(a,b), math.Max(a,c)/math.Min(a,c)
  return math.Pow(math.Log2(t/(2/(1/ab+1/ca)))/math.Log2(t), 2)
}
