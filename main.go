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
  // ElemEnv ElementalAffinization
  // ElemExt ElementalAffinization
  // // ElemAff ElementalAffinization
  // ElemInt ElementalAffinization
  Resistance [8]float64
  Name string // character name
  Class float64 // born random
  StreamStrings []Stream // all the list
  Pool struct {
    Dots []Dot
    MaxVol float64
  }
}
// type ElementalAffinization [9]Stream
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
  YourHeat player.Heat
  // cosmetical
  You Player
  verbose = false
)

func init() {
  // some login, and middleware: regen and blockchain sync WorldInit()
  environment.Welling(&Environment) // fix to partial stack
  environment.Cursing(&Environment) // and here
  PlayerBorn(0)
  // player.PlotElementalState(YourStreams.InternalElementalState, "Internal elemental state", verbose)
  go func() { // passive prcoesses block
    go func() { for You.Health.Current >= 0 { player.RegenerateDots(&YourPool, YourStreams.List, verbose) } ; fmt.Println("FATAL: You are dead.")}()
    go func() { for You.Health.Current >= 0 { player.Transferrence(&YourPool, YourStreams.InternalElementalState, YourElemState, verbose) }     ; fmt.Println("FATAL: You are dead.")}()
    go func() { time.Sleep( time.Millisecond * time.Duration( primitives.Pool_RegenerateFullTimeOut() )) ; player.CalmDown(&YourHeat, YourStreams.InternalElementalState, verbose) }()
  }()
  fmt.Printf("SYSTEM [Start]:%s Welcome to the world, %s@%1.0f.\n", ebr, You.Name, YourStreams.Class*100000)
}

func main() {
  // here must be an interface
  fmt.Println("▼ EUA growling [from everywhere]:", ebr, "I smell you... your soul, your being. LEAVE!")
  Move(3, -17.2)
  // ii := 0
  player.EnergeticSurge(&YourPool, &YourHeat, YourStreams, 0.2, verbose) ; player.PlotHeatState(YourHeat)
  for {
    time.Sleep( time.Millisecond * time.Duration( primitives.Pool_RegenerateFullTimeOut() ))
    if primitives.RNF() < 0.71 { player.EnergeticSurge(&YourPool, &YourHeat, YourStreams, 0, verbose) ; player.PlotHeatState(YourHeat) }
    if primitives.RNF() < 0.71 { player.PlotEnergyStatus(YourPool, verbose) ; player.PlotHeatState(YourHeat) }
  }
  fmt.Scanln()
}

// func WorldInit() {// emulation
//   // campfires
//   campFireFire := Stream{Element: "Fire", Creation: 3, Destruction: 5, Alteration: 8}
//   CampFire     := PowerState{ Area: 5.0, Nature: []Stream{campFireFire}, Description: "Campfire: Warm place to rest at.", Concentrated: false}
//   CampFire.XYZs = append(CampFire.XYZs, [3]float64{1014.0, -16.5, 430.0})
//   // elementalTree
//   powerCore1 := Stream{Element: RNDElem(), Creation: math.Abs(Sprimitives.RNF()), Destruction: math.Abs(Sprimitives.RNF()), Alteration: math.Abs(Sprimitives.RNF())}
//   powerCore2 := Stream{Element: RNDElem(), Creation: math.Abs(Sprimitives.RNF()), Destruction: math.Abs(Sprimitives.RNF()), Alteration: math.Abs(Sprimitives.RNF())}
//   PowerCore     := PowerState{ Area: 25.0, Nature: []Stream{powerCore1, powerCore2}, Description: "Source of pure energy.", Concentrated: true}
//   PowerCore.XYZs = append(PowerCore.XYZs, [3]float64{1011.1, -8.5, 430.0})
//   fmt.Printf("Power core is emitting %s and %s energy...\n", primitives.ES(powerCore1.Element), primitives.ES(powerCore2.Element))
//   // windySpaces
//   openSpaceAir := Stream{Element: "Air", Creation: 1, Destruction: 2, Alteration: 3}
//   OpenSpace     := PowerState{ Area: 931.0, Nature: []Stream{openSpaceAir}, Description: "Winmdy weather.", Concentrated: false}
//   OpenSpace.XYZs = append(OpenSpace.XYZs, [3]float64{600.0, 100.5, 700.0})
//   // compose
//   Environment   = append(Environment, CampFire,OpenSpace,PowerCore)
// }

