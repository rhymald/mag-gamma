package player

import "rhymald/mag-gamma/primitives"
import "math"
import "fmt"
import "math/rand"
import "time"
import "sync"

type Pool struct {
  Dots []primitives.Dot
  MaxVol float64
  // PositiveElementalState struct {
  //   FromEnv  [9]primitives.Stream
  //   Composed [9]primitives.Stream
  // }
}

func CountDotsByElements(pool Pool) ([9]int, [9]int) {
  counter, balance := [9]int{}, [9]int{}
  for _, dot := range pool.Dots { counter[primitives.ElemToInt(dot.Element)]++ }
  for e, count := range counter { balance[e] = int(math.Round(float64(count)*100/pool.MaxVol)) }
  return counter, balance
}

func ExtendPool(pool *Pool, streams []primitives.Stream, verbose bool) {
  fmt.Printf(" ╶──── INFO [Extend dot capacity to maximum]: ")
  old := *&pool.MaxVol
  new := primitives.ExtendPool_MaxVolFromStreams(streams)
  // for _, stream := range streams {
  //   new += 32*math.Sqrt(1+stream.Creation)
  // }
  *&pool.MaxVol = math.Round(new)
  if verbose {
    fmt.Printf("= %1.0f'dots\n", new)
  } else {
    if old == 0 {old = new/2}
    fmt.Printf("%+2.1f%%'dots\n", (new/old-1)*100)
  }
  // fmt.Printf(" └────────────────────────────────────────────────────────────────────────────────────────────────────\n")
}

func PlotEnergyStatus(pool Pool, verbose bool) {
  sum, mean := 0.0, 0.0
  fmt.Printf("\n ┌──── INFO [List dots]:")
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
        if verbose {fmt.Printf("─ %1.2f'%s ─", dot.Weight, primitives.ES(dot.Element))} else {fmt.Printf("%s",primitives.ES(dot.Element))}
        sum += dot.Weight
        mean += 1/dot.Weight
        count++
      }
    }
  }
  fmt.Printf("\n")
  fmt.Printf(" │ Total energy level: %2.1f%%", float64(len(pool.Dots))/pool.MaxVol*100)
  if verbose {fmt.Printf(" ─ mean:avg = %2.1f%%, %1.2f / %1.2f ─── ", float64(len(pool.Dots))/mean/(sum/float64(len(pool.Dots)))*100, float64(len(pool.Dots))/mean, sum/float64(len(pool.Dots)))}
  fmt.Printf("\n")
  fmt.Printf(" └────────────────────────────────────────────────────────────────────────────────────────────────────")
  // _, bar := CountDotsByElements(pool)
  // fmt.Printf("\n")
  // for e:=0; e<9; e++ {
  //   for i:=0; i<bar[e]; i++ {fmt.Printf("%s", primitives.ElemSigns[e] )}
  // }
}

func EmitDot(pool *Pool, streams []primitives.Stream) {
  if len(*&pool.Dots) >= int(*&pool.MaxVol) { time.Sleep( time.Millisecond * time.Duration( primitives.Pool_RegenerateFullTimeOut() )) ; return }
  picker := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(streams))
  element := streams[picker].Element
  weight, pause, _ := primitives.EmitDot_DotWeightAndTimeoutFromStreamAndMaxVol(streams[picker], *&pool.MaxVol)
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
  fulltimeout := primitives.Pool_RegenerateFullTimeOut()
  if len(*&pool.Dots) >= int(*&pool.MaxVol) {
    if verbose {fmt.Printf("\n ◦◦◦◦◦ DEBUG [Regenerating]: nothing to regenerate. ")}
    time.Sleep( time.Millisecond * time.Duration( fulltimeout ))
    return
  }
  mana := primitives.RegenerateDots_PortionFromPool(*&pool.MaxVol, len(*&pool.Dots))
  if verbose {fmt.Printf("\n ◦◦◦◦◦ DEBUG [Regenerating]: +%d dots. ", mana)}
  for i:=0; i<mana; i++ {
    if len(*&pool.Dots) >= int(*&pool.MaxVol) {
      if verbose {fmt.Printf("\n ◦◦◦◦◦ DEBUG [Regenerating]: nothing to regenerate. ")}
      time.Sleep( time.Millisecond * time.Duration( fulltimeout ))
      break
    }
    EmitDot(pool, streams)
  }
}

func CrackStream(pool *Pool, stream primitives.Stream) float64 {
  element := stream.Element
  weight := primitives.CrackStream_DotWeightFromStream(stream)
  if element == "Common" { weight = 1 - 1 / math.Phi }
  dot := primitives.Dot{Element: element, Weight: weight}
  *&pool.Dots = append(*&pool.Dots, dot)
  // element = primitives.RNDElem()
  if len(*&pool.Dots) > int(*&pool.MaxVol) {dot.Weight *= float64(len(*&pool.Dots)) / *&pool.MaxVol }
  heat := primitives.GenerateHeat_FromStreamAndDot(stream, dot)
  return heat
}

