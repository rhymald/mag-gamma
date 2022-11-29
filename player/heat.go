package player

import "rhymald/mag-gamma/primitives"
import "fmt"
import "math"
import "time"

func CalmDown(streams *Streams, verbose bool) {
  for index, _ := range *&streams.List {
    go func(i int){ for { CalmDown_CalmHeatState(&streams.List[i], *&streams.Herald, verbose) } }(index)
  }
}

func CalmDown_CalmHeatState(stream *primitives.Stream, herald float64, verbose bool) {
  oldheat := *&stream.Heat.Current
  newheat := oldheat - (math.Sqrt(oldheat)*2 + 1) / (32 + math.Log2(1+*&stream.Heat.Threshold)) * herald
  pause := primitives.Pool_RegenerateFullTimeOut()
  if newheat < 0 || oldheat <= 1/100 {
    newheat = 0
    if true {fmt.Printf("\n ◦◦◦◦◦ DEBUG The %s is chilled out", ElemSigns[primitives.ElemToInt(*&stream.Element)])}
  } else {
    pause = 256 * (primitives.RNF()+1)
    if true {fmt.Printf("\n ◦◦◦◦◦ DEBUG Chilling for %1.3f'%s for %1.3fs => Current = %1.2f", oldheat - newheat, ElemSigns[primitives.ElemToInt(*&stream.Element)], pause/1000, newheat)}
  }
  *&stream.Heat.Current = newheat
  time.Sleep( time.Millisecond * time.Duration( pause ))
}

func ConsumeHeat(stream primitives.Stream, heat float64) float64 {
  if stream.Element == "Common" {return 0}
  fmt.Printf("\n ◦◦◦◦◦ DEBUG [Consuming heat][Incoming heat]: %+1.0f'%s ", heat, ElemSigns[primitives.ElemToInt(stream.Element)] )
  newheat := stream.Heat.Current + heat
  fmt.Printf("Current heat rates: %1.0f ", newheat)
  return newheat
}

func PlotHeatState(streams []primitives.Stream) { // replace for flock
  fmt.Printf("\n ┌──── INFO [heat state]:")
  counter := 0
  for i, each := range streams {
    if each.Heat.Current>0 && each.Element!="Common" {
      fmt.Printf("\n │ %s'%d Current %1.2f / %1.2f ─── ", ElemSigns[primitives.ElemToInt(each.Element)], i+1, each.Heat.Current, each.Heat.Threshold)
      counter++
      close := each.Heat.Current/each.Heat.Threshold>math.Sqrt2/2
      danger := each.Heat.Current/each.Heat.Threshold>1
      stability, damage := primitives.GenerateHeat_EffectFromCurrentAndThresholds(each.Heat.Current, each.Heat.Threshold)
      if danger {
        fmt.Printf("Unstable: %1.1f%% ─── Danger: %1.1f%% ─── ", 100*stability, damage*100 )
      } else if close {
        fmt.Printf("Unstable: %1.1f%% ─── ", 100*stability )
      }
    }
  }
  if counter==0 {fmt.Printf("\n │ All streams are calm")}
  fmt.Println()
  fmt.Printf(" └────────────────────────────────────────────────────────────────────────────────────────────────────")
}
