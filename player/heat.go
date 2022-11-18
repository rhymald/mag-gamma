package player

import "rhymald/mag-gamma/primitives"
import "fmt"
// import "math"

type Heat struct {
  Current [9]float64
  Compared [9]float64
  Unstable [9]float64
  Danger [9]float64
}

func ConsumeHeat(heatState *Heat, streams [9]primitives.Stream, heat [9]float64) {
  fmt.Printf("\n ◦◦◦◦◦ DEBUG [Consuming heat][Incoming heat]: ") ; for i, h:=range heat { if h!=0 { fmt.Printf(" %+1.0f'%s ", h, ElemSigns[i]) } }
  bufferHeat := primitives.CollectHeat(*&heatState.Current, heat)
  avg, counter := 0.0, 0
  bufferHeat[0] = 0
  // bufferOverheat := [9]float64{}
  for i, _ := range bufferHeat {
    if bufferHeat[i] <= 0 {
      bufferHeat[i] = 0
    } else {
      avg += 1 / bufferHeat[i]
      counter++
    }
  }
  fmt.Printf("\n ◦◦◦◦◦ DEBUG [Consuming heat][Current heat rates]: "); for i, h:=range bufferHeat { if h!=0 { fmt.Printf(" %1.0f'%s ", h, ElemSigns[i]) } }
  // mean = float64(counter) * float64(counter) / mean
  // bufferOverheat = bufferHeat
  bufferHeat[0] = float64(counter) / avg //* float64(counter)
  // fmt.Printf("\n ◦◦◦◦◦ DEBUG [Consuming heat][Overheat calculatings]: %s:%1.0f ", ElemSigns[0], mean); for i, h:=range bufferOverheat { if h!=0 { fmt.Printf(" %1.0f'%s ", h, ElemSigns[i]) } }
  // for i, oh := range bufferOverheat { bufferOverheat[i] = oh/mean }
  // fmt.Printf("\n ◦◦◦◦◦ DEBUG [Consuming heat][Overheat comparsion]: %s:%1.0f ", ElemSigns[0], mean); for i, h:=range bufferOverheat { if h>0 && i!=0 { fmt.Printf(" %1.2f'%s ", h, ElemSigns[i]) } }
  *&heatState.Current = bufferHeat
  *&heatState.Compared = primitives.GenerateHeat_CompareHeat(bufferHeat, streams)
  // PlotHeatState(*heatState)
}

func PlotHeatState(heat Heat) {
  fmt.Printf("\n ┌──── INFO [heat state]:")
  for i:=0 ; i<9; i++ {
    if i==0 {fmt.Printf("\n │ %s Average: %1.2f : ", ElemSigns[0], heat.Current[0])}
    if i==0 && heat.Compared[i]>0 {fmt.Printf("%1.2f ", heat.Compared[0])}
    if i!=0 && heat.Current[i]>0 {fmt.Printf("\n │ %s Rate: %1.2f : ", ElemSigns[i], heat.Current[i])}
    if i!=0 && heat.Compared[i]>0 && heat.Current[i]>0 {fmt.Printf("%1.2f ─── ", heat.Compared[i])}
    danger := heat.Current[i]>heat.Compared[i] && heat.Current[i]>0
    close := heat.Current[i]<=heat.Compared[i] && heat.Current[i]>0
    if i!=0 && close {
      fmt.Printf("Unstable: %1.1f%% ─── ", heat.Current[i]/heat.Compared[i]*100)
    } else if i!=0 && danger {
      fmt.Printf("Unstable: %1.1f%% ─── Danger: %1.2f ─── ", heat.Current[i]/heat.Compared[i]*100, heat.Current[i]-heat.Compared[i])
    }
  }
  fmt.Println()
  fmt.Printf(" └────────────────────────────────────────────────────────────────────────────────────────────────────")
}
