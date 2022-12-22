package player

import "rhymald/mag-gamma/primitives"
import "math"
import "fmt"

type Streams struct {
  List []primitives.Stream
  SoulVolume float64 // ?
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
  resistances := make([]float64, 9)
  multiplier := 1.0
  sum, rate := 0.0, 0.0
  fmt.Printf(" ┌──── INFO [List strings]:\n")
  fmt.Printf(" │ Herald: %1.3f ─── Bender: %1.3f ─── Multiplier: %1.3f\n", list.Herald, list.Bender, list.Bender/list.Herald)
  if true { multiplier = list.Bender/list.Herald } // does it harm balance?
  for i, stream := range list.List {
    rating, askillrate, dskillrate := 1.0, 0.0, 0.0
    stats := primitives.StatsFromStream(stream)
    fmt.Printf(" │ ┌─ %d'%s ─── Basical stats: ▣ %1.3f ◈ %1.3f ▦ %1.3f \n", i+1, primitives.ES(stream.Element), stream.Creation, stream.Alteration, stream.Destruction)
    dot := math.Pow(math.Log2(primitives.Vector(stream.Destruction,stream.Creation+1,stream.Alteration)/3.5+1)+1, 2) * 0.75 - 0.75
    rating += dot*stats["M-Fuel"]*stats["M-Quickness"]*(multiplier)*multiplier
    fmt.Printf(" │ │ Consumption count: %1.0f dots", stats["M-Fuel"]*(multiplier))
    fmt.Printf(" for time: %1.0f ms", stats["M-Quickness"]*(multiplier))
    fmt.Printf(" = avg weight: %1.3f avg\n", dot)
    // Creative
    fmt.Printf(" │ │ ○ Toughness: %2.3f \n", stats["C-Toughness"])
    if stats["Cd-Decay"] > 0 { fmt.Printf(" │ │   ◎ Decay: %+1.1f%%\n", 100*stats["Cd-Decay"]) }
    if stats["Ca-Restoration"] > 0 { fmt.Printf(" │ │   ◎ Restoration: %+1.1f%%\n", 100*stats["Ca-Restoration"]) }
    if stats["Cad-Summon"] > 0 { fmt.Printf(" │ │     ◉ Summon: %+1.1f%%\n", 100*stats["Cad-Summon"]) }
    dskillrate += rating*stats["C-Toughness"]
    askillrate += rating*(stats["C-Toughness"]*stats["D-Power"])*(1+stats["Cd-Decay"])/math.Pi
    dskillrate += rating*(stats["C-Toughness"]*stats["A-Concentration"])*(1+stats["Ca-Restoration"])/math.Pi
    dskillrate += rating*(stats["C-Toughness"]*stats["D-Power"]*stats["A-Concentration"])*(1+stats["Cad-Summon"])/math.Pi/math.Pi
    // Alterative
    fmt.Printf(" │ │ ○ Concentration: %2.3f \n", stats["A-Concentration"])
    if stats["Ad-Condition"] > 0 { fmt.Printf(" │ │   ◎ Condition: %+1.1f%%\n", 100*stats["Ad-Condition"]) }
    if stats["Ac-Boon"] > 0 { fmt.Printf(" │ │   ◎ Boon: %+1.1f%%\n", 100*stats["Ac-Boon"]) }
    if stats["Adc-Transformation"] > 0 { fmt.Printf(" │ │     ◉ Transformation: %+1.1f%%\n", 100*stats["Adc-Transformation"]) }
    askillrate += rating*(stats["A-Concentration"])
    askillrate += rating*(stats["D-Power"]*stats["A-Concentration"])*(1+stats["Ad-Condition"])/math.Pi
    dskillrate += rating*(stats["C-Toughness"]*stats["A-Concentration"])*(1+stats["Ac-Boon"])/math.Pi
    dskillrate += rating*(stats["C-Toughness"]*stats["D-Power"]*stats["A-Concentration"])*(1+stats["Adc-Transformation"])/math.Pi/math.Pi
    // Destructive
    fmt.Printf(" │ │ ○ Power: %2.3f \n", stats["D-Power"])
    //   Hit Damage
    // critdmg := math.Sqrt(primitives.Log1479((primitives.Vector(stream.Destruction+1,stream.Destruction,stream.Alteration)-stream.Destruction))+1)-1
    critdmg := 2*((primitives.Vector(stream.Destruction,stream.Alteration)+math.Sqrt(stream.Destruction+1)*2-2-stream.Alteration)/primitives.Vector(stream.Destruction,stream.Alteration))
    critrate := 0.5*((primitives.Vector(stream.Destruction,stream.Alteration)+math.Sqrt(stream.Alteration+1)/2-0.5-stream.Destruction)/primitives.Vector(stream.Destruction,stream.Alteration))
    fmt.Printf(" │ │   Criticat rate: %+1.1f%% and damage: %+1.1f%% == %+1.1f%% dps \n", critrate*100, critdmg*100, ((1+critrate)*(1+critdmg)-1)*100)
    if stats["Dc-Sharpening"] > 0 { fmt.Printf(" │ │   ◎ Sharpening: %+2.1f%%\n", 100*stats["Dc-Sharpening"]) }
    if stats["Da-Barrier"] > 0 { fmt.Printf(" │ │   ◎ Barrier: %+2.1f%%\n", 100*stats["Da-Barrier"]) }
    if stats["Dac-Disaster"] > 0 { fmt.Printf(" │ │     ◉ Disaster: %+2.1f%%\n", 100*stats["Dac-Disaster"]) }
    dskillrate += rating*(stats["D-Power"])
    dskillrate += rating*(stats["D-Power"]*stats["C-Toughness"])*(1+stats["Dc-Sharpening"])/math.Pi
    askillrate += rating*(stats["D-Power"]*stats["A-Concentration"])*(1+stats["Da-Barrier"])/math.Pi
    dskillrate += rating*(stats["C-Toughness"]*stats["D-Power"]*stats["A-Concentration"])*(1+stats["Dac-Disaster"])/math.Pi/math.Pi
    fmt.Printf(" │ └──── Overheat threshold: ◮ %1.3f\n", stream.Heat.Threshold)
    rate += askillrate*math.Cbrt(stream.Heat.Threshold)*dskillrate
    sum += math.Cbrt((stream.Destruction+1)*(stream.Creation+1)*(stream.Alteration+1))-1
    resistance := primitives.Vector( primitives.Log1479(stream.Alteration), math.Sqrt(primitives.Log1479(stream.Creation)+1)-1, math.Sqrt(primitives.Log1479(stream.Alteration)+1)-1 )
    resistances[primitives.ElemToInt(stream.Element)] += resistance
    fmt.Printf(" │     Additional resisance: %+1.3f \n", resistance )
    fmt.Printf(" │            Stream volume: %1.3f and rate: %+1.3f\n", math.Cbrt((stream.Destruction+1)*(stream.Creation+1)*(stream.Alteration+1))-1, math.Log2( askillrate*(stream.Heat.Threshold*dskillrate) +1)/math.Log2(1.01) )
  }
  fmt.Printf(" │      Total resistances rate: \n │")
  for i := 1; i<5; i++ { if resistances[i] != 0 {fmt.Printf(" %s: %1.3f   ", primitives.ElemSigns[i], resistances[i]) }}
  fmt.Println()
  fmt.Printf(" │      Total stream volume: %1.0f and rate: %1.0f\n", math.Log2(sum+1)/math.Log2(1.01), math.Log2(rate+1)/math.Log2(1.01))
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
  creUp, altUp, desUp := 1.0, 1+primitives.RNF()/4+primitives.RNF()/4, 1.0
  if empowering > 0 { desUp = (1+primitives.RNF()*empowering) ; creUp = 3.5-altUp-desUp } else { creUp = (1-primitives.RNF()*empowering) ; desUp = 3.5-altUp-creUp }
  fmt.Printf("INFO [Player creation][New initial streams]: defined %d class (%d streams), %+1.0f%% of power\n", int(stringsMatrix.Class*100000), int(countOfStreams), empowering*100)
  lens, wids, pows := []float64 {}, []float64 {}, []float64 {}
  slen, swid, spow := 0.0, 0.0, 0.0
  for i:=0; i<int(countOfStreams); i++ {
    leni := (7+9*primitives.RNF())
    widi := (7+9*primitives.RNF())
    powi := (7+9*primitives.RNF())
    lens, wids, pows = append(lens, leni), append(wids, widi), append(pows, powi)
    slen += leni ; swid += widi ; spow += powi
  }
  for i:=0; i<int(countOfStreams); i++ {
    var strings primitives.Stream
    strings.Element     = AllElements[i%3+2]
    strings.Creation    = lens[i] / slen * standart * math.Cbrt(creUp)
    strings.Alteration  = wids[i] / swid * standart * math.Cbrt(altUp)
    strings.Destruction = pows[i] / spow * standart * math.Cbrt(desUp)
    // strings.InfoLWP = [3]float64{ primitives.NewBornStreams_LenFromStream(strings), primitives.NewBornStreams_WidFromStream(strings), primitives.NewBornStreams_PowFromStream(strings) }
    stringsMatrix.List = append(stringsMatrix.List, strings)
  }
  totalVol := primitives.StreamMean(AllElements[0], stringsMatrix.List)
  for i, _ := range stringsMatrix.List { stringsMatrix.List[i].Heat.Threshold = primitives.NewBorn_HeatThresholdFromStream(stringsMatrix.List[i], totalVol) }
  *yourStreams = stringsMatrix
  fmt.Printf(" │ DEBUG [Player creation][New initial streams]: done.\n")
  fmt.Printf(" └────────────────────────────────────────────────────────────────────────────────────────────────────\n")
}
