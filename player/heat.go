package player

import "rhymald/mag-gamma/primitives"
import "fmt"
// import "math"

type Heat struct {
  Current [9]float64
  Compared [9]float64
}

func ConsumeHeat(heatState *Heat, heat [9]float64) {
  fmt.Printf("\n ◦◦◦◦◦ DEBUG [Consuming heat][Incoming heat]:          ") ; for i, h:=range heat { if h!=0 { fmt.Printf(" %+1.0f'%s ", h, ElemSigns[i]) } }
  bufferHeat := primitives.CollectHeat(*&heatState.Current, heat)
  mean, counter, sum := 0.0, 0, 0.0
  bufferOverheat := [9]float64{}
  for i, _ := range bufferHeat {
    if bufferHeat[i] <= 0 {
      bufferHeat[i] = 0
    } else {
      sum += bufferHeat[i]
      mean += 1/bufferHeat[i]
      counter++
    }
  }
  fmt.Printf("\n ◦◦◦◦◦ DEBUG [Consuming heat][Current heat rates]:     "); for i, h:=range bufferHeat { if h!=0 { fmt.Printf(" %1.0f'%s ", h, ElemSigns[i]) } }
  mean = float64(counter) / mean
  bufferOverheat = bufferHeat
  bufferOverheat[0] = mean
  fmt.Printf("\n ◦◦◦◦◦ DEBUG [Consuming heat][Overheat calculatings]: %s:%1.0f ", ElemSigns[0], mean); for i, h:=range bufferOverheat { if h!=0 { fmt.Printf(" %1.0f'%s ", h, ElemSigns[i]) } }
  for i, oh := range bufferOverheat { bufferOverheat[i] = oh/mean }
  fmt.Printf("\n ◦◦◦◦◦ DEBUG [Consuming heat][Overheat comparsion]:   %s:%1.0f ", ElemSigns[0], mean); for i, h:=range bufferOverheat { if h>0 && i!=0 { fmt.Printf(" %1.2f'%s ", h, ElemSigns[i]) } }
  *&heatState.Current = bufferHeat
  *&heatState.Compared = bufferOverheat
}

func PlotHeatState(heat Heat) {
  fmt.Printf("\n ┌──── INFO [heat state]:")
  for i:=0 ; i<9; i++ {
    if i==0 {fmt.Printf("\n │ Average%s: %1.2f", ElemSigns[0], heat.Compared[0])}
    if i!=0 && heat.Compared[i]>0 {fmt.Printf("\n │ ")}
  }
  fmt.Println()
  fmt.Printf(" └────────────────────────────────────────────────────────────────────────────────────────────────────")
}
