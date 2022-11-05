package player

import "rhymald/mag-gamma/primitives"
import "math"
import "fmt"
import "math/rand"
import "time"

type Pool struct {
  Dots []primitives.Dot
  MaxVol float64
}

func ExtendPools(pool *Pool, streams []primitives.Stream, verbose bool) {

  fmt.Printf(" ┌──── INFO [Extend dot capacity to maximum]:\n │ ")
  old := *&pool.MaxVol
  new := primitives.MaxVolFromStreams(streams)
  // for _, stream := range streams {
  //   new += 32*math.Sqrt(1+stream.Creation)
  // }
  *&pool.MaxVol = math.Round(new)
  if verbose {
    fmt.Printf("DEBUG [Pool]: = %1.0f'dots\n", new)
  } else {
    if old == 0 {old = new/2}
    fmt.Printf("INFO [Pool]: %+2.1f%%'dots\n", (new/old-1)*100)
  }
  fmt.Printf(" └────────────────────────────────────────────────────────────────────────────────────────────────────\n")

}

func PlotEnergyStatus(pool Pool, verbose bool) {
  sum, mean := 0.0, 0.0
  fmt.Printf(" ┌──── INFO [List dots]:")
  if verbose == false { fmt.Printf("\n │") }
  if verbose { fmt.Printf("\n ├") }
  count := 0
  span := int(math.Sqrt(2)*math.Sqrt( float64(len(pool.Dots)+1) ))
  if span > 61 {span = 61}
  if verbose {span = 10}
  for e:=0; e<9; e++ {
    for _, dot := range pool.Dots {
      if dot.Element == AllElements[e] {
        if verbose == false && (count)%span == 0 && count != 0 { fmt.Printf("\n │") }
        if verbose && (count)%span == 0 && count != 0 { fmt.Printf("\n ├") }
        if verbose {fmt.Printf("─%5.2f'%s ─", dot.Weight, primitives.ES(dot.Element))} else {fmt.Printf("%s",primitives.ES(dot.Element))}
        sum += dot.Weight
        mean += 1/dot.Weight
        count++
      }
    }
  }
  fmt.Printf("\n")
  fmt.Printf(" ├── INFO [Energy status]: Total energy level: %2.1f%%", float64(len(pool.Dots))/pool.MaxVol*100)
  if verbose {fmt.Printf(" ─ mean:avg = %2.1f%%, %1.2f / %1.2f ─── ", float64(len(pool.Dots))/mean/(sum/float64(len(pool.Dots)))*100, float64(len(pool.Dots))/mean, sum/float64(len(pool.Dots)))}
  fmt.Printf("\n")
  fmt.Printf(" └────────────────────────────────────────────────────────────────────────────────────────────────────\n")
}

func EmitDot(pool *Pool, streams []primitives.Stream) {
  fulltimeout := primitives.RegenerateFullTimeOut()
  if len(*&pool.Dots) >= int(*&pool.MaxVol) { time.Sleep( time.Millisecond * time.Duration( fulltimeout )) ; return }
  picker := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(streams))
  element := streams[picker].Element
  weight, pause, _ := primitives.DotWeightAndTimeoutForRegenerationFromStreamAndMaxVol(streams[picker], *&pool.MaxVol)
  dot := primitives.Dot{Element: element, Weight: weight}
  *&pool.Dots = append(*&pool.Dots, dot)
  // if You.Health.Current < You.Health.Max {
  //   heal := math.Sqrt(weight)
  //   You.Health.Current += heal
  //   // if verbose {fmt.Printf("  %1.1f %1.1f  ", heal, weight)}
  // } else { You.Health.Current = You.Health.Max }
  time.Sleep( time.Millisecond * time.Duration( pause ))
}

func RegenerateDots(pool *Pool, streams []primitives.Stream, verbose bool) {
  fulltimeout := primitives.RegenerateFullTimeOut()
  if len(*&pool.Dots) >= int(*&pool.MaxVol) {
    if verbose {fmt.Printf("\nDEBUG [regenerating]: nothing to regenerate. ")}
    time.Sleep( time.Millisecond * time.Duration( fulltimeout ))
    return
  }
  mana := primitives.RegenerationPortionFromPool(*&pool.MaxVol, len(*&pool.Dots))
  if verbose {fmt.Printf("\nDEBUG [regenerating]: +%d dots. ", mana)}
  for i:=0; i<mana; i++ {
    if len(*&pool.Dots) >= int(*&pool.MaxVol) {
      if verbose {fmt.Printf("\nDEBUG [regenerating]: nothing to regenerate. ")}
      time.Sleep( time.Millisecond * time.Duration( fulltimeout ))
      break
    }
    EmitDot(pool, streams)
  }
}