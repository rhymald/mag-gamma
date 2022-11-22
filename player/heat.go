package player

import "rhymald/mag-gamma/primitives"
import "fmt"
import "math"
import "time"

type Heat struct {
  Current [9]float64
  Compared [9]float64
  Unstable [9]float64
  Danger [9]float64
}

func CalmDown(heatState *Heat, istate [9]primitives.Stream, verbose bool) {
  for i:=1; i<9; i++ {
    go func(e int){ for { CalmDown_CalmHeatState(heatState, e, istate, verbose) } }(i)
  }
}

func CalmDown_CalmHeatState(heatState *Heat, element int, istate [9]primitives.Stream, verbose bool) {
  oldheat := *&heatState.Current[element]
  newheat := oldheat - (math.Sqrt(oldheat + *&heatState.Compared[0])*2 + 1) / 8
  pause := primitives.Pool_RegenerateFullTimeOut()
  if newheat < 0 || oldheat <= 1 {
    newheat = 0
    if verbose {fmt.Printf("\n ◦◦◦◦◦ DEBUG The %s is chilled out", ElemSigns[element])}
  } else {
    pause = 256 * (primitives.RNF()+1)
    if verbose {fmt.Printf("\n ◦◦◦◦◦ DEBUG Chilling for %1.3f'%s for %1.3fs => Current = %1.2f", oldheat - newheat, ElemSigns[element], pause/1000, newheat)}
  }
  *&heatState.Current[element] = newheat
  *&heatState.Compared = primitives.GenerateHeat_CompareHeat(*&heatState.Current, istate)
  time.Sleep( time.Millisecond * time.Duration( pause ))
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
