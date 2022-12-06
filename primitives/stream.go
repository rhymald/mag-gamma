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
  buffer := make(map[string]float64)
  buffer["Power"] = math.Log2(1+stream.Destruction)
  buffer["Precision"] = math.Sqrt(stream.Alteration/Vector(stream.Creation,stream.Destruction,stream.Alteration))
  buffer["Quickness"] = 1000/Vector(math.Log2(2+stream.Destruction),math.Log2(1+stream.Alteration))
  buffer["Fuel"] = math.Log2(1024*Vector(stream.Destruction,stream.Creation))
  return buffer
}
