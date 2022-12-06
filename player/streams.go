package player

import "rhymald/mag-gamma/primitives"
import "math"
import "fmt"

type Streams struct {
  List []primitives.Stream
  SoulVolume float64
  Class  float64
  Herald float64
  Bender float64
  InternalElementalState [9]primitives.Stream
}

var (
  AllElements = primitives.AllElements
  ElemSigns   = primitives.ElemSigns
)

func PlotStreamList(list Streams, verbose bool) {
  multiplier := 1.0
  fmt.Printf(" ┌──── INFO [List strings]:\n")
  fmt.Printf(" │ Herald: %1.3f ─── Bender: %1.3f ─── Multiplier: %1.3f\n", list.Herald, list.Bender, list.Bender/list.Herald)
  if verbose==false { multiplier = list.Bender/list.Herald }
  for i, stream := range list.List {
    stats := primitives.StatsFromStream(stream)
    fmt.Printf(" │ ┌─ %d'%s ─── Basical stats: ▣ %1.3f ◈ %1.3f ▦ %1.3f \n", i+1, primitives.ES(stream.Element), stream.Creation, stream.Alteration, stream.Destruction)
    fmt.Printf(" │ │ Consumption time: %1.3f ms\n", stats["Quickness"]*multiplier)
    fmt.Printf(" │ │ Consumption count: %1.3f \n", stats["Fuel"]*multiplier)
    fmt.Printf(" │ │ Power: %1.3f\n", stats["Power"])
    fmt.Printf(" │ │ Precision: %1.3f\n", stats["Precision"]/multiplier)
    fmt.Printf(" │ └──── Overheat threshold: ◮ %1.3f\n", stream.Heat.Threshold)
  }
  fmt.Printf(" └────────────────────────────────────────────────────────────────────────────────────────────────────\n")
}

func NewBorn(yourStreams *Streams, class float64, standart float64, playerCount int) {
  fmt.Printf(" ┌──── DEBUG [Player creation][New initial streams]: start.\n │ ")
  stringsMatrix := Streams{}
  if class < 6.5 && class >= 0.5 {
    stringsMatrix.Class = class
  } else {
    for j:=0; j<2; j++ { stringsMatrix.Class += (primitives.RNF()*6+0.499999) / 2 }
  }
  stringsMatrix.Bender, stringsMatrix.Herald = primitives.NewBornStreams_BendHeraldFromClass(stringsMatrix.Class)
  countOfStreams := math.Round(stringsMatrix.Class)
  standart = standart
  empowering := ( - countOfStreams + stringsMatrix.Class )
  creUp, altUp, desUp := 1.0, 1.0, 1.0
  if empowering > 0 { desUp = (1+empowering) } else { creUp = (1-empowering) }
  altUp = (1.5 - math.Abs(empowering))
  fmt.Printf("INFO [Player creation][New initial streams]: defined %d class (%d streams), %+1.0f%% of power\n", int(stringsMatrix.Class*100000), int(countOfStreams), empowering*100)
  lens, wids, pows := []float64 {}, []float64 {}, []float64 {}
  slen, swid, spow := 0.0, 0.0, 0.0
  for i:=0; i<int(countOfStreams); i++ {
    leni := 1+primitives.RNF()
    widi := 1+primitives.RNF()
    powi := 1+primitives.RNF()
    lens, wids, pows = append(lens, leni), append(wids, widi), append(pows, powi)
    slen += leni ; swid += widi ; spow += powi
  }
  for i:=0; i<int(countOfStreams); i++ {
    var strings primitives.Stream
    strings.Element     = AllElements[i%3]
    strings.Creation    = lens[i] / slen * standart * creUp
    strings.Alteration  = wids[i] / swid * standart * altUp
    strings.Destruction = pows[i] / spow * standart * desUp
    // strings.InfoLWP = [3]float64{ primitives.NewBornStreams_LenFromStream(strings), primitives.NewBornStreams_WidFromStream(strings), primitives.NewBornStreams_PowFromStream(strings) }
    stringsMatrix.List = append(stringsMatrix.List, strings)
  }
  totalVol := primitives.StreamMean(AllElements[0], stringsMatrix.List)
  for i, _ := range stringsMatrix.List { stringsMatrix.List[i].Heat.Threshold = primitives.NewBorn_HeatThresholdFromStream(stringsMatrix.List[i], totalVol) }
  *yourStreams = stringsMatrix
  fmt.Printf(" │ DEBUG [Player creation][New initial streams]: done.\n")
  fmt.Printf(" └────────────────────────────────────────────────────────────────────────────────────────────────────\n")
}
