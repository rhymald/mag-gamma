package main
import (
  // "crypto/sha512"
  // "math/rand"
  // "encoding/binary"
  "math"
  "fmt"
  "time"
  // "sync"
  "rhymald/mag-gamma/player"
  "rhymald/mag-gamma/primitives"
  "rhymald/mag-gamma/environment"
)

type Player struct {
  Health struct {
    Current float64
    Max float64 }
  Heat map[string]float64
  OverHeat map[string]float64
  XYZ [3]float64
  Resistance [8]float64
  Name string // character name
  Class float64 // born random
  StreamStrings []Stream // all the list
  Pool struct {
    Dots []Dot
    MaxVol float64
  }
}
type Stream struct {
  Element string
  Creation float64
  Alteration float64
  Destruction float64
}
type Dot struct {
  Weight float64
  Element string
}

const (
  cbr  = "\n  ├─"   // subtitle
  elbr = "\n  │" // list
  ebr  = "\n  └─"   // end message
)
var (
  // env-related
  AllElements = primitives.AllElements
  ElemSigns   = primitives.ElemSigns
  Environment environment.Location
  // player-related
  YourElemState player.ElementalState
  YourStreams player.Streams
  YourPool player.Pool
  // cosmetical
  You Player
  verbose = false
)

func init() {
  // some login, and middleware: regen and blockchain sync WorldInit()
  environment.Welling(&Environment) // fix to partial stack
  environment.Cursing(&Environment) // and here
  PlayerBorn(0)
  go func() { // passive prcoesses block
    go func() { for You.Health.Current >= 0 { player.RegenerateDots(&YourPool, YourStreams.List, verbose) } ; fmt.Println("FATAL: You are dead.")}()
    go func() { for You.Health.Current >= 0 { player.Transferrence(&YourPool, YourStreams.InternalElementalState, YourElemState, verbose) }     ; fmt.Println("FATAL: You are dead.")}()
    go func() { time.Sleep( time.Millisecond * time.Duration( primitives.Pool_RegenerateFullTimeOut() )) ; player.CalmDown(&YourStreams, verbose) }()
  }()
  fmt.Printf("SYSTEM [Start]:%s Welcome to the world, %s@%1.0f.\n", ebr, You.Name, YourStreams.Class*100000)
}

func main() {
  // here must be an interface
  fmt.Println("▼ EUA growling [from everywhere]:", ebr, "I smell you... your soul, your being. LEAVE!")
  Move(3, -17.2)
  // ii := 0
  player.EnergeticSurge(&YourPool, &YourStreams, 10, verbose) ; player.PlotHeatState(YourStreams.List)
  for {
    time.Sleep( time.Millisecond * time.Duration( primitives.Pool_RegenerateFullTimeOut() ))
    if primitives.RNF() < 0.71 { player.EnergeticSurge(&YourPool, &YourStreams, 0, verbose) ; player.PlotHeatState(YourStreams.List) }
    if primitives.RNF() < 0.71 { player.PlotEnergyStatus(YourPool, verbose) }
  }
  fmt.Scanln()
}

func Move(x float64, y float64) {
  fmt.Printf("▲ YOU moving, hurry [to people in front of you]:%s I am coming! ", cbr)
  if verbose {player.PlotElementalState(YourStreams.InternalElementalState, "Internal elemental state", verbose)}
  if verbose {player.PlotElementalState(YourElemState.ExternalWells, "Wells around", verbose)}
  fmt.Printf("%s", elbr)
  distance := primitives.Vector(x,y)
  for t:=0.0; t<distance/0.7; t+=0.7 {
    You.XYZ[0] += x*0.7/distance
    You.XYZ[1] += y*0.7/distance
    time.Sleep( time.Millisecond * time.Duration( math.Sqrt(0.5)*100 ))
    if verbose {fmt.Printf(" 🐾")} else {fmt.Printf(" 🐾")}
    player.ReadStatesFromEnv(&YourElemState, You.XYZ, &YourStreams, Environment)
    player.InnerAffinization(&YourElemState, YourStreams.Bender, YourStreams.Herald)
  }
  if verbose {player.PlotElementalState(YourElemState.ExternalWells, "Wells around", verbose)}
  if verbose {player.PlotElementalState(YourElemState.Empowered, "Wells' affection",verbose)}
  fmt.Printf("%s Here I am: %1.2f'long-, %1.2f'graditude.", ebr, You.XYZ[0],You.XYZ[1])
}

func PlayerBorn(class float64) {
  // Nullification and rename
  You = Player{XYZ: [3]float64{1009.8, 7.3, 428.9}}
  fmt.Printf("INPUT[new player]:%s We haven't seen you yet, who are you?.. ", ebr)
  fmt.Scanln(&You.Name)
  if You.Name == "Rhymald" || You.Name == "" {verbose = true}
  You.Health.Current = 1
  player.NewBorn(&YourStreams, class, .35, 5)
  You.Health.Max += 100
  player.ExtendPool(&YourPool, YourStreams.List, verbose)
  player.ReadStatesFromEnv(&YourElemState, You.XYZ, &YourStreams, Environment)
  player.InnerAffinization(&YourElemState, YourStreams.Bender, YourStreams.Herald)
  player.PlotStreamList(YourStreams, verbose)
}
