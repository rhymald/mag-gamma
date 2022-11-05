package primitives

type Stream struct {
  Creation float64
  Alteration float64
  Destruction float64
  Element string
  HeatPrint float64
  LWP [3]float64
}

func StreamSum(element string, list []Stream) Stream {
  buffer := Stream{}
  dummy := Stream{Element: element}
  for _, each := range list {
    if each.Element == element || element == "Common" {
      buffer.Creation += each.Creation
      buffer.Alteration += each.Alteration
      buffer.Destruction += each.Destruction
    }
  }
  buffer.Element = element
  if buffer.Creation <= 0 {return dummy}
  return buffer
}
