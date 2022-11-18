package player

import "rhymald/mag-gamma/primitives"
import "math"
import "fmt"

type Streams struct {
  List []primitives.Stream
  Class  float64
  Herald float64
  Bender float64
  InternalElementalState [9]primitives.Stream
}

var (
  AllElements = primitives.AllElements
  ElemSigns   = primitives.ElemSigns
)

func CountStreamssByElements(streams Streams) [9]int {
  counter := [9]int{}
  for _, each := range streams.List { counter[primitives.ElemToInt(each.Element)]++ }
  // for e, count := range counter { balance[e] = int(math.Round(float64(count)*100/pool.MaxVol)) }
  return counter //, balance
}

func PlotStreamList(list Streams, verbose bool) {
  var counter primitives.Stream
  fmt.Printf(" ┌──── INFO [List strings]:\n")
  vols := 0.0
  counter.Creation, counter.Alteration, counter.Destruction = 0.0, 0.0, 0.0
  for i, stream := range list.List {
    vols += (stream.Alteration)*(stream.Destruction)*(stream.Creation)
    counter.Creation += stream.Creation
    counter.Alteration += stream.Alteration
    counter.Destruction += stream.Destruction
    fmt.Printf(" │ %d'%s %1.0f'len ── %1.1f'wid ── %1.2f'pow ──", i+1, primitives.ES(stream.Element), stream.LWP[0], stream.LWP[1], stream.LWP[2])
    if verbose {fmt.Printf(" %1.2f'cre ── %1.2f'alt ── %1.2f'des ── Volume: %1.2f ──", stream.Creation, stream.Alteration, stream.Destruction, primitives.Vector(stream.Alteration,stream.Destruction,stream.Creation))}
    fmt.Println()
  }
  fmt.Printf(" │ Total: %1.2f'lens + %1.2f'wids + %1.2f'pows = Volume: %1.2f ──\n", counter.Creation, counter.Alteration, counter.Destruction, primitives.Vector(counter.Creation,counter.Alteration,counter.Destruction,))
  fmt.Printf(" └────────────────────────────────────────────────────────────────────────────────────────────────────\n")
}

func NewBorn(yourStreams *Streams, class float64, standart float64, playerCount int) {
  fmt.Printf(" ┌──── DEBUG [Player creation][New initial streams]: start.\n │ ")
  stringsMatrix := Streams{}
  if class < 6.5 && class >= 0.5 {
    stringsMatrix.Class = class
  } else {
    playerCountInDB := math.Round( float64(playerCount)*math.Cos(float64(playerCount)) +3.5)
    for j:=0; j<int(playerCountInDB); j++ {
      stringsMatrix.Class += (primitives.RNF()*6+0.499999) / playerCountInDB
    }
  }
  stringsMatrix.Bender, stringsMatrix.Herald = primitives.NewBornStreams_BendHeraldFromClass(stringsMatrix.Class)
  countOfStreams := math.Round(stringsMatrix.Class)
  standart = standart
  empowering := ( - countOfStreams + stringsMatrix.Class )
  if empowering < 0 { empowering = 1 / math.Cbrt(1 + math.Abs(empowering)) } else { empowering = math.Cbrt(1 + math.Abs(empowering)) }
  empowering = math.Cbrt(empowering)
  fmt.Printf("INFO [Player creation][New initial streams]: defined %d class (%d streams), %1.0f%% of power\n", int(stringsMatrix.Class*100000), int(countOfStreams), empowering*100)
  lens, wids, pows := []float64 {}, []float64 {}, []float64 {}
  slen, swid, spow := 0.0, 0.0, 0.0
  for i:=0; i<int(countOfStreams); i++ {
    leni, widi, powi := math.Cbrt(1+primitives.RNF()), math.Cbrt(1+primitives.RNF()), math.Cbrt(1+primitives.RNF())
    lens, wids, pows = append(lens, leni), append(wids, widi), append(pows, powi)
    slen += leni ; swid += widi ; spow += powi
  }
  for i:=0; i<int(countOfStreams); i++ {
    var strings primitives.Stream
    strings.Element     = AllElements[7-i]
    strings.Creation    = lens[i] * empowering / slen * standart
    strings.Alteration  = wids[i] / swid * standart
    strings.Destruction = pows[i] / empowering / spow * standart
    strings.HeatPrint   = strings.Destruction/strings.Creation
    strings.LWP = [3]float64{ primitives.NewBornStreams_LenFromStream(strings), primitives.NewBornStreams_WidFromStream(strings), primitives.NewBornStreams_PowFromStream(strings) }
    stringsMatrix.List = append(stringsMatrix.List, strings)
  }
  *yourStreams = stringsMatrix
  fmt.Printf(" │ DEBUG [Player creation][New initial streams]: done.\n")
  fmt.Printf(" └────────────────────────────────────────────────────────────────────────────────────────────────────\n")
}