func Move(x float64, y float64) {
  fmt.Printf("▲ YOU moving, hurry [to people in front of you]:%s I am coming! ", cbr)
  if verbose {player.PlotElementalState(YourStreams.InternalElementalState, "Internal elemental state", verbose)}
  if verbose {player.PlotElementalState(YourElemState.ExternalWells, "Wells around", verbose)}
  // if verbose {player.PlotElementalState(YourElemState.Empowered, "Wells' affection",verbose)}
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
// func Orienting() {
//   var affectingPlaces []Stream
//   for _, place := range Environment.Wells {
//     for _, being := range place.XYZs {
//       distance := math.Sqrt(math.Pow(You.XYZ[0]-being[0],2)+math.Pow(You.XYZ[1]-being[1],2)+math.Pow(You.XYZ[2]-being[2],2))/place.Area
//       if distance <= 1 {
//         // fmt.Printf(" %1.2f from %s ",distance,place.Description)
//         for _, affection := range place.Nature {
//           if place.Concentrated {
//             buffer := Stream{
//               Element: affection.Element,
//               Creation: affection.Creation * math.Pow(1-distance, 2), // creation amount
//               Alteration: affection.Alteration * math.Pow(1-distance, 2), //creation quality
//               Destruction: affection.Destruction * math.Pow(1-distance, 2), // loose amount
//             }
//             affectingPlaces = append(affectingPlaces, buffer)
//           } else {
//             buffer := Stream{
//               Element: affection.Element,
//               Creation: affection.Creation * math.Sqrt(1-distance), // creation amount
//               Alteration: affection.Alteration * math.Sqrt(1-distance), //creation quality
//               Destruction: affection.Destruction * math.Sqrt(1-distance), // loose amount
//             }
//             affectingPlaces = append(affectingPlaces, buffer)
//           }
//         }
//       }
//     }
//   }
//   YourElemState = player.ElementalState{}
//   for _, affection := range affectingPlaces {
//     for i:=1;i<9;i++ {
//       if affection.Element == AllElements[i] {
//         YourElemState.External[i].Creation += affection.Creation
//         YourElemState.External[i].Alteration += affection.Alteration
//         YourElemState.External[i].Destruction += affection.Destruction
//       }
//     }
//   }
// }
// func InnerAffinization() {
//   // YourElemState.Empowered = ElementalAffinization{}
//   // YourElemState.Internal = ElementalAffinization{}
//   // You.ElemAff = ElementalAffinization{}
//   // Internal
//   for _, each := range YourStreams.List {
//     YourElemState.Internal[0].Creation += each.Creation // + YourElemState.Empowered[0].Creation)
//     YourElemState.Internal[0].Alteration += each.Alteration //  + YourElemState.Empowered[0].Alteration)
//     YourElemState.Internal[0].Destruction += each.Destruction //    + YourElemState.Empowered[0].Destruction)
//     if each.Element != "Common" {
//       YourElemState.Internal[primitives.ElemToInt(each.Element)].Creation += each.Creation // + YourElemState.Empowered[0].Creation)
//       YourElemState.Internal[primitives.ElemToInt(each.Element)].Alteration += each.Alteration //  + YourElemState.Empowered[0].Alteration)
//       YourElemState.Internal[primitives.ElemToInt(each.Element)].Destruction += each.Destruction //    + YourElemState.Empowered[0].Destruction)
//     }
//   }
//   // External basic
//   YourElemState.Empowered[1].Creation    = YourStreams.Bender * ( YourElemState.External[1].Creation    + YourElemState.External[3].Destruction) - YourElemState.External[2].Destruction * YourStreams.Herald
//   YourElemState.Empowered[1].Alteration  = YourStreams.Bender * ( YourElemState.External[1].Alteration  + YourElemState.External[3].Alteration)  - YourElemState.External[2].Alteration  * YourStreams.Herald
//   YourElemState.Empowered[1].Destruction = YourStreams.Bender * ( YourElemState.External[1].Destruction + YourElemState.External[3].Creation)    - YourElemState.External[2].Creation    * YourStreams.Herald
//   YourElemState.Empowered[2].Creation    = YourStreams.Bender * ( YourElemState.External[2].Creation    + YourElemState.External[1].Destruction) - YourElemState.External[4].Destruction * YourStreams.Herald
//   YourElemState.Empowered[2].Alteration  = YourStreams.Bender * ( YourElemState.External[2].Alteration  + YourElemState.External[1].Alteration)  - YourElemState.External[4].Alteration  * YourStreams.Herald
//   YourElemState.Empowered[2].Destruction = YourStreams.Bender * ( YourElemState.External[2].Destruction + YourElemState.External[1].Creation)    - YourElemState.External[4].Creation    * YourStreams.Herald
//   YourElemState.Empowered[3].Creation    = YourStreams.Bender * ( YourElemState.External[3].Creation    + YourElemState.External[4].Destruction) - YourElemState.External[1].Destruction * YourStreams.Herald
//   YourElemState.Empowered[3].Alteration  = YourStreams.Bender * ( YourElemState.External[3].Alteration  + YourElemState.External[4].Alteration)  - YourElemState.External[1].Alteration  * YourStreams.Herald
//   YourElemState.Empowered[3].Destruction = YourStreams.Bender * ( YourElemState.External[3].Destruction + YourElemState.External[4].Creation)    - YourElemState.External[1].Creation    * YourStreams.Herald
//   YourElemState.Empowered[4].Creation    = YourStreams.Bender * ( YourElemState.External[4].Creation    + YourElemState.External[2].Destruction) - YourElemState.External[3].Destruction * YourStreams.Herald
//   YourElemState.Empowered[4].Alteration  = YourStreams.Bender * ( YourElemState.External[4].Alteration  + YourElemState.External[2].Alteration)  - YourElemState.External[3].Alteration  * YourStreams.Herald
//   YourElemState.Empowered[4].Destruction = YourStreams.Bender * ( YourElemState.External[4].Destruction + YourElemState.External[2].Creation)    - YourElemState.External[3].Creation    * YourStreams.Herald
//   // v void - extra consumption
//   YourElemState.Empowered[5].Creation    = YourStreams.Bender * ( YourElemState.External[5].Creation) //    + YourElemState.Empowered[5].Creation)
//   YourElemState.Empowered[5].Alteration  = YourStreams.Bender * ( YourElemState.External[5].Alteration) //  + YourElemState.Empowered[5].Alteration)
//   YourElemState.Empowered[5].Destruction = YourStreams.Bender * ( YourElemState.External[5].Destruction) // + YourElemState.Empowered[5].Destruction)
//   // v deviant - extra overheat
//   YourElemState.Empowered[6].Creation    = YourStreams.Bender * ( YourElemState.External[6].Creation    + 2 * math.Sqrt(YourElemState.External[4].Creation    * YourElemState.External[2].Creation)) // + YourElemState.Empowered[6].Creation    )
//   YourElemState.Empowered[6].Alteration  = YourStreams.Bender * ( YourElemState.External[6].Alteration  + 2 * math.Sqrt(YourElemState.External[4].Alteration  * YourElemState.External[2].Creation)) // + YourElemState.Empowered[6].Alteration  )
//   YourElemState.Empowered[6].Destruction = YourStreams.Bender * ( YourElemState.External[6].Destruction + 2 * math.Sqrt(YourElemState.External[4].Destruction * YourElemState.External[2].Creation)) // + YourElemState.Empowered[6].Destruction )
//   YourElemState.Empowered[7].Creation    = YourStreams.Bender * ( YourElemState.External[7].Creation    + 2 * math.Sqrt(YourElemState.External[3].Creation    * YourElemState.External[1].Creation)) //  + YourElemState.Empowered[7].Creation    )
//   YourElemState.Empowered[7].Alteration  = YourStreams.Bender * ( YourElemState.External[7].Alteration  + 2 * math.Sqrt(YourElemState.External[3].Alteration  * YourElemState.External[1].Creation)) //  + YourElemState.Empowered[7].Alteration  )
//   YourElemState.Empowered[7].Destruction = YourStreams.Bender * ( YourElemState.External[7].Destruction + 2 * math.Sqrt(YourElemState.External[3].Destruction * YourElemState.External[1].Creation)) //  + YourElemState.Empowered[7].Destruction )
//   // v rarest - extra overheat and consumption
//   YourElemState.Empowered[8].Creation    = YourStreams.Bender * ( YourElemState.External[8].Alteration + YourElemState.External[8].Creation   )//  + YourElemState.Empowered[8].Creation)
//   YourElemState.Empowered[8].Alteration  = YourStreams.Bender * ( YourElemState.External[8].Creation   + YourElemState.External[8].Destruction)//  + YourElemState.Empowered[8].Alteration)
//   YourElemState.Empowered[8].Destruction = YourStreams.Bender * ( YourElemState.External[8].Alteration + YourElemState.External[8].Destruction) // + YourElemState.Empowered[8].Destruction)
//   // Finalizing
//   YourElemState.Empowered[0].Creation = YourStreams.Bender * YourElemState.External[8].Creation - YourElemState.External[5].Destruction * YourStreams.Herald + YourElemState.Internal[0].Creation
//   YourElemState.Empowered[0].Alteration = YourStreams.Bender * YourElemState.External[8].Alteration - YourElemState.External[5].Alteration * YourStreams.Herald + YourElemState.Internal[0].Alteration
//   YourElemState.Empowered[0].Destruction = YourStreams.Bender * YourElemState.External[8].Destruction - YourElemState.External[5].Creation * YourStreams.Herald + YourElemState.Internal[0].Destruction
//   for i:=1; i<9; i++ {
//     // if YourElemState.Empowered[i].Creation != 0 {
//     //   You.ElemAff[i].Creation += YourElemState.Empowered[i].Creation + YourElemState.Internal[i].Creation
//     //   You.ElemAff[i].Alteration += YourElemState.Empowered[i].Alteration + YourElemState.Internal[i].Alteration
//     //   You.ElemAff[i].Destruction += YourElemState.Empowered[i].Destruction + YourElemState.Internal[i].Destruction
//     // }
//     if YourElemState.Internal[i].Creation*YourElemState.Internal[i].Destruction != 0 { You.Resistance[i-1] = YourElemState.Internal[i].Destruction * YourElemState.Internal[i].Creation }
//   }
// }
// func PlotEnvAff() {
//   i := 0
//   if verbose {
//     fmt.Printf("\nDEBUG [fundamental attribute stats]: ")
//     if YourElemState.Empowered[0].Creation != 0 {fmt.Printf("%s ✳️  ─── %1.2f'vitality ─── %1.2f'concentration ─── %1.2f'power", elbr, YourElemState.Empowered[0].Creation, YourElemState.Empowered[0].Alteration, YourElemState.Empowered[0].Destruction)}
//   } else {
//     fmt.Printf("%s INFO [fundamental attribute balance]: ", cbr)
//     summ := YourElemState.Empowered[0].Creation+YourElemState.Empowered[0].Alteration+YourElemState.Empowered[0].Destruction
//     if YourElemState.Empowered[0].Creation != 0 {fmt.Printf("%s ✳️  ─── %2.1f%%'vitality ─── %2.1f%%'concentration ─── %2.1f%%'power", elbr, YourElemState.Empowered[0].Creation/summ*100, YourElemState.Empowered[0].Alteration/summ*100, YourElemState.Empowered[0].Destruction/summ*100)}
//   }
//   fmt.Printf("\nINFO [resistances]: ")
//   fmt.Printf(" %s:%1.2f ───", ElemSigns[0], (YourElemState.Internal[0].Creation)*(YourElemState.Internal[0].Alteration)*(YourElemState.Internal[0].Destruction) )
//   for i:=0; i<8; i++ { if You.Resistance[i] != 0 { fmt.Printf(" %s:%1.2f ───", ElemSigns[i+1], You.Resistance[i] ) } }
//   if verbose {
//     fmt.Printf("\nDEBUG [surrending elemental state]: ")
//     if YourElemState.External[1].Creation != 0 {fmt.Printf("%s ☁️  ──── %1.2f'pressure ────── %1.2f'spreading ──── %1.2f'puncture ───────", elbr, YourElemState.External[1].Creation, YourElemState.External[1].Alteration, YourElemState.External[1].Destruction) ; i++}
//     if YourElemState.External[2].Creation != 0 {fmt.Printf("%s 🔥 ──── %1.2f'warming ─────── %1.2f'burning ────── %1.2f'detonation ─────", elbr, YourElemState.External[2].Creation, YourElemState.External[2].Alteration, YourElemState.External[2].Destruction) ; i++}
//     if YourElemState.External[3].Creation != 0 {fmt.Printf("%s ⛰  ──── %1.2f'structure ───── %1.2f'mass ───────── %1.2f'fragility ──────", elbr, YourElemState.External[3].Creation, YourElemState.External[3].Alteration, YourElemState.External[3].Destruction) ; i++}
//     if YourElemState.External[4].Creation != 0 {fmt.Printf("%s 🧊 ──── %1.2f'form ─────────── %1.2f'direction ─── %1.2f'calm ───────────", elbr, YourElemState.External[4].Creation, YourElemState.External[4].Alteration, YourElemState.External[4].Destruction) ; i++}
//     if YourElemState.External[5].Creation != 0 {fmt.Printf("%s 🌑 ───── %1.2f'shadows ─────── %1.2f'curse ──────── %1.2f'pain ───────────", elbr, YourElemState.External[5].Creation, YourElemState.External[5].Alteration, YourElemState.External[5].Destruction) ; i++}
//     if YourElemState.External[6].Creation != 0 {fmt.Printf("%s 🩸 ───── %1.2f'growing ─────── %1.2f'corruption ─── %1.2f'consumption ────", elbr, YourElemState.External[6].Creation, YourElemState.External[6].Alteration, YourElemState.External[6].Destruction) ; i++}
//     if YourElemState.External[7].Creation != 0 {fmt.Printf("%s 🎶 ───── %1.2f'inspiration ─── %1.2f'echo ───────── %1.2f'roar ───────────", elbr, YourElemState.External[7].Creation, YourElemState.External[7].Alteration, YourElemState.External[7].Destruction) ; i++}
//     if YourElemState.External[8].Creation != 0 {fmt.Printf("%s 🌟 ────── %1.2f'mirage ──────── %1.2f'matter ─────── %1.2f'desintegration ─", elbr, YourElemState.External[8].Creation, YourElemState.External[8].Alteration, YourElemState.External[8].Destruction) ; i++}
//     if i==0 {fmt.Printf("%s Totaly not affected by environment ", elbr)}
//     fmt.Printf("\nDEBUG [incoming elemental affection]: ")
//     if YourElemState.Empowered[1].Creation != 0 {fmt.Printf("%s ☁️  ──── %1.2f'pressure ────── %1.2f'spreading ──── %1.2f'puncture ───────", elbr, YourElemState.Empowered[1].Creation, YourElemState.Empowered[1].Alteration, YourElemState.Empowered[1].Destruction) }
//     if YourElemState.Empowered[2].Creation != 0 {fmt.Printf("%s 🔥 ──── %1.2f'warming ─────── %1.2f'burning ────── %1.2f'detonation ─────", elbr, YourElemState.Empowered[2].Creation, YourElemState.Empowered[2].Alteration, YourElemState.Empowered[2].Destruction) }
//     if YourElemState.Empowered[3].Creation != 0 {fmt.Printf("%s ⛰  ──── %1.2f'structure ───── %1.2f'mass ───────── %1.2f'fragility ──────", elbr, YourElemState.Empowered[3].Creation, YourElemState.Empowered[3].Alteration, YourElemState.Empowered[3].Destruction) }
//     if YourElemState.Empowered[4].Creation != 0 {fmt.Printf("%s 🧊 ──── %1.2f'form ─────────── %1.2f'direction ─── %1.2f'calm ───────────", elbr, YourElemState.Empowered[4].Creation, YourElemState.Empowered[4].Alteration, YourElemState.Empowered[4].Destruction) }
//     if YourElemState.Empowered[5].Creation != 0 {fmt.Printf("%s 🌑 ───── %1.2f'shadows ─────── %1.2f'curse ──────── %1.2f'pain ───────────", elbr, YourElemState.Empowered[5].Creation, YourElemState.Empowered[5].Alteration, YourElemState.Empowered[5].Destruction) }
//     if YourElemState.Empowered[6].Creation != 0 {fmt.Printf("%s 🩸 ───── %1.2f'growing ─────── %1.2f'corruption ─── %1.2f'consumption ────", elbr, YourElemState.Empowered[6].Creation, YourElemState.Empowered[6].Alteration, YourElemState.Empowered[6].Destruction) }
//     if YourElemState.Empowered[7].Creation != 0 {fmt.Printf("%s 🎶 ───── %1.2f'inspiration ─── %1.2f'echo ───────── %1.2f'roar ───────────", elbr, YourElemState.Empowered[7].Creation, YourElemState.Empowered[7].Alteration, YourElemState.Empowered[7].Destruction) }
//     if YourElemState.Empowered[8].Creation != 0 {fmt.Printf("%s 🌟 ────── %1.2f'mirage ──────── %1.2f'matter ─────── %1.2f'desintegration ─", elbr, YourElemState.Empowered[8].Creation, YourElemState.Empowered[8].Alteration, YourElemState.Empowered[8].Destruction) }
//     if i==0 {fmt.Printf("%s Totaly not affected by environment ", elbr)}
//     // fmt.Printf("\nDEBUG [finaly affecting elemental state]: ") -- dat will be curses
//     // if You.ElemAff[1].Creation != 0 {fmt.Printf("%s ☁️  ──── %1.2f'pressure ────── %1.2f'spreading ──── %1.2f'puncture ───────", elbr, You.ElemAff[1].Creation, You.ElemAff[1].Alteration, You.ElemAff[1].Destruction) }
//     // if You.ElemAff[2].Creation != 0 {fmt.Printf("%s 🔥 ──── %1.2f'warming ─────── %1.2f'burning ────── %1.2f'detonation ─────", elbr, You.ElemAff[2].Creation, You.ElemAff[2].Alteration, You.ElemAff[2].Destruction) }
//     // if You.ElemAff[3].Creation != 0 {fmt.Printf("%s ⛰  ──── %1.2f'structure ───── %1.2f'mass ───────── %1.2f'fragility ──────", elbr, You.ElemAff[3].Creation, You.ElemAff[3].Alteration, You.ElemAff[3].Destruction) }
//     // if You.ElemAff[4].Creation != 0 {fmt.Printf("%s 🧊 ──── %1.2f'form ─────────── %1.2f'direction ─── %1.2f'calm ───────────", elbr, You.ElemAff[4].Creation, You.ElemAff[4].Alteration, You.ElemAff[4].Destruction) }
//     // if You.ElemAff[5].Creation != 0 {fmt.Printf("%s 🌑 ───── %1.2f'shadows ─────── %1.2f'curse ──────── %1.2f'pain ───────────", elbr, You.ElemAff[5].Creation, You.ElemAff[5].Alteration, You.ElemAff[5].Destruction) }
//     // if You.ElemAff[6].Creation != 0 {fmt.Printf("%s 🩸 ───── %1.2f'growing ─────── %1.2f'corruption ─── %1.2f'consumption ────", elbr, You.ElemAff[6].Creation, You.ElemAff[6].Alteration, You.ElemAff[6].Destruction) }
//     // if You.ElemAff[7].Creation != 0 {fmt.Printf("%s 🎶 ───── %1.2f'inspiration ─── %1.2f'echo ───────── %1.2f'roar ───────────", elbr, You.ElemAff[7].Creation, You.ElemAff[7].Alteration, You.ElemAff[7].Destruction) }
//     // if You.ElemAff[8].Creation != 0 {fmt.Printf("%s 🌟 ────── %1.2f'mirage ──────── %1.2f'matter ─────── %1.2f'desintegration ─", elbr, You.ElemAff[8].Creation, You.ElemAff[8].Alteration, You.ElemAff[8].Destruction) }
//     // if i==0 {fmt.Printf("%s Totaly not affected by environment ", elbr)}
//     fmt.Printf("\nDEBUG [internal elemental state]: ")
//     if YourElemState.Internal[1].Creation != 0 {fmt.Printf("%s ☁️  ──── %1.2f'pressure ────── %1.2f'spreading ──── %1.2f'puncture ───────", elbr, YourElemState.Internal[1].Creation, YourElemState.Internal[1].Alteration, YourElemState.Internal[1].Destruction) }
//     if YourElemState.Internal[2].Creation != 0 {fmt.Printf("%s 🔥 ──── %1.2f'warming ─────── %1.2f'burning ────── %1.2f'detonation ─────", elbr, YourElemState.Internal[2].Creation, YourElemState.Internal[2].Alteration, YourElemState.Internal[2].Destruction) }
//     if YourElemState.Internal[3].Creation != 0 {fmt.Printf("%s ⛰  ──── %1.2f'structure ───── %1.2f'mass ───────── %1.2f'fragility ──────", elbr, YourElemState.Internal[3].Creation, YourElemState.Internal[3].Alteration, YourElemState.Internal[3].Destruction) }
//     if YourElemState.Internal[4].Creation != 0 {fmt.Printf("%s 🧊 ──── %1.2f'form ─────────── %1.2f'direction ─── %1.2f'calm ───────────", elbr, YourElemState.Internal[4].Creation, YourElemState.Internal[4].Alteration, YourElemState.Internal[4].Destruction) }
//     if YourElemState.Internal[5].Creation != 0 {fmt.Printf("%s 🌑 ───── %1.2f'shadows ─────── %1.2f'curse ──────── %1.2f'pain ───────────", elbr, YourElemState.Internal[5].Creation, YourElemState.Internal[5].Alteration, YourElemState.Internal[5].Destruction) }
//     if YourElemState.Internal[6].Creation != 0 {fmt.Printf("%s 🩸 ───── %1.2f'growing ─────── %1.2f'corruption ─── %1.2f'consumption ────", elbr, YourElemState.Internal[6].Creation, YourElemState.Internal[6].Alteration, YourElemState.Internal[6].Destruction) }
//     if YourElemState.Internal[7].Creation != 0 {fmt.Printf("%s 🎶 ───── %1.2f'inspiration ─── %1.2f'echo ───────── %1.2f'roar ───────────", elbr, YourElemState.Internal[7].Creation, YourElemState.Internal[7].Alteration, YourElemState.Internal[7].Destruction) }
//     if YourElemState.Internal[8].Creation != 0 {fmt.Printf("%s 🌟 ────── %1.2f'mirage ──────── %1.2f'matter ─────── %1.2f'desintegration ─", elbr, YourElemState.Internal[8].Creation, YourElemState.Internal[8].Alteration, YourElemState.Internal[8].Destruction) }
//   }
// }

// func RNDElem() string { return AllElements[ rand.New(rand.NewSource(time.Now().UnixNano())).Intn( len(AllElements)-1 ) +1 ] }
// func primitives.ES(elem string) string { return ElemSigns[primitives.ElemToInt(elem)] }
// func primitives.ElemToInt(elem string) int { for i, each := range AllElements {if elem == each {return i}} ; return -1 }
// func primitives.RNF() float64 {
//   x:= (time.Now().UnixNano())
//   // var in_bytes []byte = big.NewInt(x).Bytes()
//   in_bytes := make([]byte, 8)
//   binary.LittleEndian.PutUint64(in_bytes, uint64(x))
//
//   hsum := sha512.Sum512(in_bytes)
//   sum := binary.BigEndian.Uint64(hsum[:])
//   return rand.New(rand.NewSource( int64(sum) )).Float64()
// }
// // func LogMM(a float64, b float64) float64 { return primitives.Log1479(math.Max(a,b)-math.Min(a,b)+1) }
// func primitives.Log1479(a float64) float64 { return math.Log2(math.Abs(a)+1)/math.Log2(1.479) }
// func primitives.Sign(a float64) float64 { if a != 0 {return a/math.Abs(a)} else {return 0} }
// // func Chance(a float64) float64 { return primitives.Log1479(a)/primitives.Log1479(a+1) } //NU
// // func Phi(a int) float64 { return math.Round((math.Pow((math.Sqrt(5)+1)/2, float64(a))-math.Pow((1-math.Sqrt(5))/2, float64(a)))/math.Sqrt(5)) } //NU
// func Sprimitives.RNF()float64 { return ( 4*primitives.RNF() - 3*primitives.RNF() + primitives.RNF() - 2*primitives.RNF() )/( 2*primitives.RNF() - 4*primitives.RNF() + 3*primitives.RNF() - primitives.RNF() ) } //NU
// func primitives.ChancedRound(a float64) int {
//   b,l:=math.Ceil(a),math.Floor(a)
//   c:=math.Abs(math.Abs(a)-math.Abs(math.Min(b, l)))
//   if a<0 {c = 1-c}
//   // if a!= 0 {fmt.Printf("\n %1.2f - between %1.0f and %1.0f - with chance %1.2f",a, b, l, c)}
//   if primitives.RNF() < c {return int(b)} else {return int(l)}
//   return 0
// }

func PlayerBorn(class float64) {
  // Nullification and rename
  You = Player{XYZ: [3]float64{1009.8, 7.3, 428.9}}
  fmt.Printf("INPUT[new player]:%s We haven't seen you yet, who are you?.. ", ebr)
  fmt.Scanln(&You.Name)
  if You.Name == "Rhymald" || You.Name == "" {verbose = true}
  You.Health.Current = 1

  YourHeat = player.Heat{}
  player.NewBorn(&YourStreams, class, .35, 5)
  // Class randomizing
  // if class < 6.5 && class >= 0.5 {
  //   YourStreams.Class = class
  // } else {
  //   playerCountInDB := math.Round(math.Log10(  4917  )+3.5)
  //   YourStreams.Class = class
  //   for j:=0; j<int(playerCountInDB); j++ {
  //     YourStreams.Class += (primitives.RNF()*6+0.499999) /playerCountInDB
  //   }
  // }
  // // Stream attachment
  // countOfStreams := math.Round(YourStreams.Class)
  // YourStreams.Bender = math.Sqrt(7.5-YourStreams.Class)
  // YourStreams.Herald = math.Sqrt(0.5+YourStreams.Class)
  // standart := .9
  // standart = math.Cbrt(standart / countOfStreams)
  // // fmt.Println(standart)
  // You.Health.Max = standart * 16
  // empowering := ( - countOfStreams + YourStreams.Class )
  // if empowering < 0 { empowering = 1 / (1 + math.Abs(empowering)) } else { empowering = (1 + math.Abs(empowering)) }
  // empowering = math.Cbrt(empowering)
  // if verbose {fmt.Printf("DEBUG [Player creation]: %d streams, %1.0f%% of power\n", int(countOfStreams), empowering*100)}
  // // fmt.Printf("DEBUG [Player creation]:%s Count: %d, %1.2f%% of power\n", ebr, int(countOfStreams), empowering*100)
  // stringsMatrix := []primitives.Stream{}
  // lens, wids, pows, geomean := []float64 {}, []float64 {}, []float64 {}, 0.0
  // for i:=0; i<int(countOfStreams); i++ {
  //   leni, widi, powi := 0.1+primitives.RNF(), 0.1+primitives.RNF(), 0.1+primitives.RNF()
  //   lens, wids, pows = append(lens, leni), append(wids, widi), append(pows, powi)
  //   geomean := math.Cbrt(leni*widi*powi)
  //   lens[i], wids[i], pows[i] = lens[i]*(standart/geomean), wids[i]*(standart/geomean), pows[i]*(standart/geomean)
  //   // fmt.Println(leni, widi, powi, "=>",geomean, "=>", lens[i], wids[i], pows[i], "=>", lens[i]*wids[i]*pows[i])
  // }
  // geomean = math.Pow(geomean, float64(1/countOfStreams/3))
  // for i:=0; i<int(countOfStreams); i++ {
  //   var strings primitives.Stream
  //   strings.Element     = AllElements[0]
  //   strings.Creation    = lens[i] * empowering
  //   strings.Alteration  = wids[i]
  //   strings.Destruction = pows[i] / empowering
  //   stringsMatrix = append(stringsMatrix, strings)
  //   healthBonus := 0.0
  //   for _, each := range stringsMatrix { healthBonus += primitives.Log1479((each.Creation*each.Creation+1)*each.Alteration) }
  // }
  // fmt.Println("DEBUG: Max health:", You.Health.Max)
  // YourStreams.List = stringsMatrix
  You.Health.Max += 100
  player.ExtendPool(&YourPool, YourStreams.List, verbose)
  player.ReadStatesFromEnv(&YourElemState, You.XYZ, &YourStreams, Environment)
  player.InnerAffinization(&YourElemState, YourStreams.Bender, YourStreams.Herald)
  player.PlotStreamList(YourStreams, verbose)
}
// func ListStrings() {
//   var counter Stream
//   fmt.Printf("INFO [List strings]:")
//   vols := 0.0
//   counter.Creation, counter.Alteration, counter.Destruction = 0.0, 0.0, 0.0
//   for i, stream := range YourStreams.List {
//     vols += (stream.Alteration)*(stream.Destruction)*(stream.Creation)
//     counter.Creation += stream.Creation
//     counter.Alteration += stream.Alteration
//     counter.Destruction += stream.Destruction
//     fmt.Printf("%s %d ─── %s %1.1f'len ─── %1.2f'wid ─── %1.3f'pow ", elbr,i+1, primitives.ES(stream.Element), stream.LWP[0], stream.LWP[1], stream.LWP[2])
//     if verbose {fmt.Printf("%s ─── %1.2f'len ─── %1.2f'wid ─── %1.2f'pow ────── Volume: %1.1f ─────", primitives.ES(stream.Element), stream.Creation, stream.Alteration, stream.Destruction, (stream.Alteration)*(stream.Destruction)*(stream.Creation))}
//   }
//   if verbose == false {PlotEnvAff()}
//   fmt.Printf("%s Total: %1.2f'lens + %1.2f'wids + %1.2f'pows = Volume: %1.1f\n", ebr, counter.Creation, counter.Alteration, counter.Destruction, vols)
// }
// func ExtendPools() {
//   fmt.Printf("INFO [Extend dot capacity to maximum]:")
//   old := You.Pool.MaxVol
//   for _, stream := range YourStreams.List {
//     You.Pool.MaxVol += 32*math.Sqrt(1+stream.Creation)
//   }
//   You.Pool.MaxVol = math.Round(You.Pool.MaxVol)
//   if verbose {
//     fmt.Printf("\nDEBUG [Pool]: %1.0f'dots\n", You.Pool.MaxVol)
//   } else {
//     if old == 0 {old = You.Pool.MaxVol/2}
//     fmt.Printf("%s INFO [Pool]: %+2.1f%%'dots\n", ebr, (You.Pool.MaxVol/old-1)*100)
//   }
// }
// func EnergyStatus() {
//   sum, mean := 0.0, 0.0
//   fmt.Printf("\nINFO [List dots]:%s", elbr)
//   count := 0
//   span := int(math.Sqrt(2)*math.Sqrt( float64(len(You.Pool.Dots)+1) ))
//   if span > 61 {span = 61}
//   if verbose {span = 10}
//   for e:=0; e<9; e++ {
//     for _, dot := range You.Pool.Dots {
//       if dot.Element == AllElements[e] {
//         if (count)%span == 0 && count != 0 {
//           fmt.Printf("%s",elbr)
//         }
//         if verbose {fmt.Printf("─%5.2f'%s ──", dot.Weight, primitives.ES(dot.Element))} else {fmt.Printf("%s",primitives.ES(dot.Element))}
//         sum += dot.Weight
//         mean += 1/dot.Weight
//         count++
//       }
//     }
//     // if verbose && count != 0 {fmt.Printf("%s", elbr)}
//   }
//   // if (count)%span == 0 {
//   //   if count != len(You.Pool.Dots) {fmt.Printf("%s",elbr)}
//   // }
//   // if verbose != true {
//   //   for e:=0; e<int(You.Pool.MaxVol)-len(You.Pool.Dots); e++ {
//   //     if (count)%span == 0 && count != len(You.Pool.Dots) {
//   //       fmt.Printf("%s",elbr)
//   //     }
//   //     fmt.Printf("◯ ")
//   //     count++
//   //   }
//   // }
//   fmt.Printf("\n")
//   fmt.Printf("INFO [Energy status]:%s Total energy level: %2.1f%%", ebr, float64(len(You.Pool.Dots))/You.Pool.MaxVol*100)
//   if verbose {fmt.Printf(" ─ mean:avg = %2.1f%%, %1.2f / %1.2f ─── Life: %2.1f%%", float64(len(You.Pool.Dots))/mean/(sum/float64(len(You.Pool.Dots)))*100, float64(len(You.Pool.Dots))/mean, sum/float64(len(You.Pool.Dots)), You.Health.Current/You.Health.Max*100)}
// }

// func GainDot() {
//   if len(You.Pool.Dots) >= int(You.Pool.MaxVol) { time.Sleep( time.Millisecond * time.Duration( 4000 )) ; return }
//   picker := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(YourStreams.List))
//   element := YourStreams.List[picker].Element
//   weight := primitives.Log1479( YourStreams.List[picker].Alteration ) * (1 + primitives.RNF()) / 2
//   dot := Dot{Element: element, Weight: weight}
//   You.Pool.Dots = append(You.Pool.Dots, dot)
//   if You.Health.Current < You.Health.Max {
//     heal := math.Sqrt(weight)
//     You.Health.Current += heal
//     // if verbose {fmt.Printf("  %1.1f %1.1f  ", heal, weight)}
//   } else { You.Health.Current = You.Health.Max }
//   time.Sleep( time.Millisecond * time.Duration( 1000*primitives.Log1479(You.Pool.MaxVol)/math.Sqrt(You.Pool.MaxVol)*math.Sqrt(weight) ))
// }
// func Regenerate() {
//   if len(You.Pool.Dots) >= int(You.Pool.MaxVol) {
//     if verbose {fmt.Printf("\nDEBUG [regenerating]: nothing to regenerate. ")}
//     time.Sleep( time.Millisecond * time.Duration( 4000 ))
//     return
//   }
//   mana := int( math.Sqrt(You.Pool.MaxVol-float64(len(You.Pool.Dots))) )
//   if verbose {fmt.Printf("\nDEBUG [regenerating]: +%d dots. ", mana)}
//   for i:=0; i<mana; i++ {
//     if len(You.Pool.Dots) >= int(You.Pool.MaxVol) {
//       if verbose {fmt.Printf("\nDEBUG [regenerating]: nothing to regenerate. ")}
//       time.Sleep( time.Millisecond * time.Duration( 4000 ))
//       break
//     }
//     GainDot()
//   }
// }
//
// func CrackStream(stream primitives.Stream) { // need heat {
//   element := stream.Element
//   weight := primitives.Log1479( stream.Destruction ) * (primitives.RNF() + primitives.RNF()) / 2
//   dot := Dot{Element: element, Weight: weight}
//   You.Pool.Dots = append(You.Pool.Dots, dot)
//   // return heat[element] = sqrt(sqr(d+1)/sqr(l-1)/sqr(w-1)+1)
// }
// func EnergeticSurge(doze float64) { // need in time
//   fmt.Printf("\n▲ YOU yelling [around]:%s CHEERS! A-ah...", ebr)
//   if doze == 0 {
//     doze = 1 / YourStreams.List[0].Destruction
//     for _, string := range YourStreams.List { doze = math.Max(doze, 1 / string.Destruction) }
//   }
//   for _, string := range YourStreams.List {
//     i := 0.0
//     for {
//       CrackStream(string)
//       i += 1 / doze
//       if string.Destruction < i { break }
//     }
//   }
// }

// func MinusDot(index int) (string, float64) {
//   if index >= len(You.Pool.Dots) { index = rand.New(rand.NewSource(time.Now().UnixNano())).Intn( len(You.Pool.Dots) ) }
//   ddelement := You.Pool.Dots[index].Element
//   ddweight := You.Pool.Dots[index].Weight
//   You.Pool.Dots[index] = You.Pool.Dots[len(You.Pool.Dots)-1]
//   You.Pool.Dots = You.Pool.Dots[:len(You.Pool.Dots)-1]
//   return ddelement, ddweight
// }
// func DotTransferIn(e int) {
//   if verbose {fmt.Printf("Absorbing dots:")}
//   element := AllElements[e]
//   if float64(len(You.Pool.Dots)) >= You.Pool.MaxVol+math.Sqrt(float64(len(You.Pool.Dots))) { if verbose {fmt.Printf(" Full is energy.\n")} ; time.Sleep( time.Millisecond * time.Duration( 4000 )) ; return }
//   weight := primitives.Log1479( YourElemState.Empowered[e].Alteration ) * (1 + primitives.RNF()) / 2
//   dot := Dot{Element: element, Weight: weight}
//   You.Pool.Dots = append(You.Pool.Dots, dot)
//   step := 32*math.Sqrt(1+math.Abs(YourElemState.Empowered[e].Creation))
//   if verbose {fmt.Printf(" +%s'%1.2f", primitives.ES(element), weight )}
//   if verbose {fmt.Printf(", - dot absorbed for %1.3fs.\n", math.Log2(step)/math.Sqrt(step)*math.Sqrt(weight))}
//   time.Sleep( time.Millisecond * time.Duration( 1000* math.Log2(step)/math.Sqrt(step) *math.Sqrt(weight) ))
// }
// func DotTransferOut(e int) {
//   if verbose {fmt.Printf("Losing dots:")}
//   element := AllElements[e]
//   presense := []int{}
//   for i, dot := range You.Pool.Dots { if dot.Element == element {presense = append(presense, i)} }
//   if len(presense) == 0 { if verbose{fmt.Printf(" No such dots.\n")} ; time.Sleep( time.Millisecond * time.Duration( 4000 )) ; return }
//   killer := presense[rand.New(rand.NewSource(time.Now().UnixNano())).Intn( len(presense) )]
//   _, weight := MinusDot(killer)
//   step := 32*math.Sqrt(1+math.Abs(YourElemState.Empowered[e].Destruction)) //primitives.Log1479(math.Abs(YourElemState.Empowered[e].Destruction)) * (1 + primitives.RNF()) / 2
//   if verbose {fmt.Printf(" -%s'%1.2f", primitives.ES(element), weight)}
//   if verbose {fmt.Printf(", - dot is lost for %1.3fs.\n", math.Log2(step)/math.Sqrt(step)/math.Sqrt(weight))}
//   time.Sleep( time.Millisecond * time.Duration( 1000* math.Log2(step)/math.Sqrt(step) /math.Sqrt(weight) ))
// }
// func Transferrence() {
//   demand := [9]int{}
//   cooldown := 0.0
//   for i, source := range YourElemState.Empowered {
//     count := 0.0
//     if source.Creation < 0 { count = - math.Sqrt(1+math.Abs(source.Destruction)) * (1 + primitives.RNF()) / 2 } else { count = math.Sqrt(1+math.Abs(source.Creation)) * (1 + primitives.RNF()) / 2 }
//     if i == 0 { count = 0 }
//     demand[i] = primitives.ChancedRound(count * primitives.Sign(YourElemState.External[i].Creation))
//     cooldown = math.Max(math.Abs(count) * 500, cooldown)
//   }
//   if cooldown == 0 { cooldown = 2000 }
//   if verbose {fmt.Printf("\nDEBUG [transferrence]: %v dots, cooldown: %1.3fs \n", demand, cooldown/1000)}
//   wg := sync.WaitGroup{}
//   for e, _ := range demand {
//     amount := demand[e]
//     if demand[e] > 0 {
//       if verbose {fmt.Println("Gaining", amount, AllElements[e])}
//       wg.Add(1)
//       go func(e int){
//         defer wg.Done()
//         for j:=0; j<amount; j++ { DotTransferIn(e) }
//       }(e)
//     } else if demand[e] < 0 {
//       amount = 0 - demand[e]
//       if verbose {fmt.Println("Loosing", amount, AllElements[e])}
//       wg.Add(1)
//       go func(e int){
//         defer wg.Done()
//         for j:=0; j<amount; j++ { DotTransferOut(e) }
//       }(e)
//     }
//   }
//   time.Sleep( time.Millisecond * time.Duration( cooldown ))
//   wg.Wait()
// }
