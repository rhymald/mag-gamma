package player

import "rhymald/mag-gamma/primitives"
import "fmt"
import "math"

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
  bufferHeat[0] = float64(counter) / avg //* float64(counter)
  *&heatState.Current = bufferHeat
  *&heatState.Compared = primitives.GenerateHeat_CompareHeat(bufferHeat, streams)
}

func PlotHeatState(heat Heat) {
  fmt.Printf("\n ┌──── INFO [heat state]:")
  for i:=0 ; i<9; i++ {
    if i==0 {fmt.Printf("\n │ %s Average: %1.2f : ", ElemSigns[0], heat.Current[0])}
    if i==0 && heat.Compared[i]>0 {fmt.Printf("%1.2f ", heat.Compared[0])}
    if i!=0 && heat.Current[i]>0 {fmt.Printf("\n │ %s Rate: %1.2f : ", ElemSigns[i], heat.Current[i])}
    if i!=0 && heat.Compared[i]>0 && heat.Current[i]>0 {fmt.Printf("%1.2f ─── ", heat.Compared[i])}
    // fatal := heat.Current[i]/heat.Compared[i]>math.Sqrt2 && heat.Current[i]>0
    // danger := heat.Current[i]>heat.Compared[i] && heat.Current[i]>0
    close := heat.Current[i]/heat.Compared[i]>math.Sqrt2/2 && heat.Current[i]>0
    danger := heat.Current[i]/heat.Compared[i]>1 && heat.Current[i]>0
    stability, damage := primitives.GenerateHeat_EffectFromCurrentAndThresholds(heat.Current[i], heat.Compared[i])
    if i!=0 && danger {
    //   fmt.Printf("Unstable: %1.1f%% ─── Danger: %1.1f%% ─── Fatal: %1.1f%% ─── ", 100/(heat.Current[i]/heat.Compared[i]+0.3), math.Sqrt(math.Log10(1+heat.Current[i]-heat.Compared[i])/10)*100, math.Sqrt(math.Log10(1+heat.Current[i]-heat.Compared[i])/10)*100 )
    // } else if i!=0 && danger {
      fmt.Printf("Unstable: %1.1f%% ─── Danger: %1.1f%% ─── ", 100*stability, damage*100 )
    } else if i!=0 && close {
      fmt.Printf("Unstable: %1.1f%% ─── ", 100*stability )
    }
  }
  fmt.Println()
  fmt.Printf(" └────────────────────────────────────────────────────────────────────────────────────────────────────")
}