func EnergeticSurge(pool *Pool, streams *Streams, doze float64, verbose bool) {
  verbose = true
  // heatGenerated := [9]float64{}
  if doze == 0 { doze = 1 / *&streams.List[0].Destruction ; for _, string := range *&streams.List { doze = math.Max(doze, 1 / string.Destruction) } }
  fmt.Printf("\n  ▲ YOU [yelling around]: CHEERS! A-ah... [drink %0.3f ml]", doze)
  for index, string := range streams.List {
    i := 0.0
    for {
      // heat := CrackStream(pool, string) // compose heat
      i += 1 / doze
      *&streams.List[index].Heat.Current = ConsumeHeat(string, CrackStream(pool, string) / *&streams.Bender )
      // heatGenerated = primitives.CollectHeat(heatGenerated, heat)
      if string.Destruction <= i { break }
    }
  }
  // heatGenerated = primitives.GenerateHeat_ComposeHeat(heatGenerated)
}

func MinusDot(pool *Pool, index int) (string, float64) {
  if index >= len(*&pool.Dots) { index = rand.New(rand.NewSource(time.Now().UnixNano())).Intn( len(*&pool.Dots) ) }
  ddelement := *&pool.Dots[index].Element
  ddweight := *&pool.Dots[index].Weight
  buffer := *&pool.Dots
  buffer[index] = buffer[len(buffer)-1]
  *&pool.Dots = buffer[:len(buffer)-1]
  return ddelement, ddweight
}

func DotTransferIn(pool *Pool, estate ElementalState, verbose bool, e int) {
  if verbose {fmt.Printf("\n ◦◦◦◦◦ DEBUG [Transference][Absorbing dots]:")}
  element := AllElements[e]
  if float64(len(*&pool.Dots)) >= *&pool.MaxVol+math.Sqrt(float64(len(*&pool.Dots))) { if verbose {fmt.Printf(" Energy full")} ; time.Sleep( time.Millisecond * time.Duration( primitives.Pool_RegenerateFullTimeOut() )) ; return }
  weight, pause := primitives.DotTransferIn_DotWeightAndTimeoutFromState(estate.Empowered[e])
  dot := primitives.Dot{Element: element, Weight: weight}
  *&pool.Dots = append(*&pool.Dots, dot)
  if verbose {fmt.Printf(" +%s'%1.2f", primitives.ES(element), weight )}
  if verbose {fmt.Printf(" for %1.3fs", pause/1000)}
  time.Sleep( time.Millisecond * time.Duration( pause ))
}

func DotTransferOut(pool *Pool, estate ElementalState, verbose bool, e int) {
  if verbose {fmt.Printf("\n ◦◦◦◦◦ DEBUG [Transference][Losing dots]:")}
  element := AllElements[e]
  presense := []int{}
  for i, dot := range *&pool.Dots { if dot.Element == element {presense = append(presense, i)} }
  if len(presense) == 0 { if verbose{fmt.Printf(" No such dots")} ; time.Sleep( time.Millisecond * time.Duration( primitives.Pool_RegenerateFullTimeOut() )) ; return }
  killer := presense[rand.New(rand.NewSource(time.Now().UnixNano())).Intn( len(presense) )]
  _, weight := MinusDot(pool, killer)
  pause := primitives.DotTransferOut_TimeoutFromWeightAndState(weight, estate.Empowered[e]) //primitives.Log1479(math.Abs(estate.Empowered[e].Destruction)) * (1 + primitives.RNF()) / 2
  if verbose {fmt.Printf(" -%s'%1.2f", primitives.ES(element), weight)}
  if verbose {fmt.Printf(" for %1.3fs", pause/1000)}
  time.Sleep( time.Millisecond * time.Duration( pause ))
}

func Transferrence(pool *Pool, istate [9]primitives.Stream, estate ElementalState, verbose bool) {
  // demand := [9]int{}
  // // cooldown := 0.0
  // for i, source := range estate.Empowered {
  //   count := primitives.Transference_DotCountFromState(source)
  //   // if source.Creation < 0 { count = - math.Sqrt(1+math.Abs(source.Destruction)) * (1 + primitives.RNF()) / 2 } else { count = math.Sqrt(1+math.Abs(source.Creation)) * (1 + primitives.RNF()) / 2 }
  //   if i == 0 { count = 0 }
  //   demand[i] = primitives.ChancedRound(count * primitives.Sign(estate.External[i].Creation))
  //   // cooldown = math.Max(math.Abs(count) * 500, cooldown)
  // }
  // // if cooldown == 0 { cooldown = 2000 }
  // cooldown := primitives.Transference_TotalCooldownFromDemand(demand)
  cooldown, demand := primitives.Transference_DotCountDemandAndTotalCooldownFromStates(estate.Empowered, istate)
  if verbose {fmt.Printf("\n ◦◦◦◦◦ DEBUG [Transference][Demand calculation]: %v dots, cooldown: %1.3fs ", demand, cooldown/1000)}
  wg := sync.WaitGroup{}
  for e, _ := range demand {
    amount := demand[e]
    if demand[e] > 0 {
      if verbose {fmt.Printf("\n ◦◦◦◦◦ DEBUG [Transference][Gaining]: %d %s", amount, AllElements[e])}
      wg.Add(1)
      go func(e int){
        defer wg.Done()
        for j:=0; j<amount; j++ { DotTransferIn(pool, estate, verbose, e) }
      }(e)
    } else if demand[e] < 0 {
      amount = 0 - demand[e]
      if verbose {fmt.Printf("\n ◦◦◦◦◦ DEBUG [Transference][Loosing] %d %s", amount, AllElements[e])}
      wg.Add(1)
      go func(e int){
        defer wg.Done()
        for j:=0; j<amount; j++ { DotTransferOut(pool, estate, verbose, e) }
      }(e)
    }
  }
  time.Sleep( time.Millisecond * time.Duration( cooldown ))
  wg.Wait()
}
