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
  coefficiaent := 1.05
  buffer := make(map[string]float64)
  volume := Vector(stream.Creation+1,stream.Destruction+1,stream.Alteration+1)
  // Destruction = damage, power
  buffer["D-Power-Damage"]   = stream.Destruction/volume
  if stream.Destruction/stream.Creation > 1 && stream.Destruction/stream.Creation < coefficiaent {
    // Antibarrier = +AddDamage, +ticks, - if D>C close to each other
    buffer["Dc-Sharpening"] = math.Sqrt(Vector(stream.Destruction+1,stream.Creation)/volume)
  }
  if stream.Destruction/stream.Alteration > 1 && stream.Destruction/stream.Alteration < coefficiaent {
    // Permanent debuff (hard to clean) = +Speed, +effectiveness, - if D>A close to each other
    buffer["Da-Barrier"]       = math.Sqrt(Vector(stream.Destruction+1,stream.Alteration)/volume)
  }
  if stream.Creation/stream.Alteration > 1/coefficiaent && stream.Creation/stream.Alteration < coefficiaent && stream.Destruction/math.Max(stream.Creation,stream.Alteration)>coefficiaent*coefficiaent {
    // Pulsing damage = +efectiveness, +damage, +speed, - if D>(A=C) when ac close to each other
    buffer["Dac-Disaster"]   = math.Cbrt(Vector(stream.Destruction+1,stream.Destruction)/volume)
  }
  // Alteration = luck, dexterity
  buffer["A-Concentration"]   = stream.Alteration/volume
  if stream.Alteration/stream.Destruction > 1 && stream.Alteration/stream.Destruction < coefficiaent {
    // Smooth damaging conditions (easy to clean) = +time, +damage : A>D
    buffer["Ad-Condition"]       = math.Sqrt(Vector(stream.Destruction,stream.Alteration+1)/volume)
  }
  if stream.Alteration/stream.Creation > 1 && stream.Alteration/stream.Creation < coefficiaent {
    // Smooth buff (easy to rip-off) = +time, +edfectiveness : A>C
    buffer["Ac-Boon"]            = math.Sqrt(Vector(stream.Creation,stream.Alteration+1)/volume)
  }
  if stream.Creation/stream.Destruction > 1/coefficiaent && stream.Creation/stream.Destruction < coefficiaent && stream.Alteration/math.Max(stream.Creation,stream.Destruction)>coefficiaent*coefficiaent {
    // Permanent buff trigger = +effectiveness, +chance, +speed : A>(D=C)
    buffer["Adc-Transformation"] = math.Cbrt(Vector(stream.Alteration+1,stream.Alteration)/volume)
  }
  // Creation = give, intelligence
  buffer["C-Creation"]       = stream.Creation/volume
  if stream.Creation/stream.Destruction > 1 && stream.Creation/stream.Destruction < coefficiaent {
    // Shield = +amount, +time : C>D
    buffer["Cd-Decay"]     = math.Sqrt(Vector(stream.Creation,stream.Alteration+1)/volume)
  }
  if stream.Creation/stream.Alteration > 1 && stream.Creation/stream.Alteration < coefficiaent {
    // Heal recovery = +efectiveness, +speed : C>A
    buffer["Ca-Restoration"] = math.Sqrt(Vector(stream.Creation,stream.Alteration+1)/volume)
  }
  if stream.Alteration/stream.Destruction > 1/coefficiaent && stream.Alteration/stream.Destruction < coefficiaent && stream.Creation/math.Max(stream.Alteration,stream.Destruction)>coefficiaent*coefficiaent {
    // Conjuration local shadows, wells = +volume, +activity, +efectiveness : C>(A=D)
    buffer["Cad-Summon"]     = math.Cbrt(Vector(stream.Creation+1,stream.Creation)/volume)
  }
  // Main meta
  buffer["M-Quickness"] = 1000/Vector(math.Log2(2+stream.Destruction),math.Log2(1+stream.Alteration))
  buffer["M-Fuel"]      = math.Log2(1024*Vector(stream.Destruction,stream.Creation))
  // buffer["M-Heat"]   = math.Log2(1024*Vector(stream.Destruction,stream.Creation))
  return buffer
}
