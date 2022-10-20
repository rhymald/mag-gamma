package player

import "rhymald/mag-gamma/primitives"
import "math"
import "fmt"

type Streams struct {
  List []primitives.Stream
  Class  float64
  Herald float64
  Bender float64
}

var (
  AllElements = primitives.AllElements
  ElemSigns   = primitives.ElemSigns
)

func NewBorn(yourStreams *Streams, class float64, standart float64, playerCount int) {
  fmt.Printf(" ┌─ DEBUG [Player creation]: new initial streams - start.\n ├─ ")
  stringsMatrix := Streams{}
  if class < 6.5 && class >= 0.5 {
    stringsMatrix.Class = class
  } else {
    playerCountInDB := math.Round(math.Log10( float64(playerCount) )+3.5)
    for j:=0; j<int(playerCountInDB); j++ {
      stringsMatrix.Class += (primitives.RNF()*6+0.499999) / playerCountInDB
    }
  }
  stringsMatrix.Bender = math.Cbrt(7.5-stringsMatrix.Class)
  stringsMatrix.Herald = math.Cbrt(0.5+stringsMatrix.Class)
  countOfStreams := math.Round(stringsMatrix.Class)
  standart = standart
  empowering := ( - countOfStreams + stringsMatrix.Class )
  if empowering < 0 { empowering = 1 / math.Cbrt(1 + math.Abs(empowering)) } else { empowering = math.Cbrt(1 + math.Abs(empowering)) }
  empowering = math.Cbrt(empowering)
  fmt.Printf("INFO  [Player creation]: defined %d class (%d streams), %1.0f%% of power\n", int(stringsMatrix.Class*100000), int(countOfStreams), empowering*100)
  lens, wids, pows := []float64 {}, []float64 {}, []float64 {}
  slen, swid, spow := 0.0, 0.0, 0.0
  for i:=0; i<int(countOfStreams); i++ {
    leni, widi, powi := math.Cbrt(1+primitives.RNF()), math.Cbrt(1+primitives.RNF()), math.Cbrt(1+primitives.RNF())
    lens, wids, pows = append(lens, leni), append(wids, widi), append(pows, powi)
    slen += leni ; swid += widi ; spow += powi
  }
  for i:=0; i<int(countOfStreams); i++ {
    var strings primitives.Stream
    strings.Element     = AllElements[0]
    strings.Creation    = lens[i] * empowering / slen * standart
    strings.Alteration  = wids[i] / swid * standart
    strings.Destruction = pows[i] / empowering / spow * standart
    strings.HeatPrint   = strings.Destruction/strings.Creation
    strings.LWP = [3]float64{ strings.Creation*1024, strings.Alteration/(1+strings.Creation), strings.Destruction/(strings.Alteration+1)/(strings.Creation+1) }
    stringsMatrix.List = append(stringsMatrix.List, strings)
  }
  *yourStreams = stringsMatrix
  fmt.Printf(" └─ DEBUG [Player creation]: new initial streams - done.\n")
}
