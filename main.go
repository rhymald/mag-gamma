package main
import (
  "crypto/sha512"
  "math/rand"
  "encoding/binary"
  "math"
  "fmt"
  "time"
  "sync"
)

type PowerState struct {
  Description string
  Nature []Stream
  Area float64
  XYZs [][3]float64
  Concentrated bool
}

type Player struct {
  Health struct {
    Current float64
    Max float64 }
  Heat map[string]float64
  OverHeat map[string]float64
  XYZ [3]float64
  ElemEnv ElementalAffinization
  ElemExt ElementalAffinization
  // ElemAff ElementalAffinization
  ElemInt ElementalAffinization
  Resistance [8]float64
  Name string // character name
  Class float64 // born random
  StreamStrings []Stream // all the list
  Pool struct {
    Dots []Dot
    MaxVol float64
  }
}
type ElementalAffinization [9]Stream
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
  elbr = "\n  ├───" // list
  ebr  = "\n  └─"   // end message
)
var (
  verbose = false
  AllElements [9]string = [9]string{"Common", "Air", "Fire", "Earth", "Water", "Void", "Mallom", "Noise", "Resonance"}
  ElemSigns [9]string = [9]string{"✳️ ", "☁️ ", "🔥", "⛰ ", "🧊", "🌑", "🩸", "🎶", "🌟"}
  Environment []PowerState
  You Player
)

func init() {
  fmt.Println("▼ EUA growling [from everywhere]:", ebr, "I smell you... your soul, your being. LEAVE!")
  WorldInit()
  PlayerBorn(6.4)
  go func() { // passive prcoesses block
    go func() { for You.Health.Current > 0 {Regenerate()}    ; fmt.Println("FATAL: You are dead.")}()
    go func() { for You.Health.Current > 0 {Transferrence()} ; fmt.Println("FATAL: You are dead.")}()
  }()
  fmt.Printf("SYSTEM [Start]:%s Welcome to the world, %s@%1.0f.\n", ebr, You.Name, You.Class*100000)
}

func main() {
  Move(3, -17.2)
  for {
    if RNF() < 0.25 {EnergeticSurge(0.11)}
    EnergyStatus()
    time.Sleep( time.Second * time.Duration( 10 ))
    // if len(You.Pool.Dots) =
  }
  fmt.Scanln()
}

func WorldInit() {// emulation
  // campfires
  campFireFire := Stream{Element: "Fire", Creation: 3, Destruction: 5, Alteration: 8}
  CampFire     := PowerState{ Area: 5.0, Nature: []Stream{campFireFire}, Description: "Campfire: Warm place to rest at.", Concentrated: false}
  CampFire.XYZs = append(CampFire.XYZs, [3]float64{1014.0, -16.5, 430.0})
  // elementalTree
  powerCore1 := Stream{Element: RNDElem(), Creation: math.Abs(SRNF()), Destruction: math.Abs(SRNF()), Alteration: math.Abs(SRNF())}
  powerCore2 := Stream{Element: RNDElem(), Creation: math.Abs(SRNF()), Destruction: math.Abs(SRNF()), Alteration: math.Abs(SRNF())}
  PowerCore     := PowerState{ Area: 25.0, Nature: []Stream{powerCore1, powerCore2}, Description: "Source of pure energy.", Concentrated: true}
  PowerCore.XYZs = append(PowerCore.XYZs, [3]float64{1011.1, -8.5, 430.0})
  fmt.Printf("Power core is emitting %s and %s energy...\n", ES(powerCore1.Element), ES(powerCore2.Element))
  // windySpaces
  openSpaceAir := Stream{Element: "Air", Creation: 1, Destruction: 2, Alteration: 3}
  OpenSpace     := PowerState{ Area: 931.0, Nature: []Stream{openSpaceAir}, Description: "Winmdy weather.", Concentrated: false}
  OpenSpace.XYZs = append(OpenSpace.XYZs, [3]float64{600.0, 100.5, 700.0})
  // compose
  Environment   = append(Environment, CampFire,OpenSpace,PowerCore)
}

func Move(x float64, y float64) {
  fmt.Printf("▲ YOU moving, hurry [to people in front of you]:%s I am coming! ", cbr)
  if verbose {PlotEnvAff()}
  fmt.Printf("%s", elbr)
  distance := math.Sqrt(x*x+y*y)
  for t:=0.0; t<distance/0.7; t+=0.7 {
    You.XYZ[0] += x*0.7/distance
    You.XYZ[1] += y*0.7/distance
    time.Sleep( time.Millisecond * time.Duration( math.Sqrt(0.5)*1000 ))
    if verbose {fmt.Printf(" 🐾")} else {fmt.Printf(" 🐾")}
    Orienting()
    InnerAffinization()
  }
  if verbose {PlotEnvAff()}
  fmt.Printf("%s Here I am: %1.2f'long-, %1.2f'graditude.", ebr, You.XYZ[0],You.XYZ[1])
}
func Orienting() {
  var affectingPlaces []Stream
  for _, place := range Environment {
    for _, being := range place.XYZs {
      distance := math.Sqrt(math.Pow(You.XYZ[0]-being[0],2)+math.Pow(You.XYZ[1]-being[1],2)+math.Pow(You.XYZ[2]-being[2],2))/place.Area
      if distance <= 1 {
        // fmt.Printf(" %1.2f from %s ",distance,place.Description)
        for _, affection := range place.Nature {
          if place.Concentrated {
            buffer := Stream{
              Element: affection.Element,
              Creation: affection.Creation * math.Pow(1-distance, 2), // creation amount
              Alteration: affection.Alteration * math.Pow(1-distance, 2), //creation quality
              Destruction: affection.Destruction * math.Pow(1-distance, 2), // loose amount
            }
            affectingPlaces = append(affectingPlaces, buffer)
          } else {
            buffer := Stream{
              Element: affection.Element,
              Creation: affection.Creation * math.Sqrt(1-distance), // creation amount
              Alteration: affection.Alteration * math.Sqrt(1-distance), //creation quality
              Destruction: affection.Destruction * math.Sqrt(1-distance), // loose amount
            }
            affectingPlaces = append(affectingPlaces, buffer)
          }
        }
      }
    }
  }
  You.ElemEnv = ElementalAffinization{}
  for _, affection := range affectingPlaces {
    for i:=1;i<9;i++ {
      if affection.Element == AllElements[i] {
        You.ElemEnv[i].Creation += affection.Creation
        You.ElemEnv[i].Alteration += affection.Alteration
        You.ElemEnv[i].Destruction += affection.Destruction
      }
    }
  }
}
func InnerAffinization() {
  You.ElemExt = ElementalAffinization{}
  You.ElemInt = ElementalAffinization{}
  // You.ElemAff = ElementalAffinization{}
  // Internal
  for _, each := range You.StreamStrings {
    You.ElemInt[0].Creation += each.Creation // + You.ElemExt[0].Creation)
    You.ElemInt[0].Alteration += each.Alteration //  + You.ElemExt[0].Alteration)
    You.ElemInt[0].Destruction += each.Destruction //    + You.ElemExt[0].Destruction)
    if each.Element != "Common" {
      You.ElemInt[ElemToInt(each.Element)].Creation += each.Creation // + You.ElemExt[0].Creation)
      You.ElemInt[ElemToInt(each.Element)].Alteration += each.Alteration //  + You.ElemExt[0].Alteration)
      You.ElemInt[ElemToInt(each.Element)].Destruction += each.Destruction //    + You.ElemExt[0].Destruction)
    }
  }
  // External basic
  You.ElemExt[1].Creation    = math.Sqrt(7-You.Class) * ( You.ElemEnv[1].Creation    + You.ElemEnv[3].Destruction) - You.ElemEnv[2].Destruction * math.Sqrt(You.Class)
  You.ElemExt[1].Alteration  = math.Sqrt(7-You.Class) * ( You.ElemEnv[1].Alteration  + You.ElemEnv[3].Alteration)  - You.ElemEnv[2].Alteration  * math.Sqrt(You.Class)
  You.ElemExt[1].Destruction = math.Sqrt(7-You.Class) * ( You.ElemEnv[1].Destruction + You.ElemEnv[3].Creation)    - You.ElemEnv[2].Creation    * math.Sqrt(You.Class)
  You.ElemExt[2].Creation    = math.Sqrt(7-You.Class) * ( You.ElemEnv[2].Creation    + You.ElemEnv[1].Destruction) - You.ElemEnv[4].Destruction * math.Sqrt(You.Class)
  You.ElemExt[2].Alteration  = math.Sqrt(7-You.Class) * ( You.ElemEnv[2].Alteration  + You.ElemEnv[1].Alteration)  - You.ElemEnv[4].Alteration  * math.Sqrt(You.Class)
  You.ElemExt[2].Destruction = math.Sqrt(7-You.Class) * ( You.ElemEnv[2].Destruction + You.ElemEnv[1].Creation)    - You.ElemEnv[4].Creation    * math.Sqrt(You.Class)
  You.ElemExt[3].Creation    = math.Sqrt(7-You.Class) * ( You.ElemEnv[3].Creation    + You.ElemEnv[4].Destruction) - You.ElemEnv[1].Destruction * math.Sqrt(You.Class)
  You.ElemExt[3].Alteration  = math.Sqrt(7-You.Class) * ( You.ElemEnv[3].Alteration  + You.ElemEnv[4].Alteration)  - You.ElemEnv[1].Alteration  * math.Sqrt(You.Class)
  You.ElemExt[3].Destruction = math.Sqrt(7-You.Class) * ( You.ElemEnv[3].Destruction + You.ElemEnv[4].Creation)    - You.ElemEnv[1].Creation    * math.Sqrt(You.Class)
  You.ElemExt[4].Creation    = math.Sqrt(7-You.Class) * ( You.ElemEnv[4].Creation    + You.ElemEnv[2].Destruction) - You.ElemEnv[3].Destruction * math.Sqrt(You.Class)
  You.ElemExt[4].Alteration  = math.Sqrt(7-You.Class) * ( You.ElemEnv[4].Alteration  + You.ElemEnv[2].Alteration)  - You.ElemEnv[3].Alteration  * math.Sqrt(You.Class)
  You.ElemExt[4].Destruction = math.Sqrt(7-You.Class) * ( You.ElemEnv[4].Destruction + You.ElemEnv[2].Creation)    - You.ElemEnv[3].Creation    * math.Sqrt(You.Class)
  // v void - extra consumption
  You.ElemExt[5].Creation    = math.Sqrt(7-You.Class) * ( You.ElemEnv[5].Creation) //    + You.ElemExt[5].Creation)
  You.ElemExt[5].Alteration  = math.Sqrt(7-You.Class) * ( You.ElemEnv[5].Alteration) //  + You.ElemExt[5].Alteration)
  You.ElemExt[5].Destruction = math.Sqrt(7-You.Class) * ( You.ElemEnv[5].Destruction) // + You.ElemExt[5].Destruction)
  // v deviant - extra overheat
  You.ElemExt[6].Creation    = math.Sqrt(7-You.Class) * ( You.ElemEnv[6].Creation    + 2 * math.Sqrt(You.ElemEnv[4].Creation    * You.ElemEnv[2].Creation)) // + You.ElemExt[6].Creation    )
  You.ElemExt[6].Alteration  = math.Sqrt(7-You.Class) * ( You.ElemEnv[6].Alteration  + 2 * math.Sqrt(You.ElemEnv[4].Alteration  * You.ElemEnv[2].Creation)) // + You.ElemExt[6].Alteration  )
  You.ElemExt[6].Destruction = math.Sqrt(7-You.Class) * ( You.ElemEnv[6].Destruction + 2 * math.Sqrt(You.ElemEnv[4].Destruction * You.ElemEnv[2].Creation)) // + You.ElemExt[6].Destruction )
  You.ElemExt[7].Creation    = math.Sqrt(7-You.Class) * ( You.ElemEnv[7].Creation    + 2 * math.Sqrt(You.ElemEnv[3].Creation    * You.ElemEnv[1].Creation)) //  + You.ElemExt[7].Creation    )
  You.ElemExt[7].Alteration  = math.Sqrt(7-You.Class) * ( You.ElemEnv[7].Alteration  + 2 * math.Sqrt(You.ElemEnv[3].Alteration  * You.ElemEnv[1].Creation)) //  + You.ElemExt[7].Alteration  )
  You.ElemExt[7].Destruction = math.Sqrt(7-You.Class) * ( You.ElemEnv[7].Destruction + 2 * math.Sqrt(You.ElemEnv[3].Destruction * You.ElemEnv[1].Creation)) //  + You.ElemExt[7].Destruction )
  // v rarest - extra overheat and consumption
  You.ElemExt[8].Creation    = math.Sqrt(7-You.Class) * ( You.ElemEnv[8].Alteration + You.ElemEnv[8].Creation   )//  + You.ElemExt[8].Creation)
  You.ElemExt[8].Alteration  = math.Sqrt(7-You.Class) * ( You.ElemEnv[8].Creation   + You.ElemEnv[8].Destruction)//  + You.ElemExt[8].Alteration)
  You.ElemExt[8].Destruction = math.Sqrt(7-You.Class) * ( You.ElemEnv[8].Alteration + You.ElemEnv[8].Destruction) // + You.ElemExt[8].Destruction)
  // Finalizing
  You.ElemExt[0].Creation = math.Sqrt(7-You.Class) * You.ElemEnv[8].Creation - You.ElemEnv[5].Destruction * math.Sqrt(You.Class) + You.ElemInt[0].Creation
  You.ElemExt[0].Alteration = math.Sqrt(7-You.Class) * You.ElemEnv[8].Alteration - You.ElemEnv[5].Alteration * math.Sqrt(You.Class) + You.ElemInt[0].Alteration
  You.ElemExt[0].Destruction = math.Sqrt(7-You.Class) * You.ElemEnv[8].Destruction - You.ElemEnv[5].Creation * math.Sqrt(You.Class) + You.ElemInt[0].Destruction
  for i:=1; i<9; i++ {
    // if You.ElemExt[i].Creation != 0 {
    //   You.ElemAff[i].Creation += You.ElemExt[i].Creation + You.ElemInt[i].Creation
    //   You.ElemAff[i].Alteration += You.ElemExt[i].Alteration + You.ElemInt[i].Alteration
    //   You.ElemAff[i].Destruction += You.ElemExt[i].Destruction + You.ElemInt[i].Destruction
    // }
    if You.ElemInt[i].Creation*You.ElemInt[i].Destruction != 0 { You.Resistance[i-1] = You.ElemInt[i].Destruction * You.ElemInt[i].Creation }
  }
}
func PlotEnvAff() {
  i := 0
  if verbose {
    fmt.Printf("\nDEBUG [fundamental attribute stats]: ")
    if You.ElemExt[0].Creation != 0 {fmt.Printf("%s ✳️  ─── %1.2f'vitality ─── %1.2f'concentration ─── %1.2f'power", elbr, You.ElemExt[0].Creation, You.ElemExt[0].Alteration, You.ElemExt[0].Destruction)}
  } else {
    fmt.Printf("%s INFO [fundamental attribute balance]: ", cbr)
    summ := You.ElemExt[0].Creation+You.ElemExt[0].Alteration+You.ElemExt[0].Destruction
    if You.ElemExt[0].Creation != 0 {fmt.Printf("%s ✳️  ─── %2.1f%%'vitality ─── %2.1f%%'concentration ─── %2.1f%%'power", elbr, You.ElemExt[0].Creation/summ*100, You.ElemExt[0].Alteration/summ*100, You.ElemExt[0].Destruction/summ*100)}
  }
  fmt.Printf("\nINFO [resistances]: ")
  fmt.Printf(" %s:%1.2f ───", ElemSigns[0], (You.ElemInt[0].Creation)*(You.ElemInt[0].Alteration)*(You.ElemInt[0].Destruction) )
  for i:=0; i<8; i++ { if You.Resistance[i] != 0 { fmt.Printf(" %s:%1.2f ───", ElemSigns[i+1], You.Resistance[i] ) } }
  if verbose {
    fmt.Printf("\nDEBUG [surrending elemental state]: ")
    if You.ElemEnv[1].Creation != 0 {fmt.Printf("%s ☁️  ──── %1.2f'pressure ────── %1.2f'spreading ──── %1.2f'puncture ───────", elbr, You.ElemEnv[1].Creation, You.ElemEnv[1].Alteration, You.ElemEnv[1].Destruction) ; i++}
    if You.ElemEnv[2].Creation != 0 {fmt.Printf("%s 🔥 ──── %1.2f'warming ─────── %1.2f'burning ────── %1.2f'detonation ─────", elbr, You.ElemEnv[2].Creation, You.ElemEnv[2].Alteration, You.ElemEnv[2].Destruction) ; i++}
    if You.ElemEnv[3].Creation != 0 {fmt.Printf("%s ⛰  ──── %1.2f'structure ───── %1.2f'mass ───────── %1.2f'fragility ──────", elbr, You.ElemEnv[3].Creation, You.ElemEnv[3].Alteration, You.ElemEnv[3].Destruction) ; i++}
    if You.ElemEnv[4].Creation != 0 {fmt.Printf("%s 🧊 ──── %1.2f'form ─────────── %1.2f'direction ─── %1.2f'calm ───────────", elbr, You.ElemEnv[4].Creation, You.ElemEnv[4].Alteration, You.ElemEnv[4].Destruction) ; i++}
    if You.ElemEnv[5].Creation != 0 {fmt.Printf("%s 🌑 ───── %1.2f'shadows ─────── %1.2f'curse ──────── %1.2f'pain ───────────", elbr, You.ElemEnv[5].Creation, You.ElemEnv[5].Alteration, You.ElemEnv[5].Destruction) ; i++}
    if You.ElemEnv[6].Creation != 0 {fmt.Printf("%s 🩸 ───── %1.2f'growing ─────── %1.2f'corruption ─── %1.2f'consumption ────", elbr, You.ElemEnv[6].Creation, You.ElemEnv[6].Alteration, You.ElemEnv[6].Destruction) ; i++}
    if You.ElemEnv[7].Creation != 0 {fmt.Printf("%s 🎶 ───── %1.2f'inspiration ─── %1.2f'echo ───────── %1.2f'roar ───────────", elbr, You.ElemEnv[7].Creation, You.ElemEnv[7].Alteration, You.ElemEnv[7].Destruction) ; i++}
    if You.ElemEnv[8].Creation != 0 {fmt.Printf("%s 🌟 ────── %1.2f'mirage ──────── %1.2f'matter ─────── %1.2f'desintegration ─", elbr, You.ElemEnv[8].Creation, You.ElemEnv[8].Alteration, You.ElemEnv[8].Destruction) ; i++}
    if i==0 {fmt.Printf("%s Totaly not affected by environment ", elbr)}
    fmt.Printf("\nDEBUG [incoming elemental affection]: ")
    if You.ElemExt[1].Creation != 0 {fmt.Printf("%s ☁️  ──── %1.2f'pressure ────── %1.2f'spreading ──── %1.2f'puncture ───────", elbr, You.ElemExt[1].Creation, You.ElemExt[1].Alteration, You.ElemExt[1].Destruction) }
    if You.ElemExt[2].Creation != 0 {fmt.Printf("%s 🔥 ──── %1.2f'warming ─────── %1.2f'burning ────── %1.2f'detonation ─────", elbr, You.ElemExt[2].Creation, You.ElemExt[2].Alteration, You.ElemExt[2].Destruction) }
    if You.ElemExt[3].Creation != 0 {fmt.Printf("%s ⛰  ──── %1.2f'structure ───── %1.2f'mass ───────── %1.2f'fragility ──────", elbr, You.ElemExt[3].Creation, You.ElemExt[3].Alteration, You.ElemExt[3].Destruction) }
    if You.ElemExt[4].Creation != 0 {fmt.Printf("%s 🧊 ──── %1.2f'form ─────────── %1.2f'direction ─── %1.2f'calm ───────────", elbr, You.ElemExt[4].Creation, You.ElemExt[4].Alteration, You.ElemExt[4].Destruction) }
    if You.ElemExt[5].Creation != 0 {fmt.Printf("%s 🌑 ───── %1.2f'shadows ─────── %1.2f'curse ──────── %1.2f'pain ───────────", elbr, You.ElemExt[5].Creation, You.ElemExt[5].Alteration, You.ElemExt[5].Destruction) }
    if You.ElemExt[6].Creation != 0 {fmt.Printf("%s 🩸 ───── %1.2f'growing ─────── %1.2f'corruption ─── %1.2f'consumption ────", elbr, You.ElemExt[6].Creation, You.ElemExt[6].Alteration, You.ElemExt[6].Destruction) }
    if You.ElemExt[7].Creation != 0 {fmt.Printf("%s 🎶 ───── %1.2f'inspiration ─── %1.2f'echo ───────── %1.2f'roar ───────────", elbr, You.ElemExt[7].Creation, You.ElemExt[7].Alteration, You.ElemExt[7].Destruction) }
    if You.ElemExt[8].Creation != 0 {fmt.Printf("%s 🌟 ────── %1.2f'mirage ──────── %1.2f'matter ─────── %1.2f'desintegration ─", elbr, You.ElemExt[8].Creation, You.ElemExt[8].Alteration, You.ElemExt[8].Destruction) }
    if i==0 {fmt.Printf("%s Totaly not affected by environment ", elbr)}
    // fmt.Printf("\nDEBUG [finaly affecting elemental state]: ") -- dat will be curses
    // if You.ElemAff[1].Creation != 0 {fmt.Printf("%s ☁️  ──── %1.2f'pressure ────── %1.2f'spreading ──── %1.2f'puncture ───────", elbr, You.ElemAff[1].Creation, You.ElemAff[1].Alteration, You.ElemAff[1].Destruction) }
    // if You.ElemAff[2].Creation != 0 {fmt.Printf("%s 🔥 ──── %1.2f'warming ─────── %1.2f'burning ────── %1.2f'detonation ─────", elbr, You.ElemAff[2].Creation, You.ElemAff[2].Alteration, You.ElemAff[2].Destruction) }
    // if You.ElemAff[3].Creation != 0 {fmt.Printf("%s ⛰  ──── %1.2f'structure ───── %1.2f'mass ───────── %1.2f'fragility ──────", elbr, You.ElemAff[3].Creation, You.ElemAff[3].Alteration, You.ElemAff[3].Destruction) }
    // if You.ElemAff[4].Creation != 0 {fmt.Printf("%s 🧊 ──── %1.2f'form ─────────── %1.2f'direction ─── %1.2f'calm ───────────", elbr, You.ElemAff[4].Creation, You.ElemAff[4].Alteration, You.ElemAff[4].Destruction) }
    // if You.ElemAff[5].Creation != 0 {fmt.Printf("%s 🌑 ───── %1.2f'shadows ─────── %1.2f'curse ──────── %1.2f'pain ───────────", elbr, You.ElemAff[5].Creation, You.ElemAff[5].Alteration, You.ElemAff[5].Destruction) }
    // if You.ElemAff[6].Creation != 0 {fmt.Printf("%s 🩸 ───── %1.2f'growing ─────── %1.2f'corruption ─── %1.2f'consumption ────", elbr, You.ElemAff[6].Creation, You.ElemAff[6].Alteration, You.ElemAff[6].Destruction) }
    // if You.ElemAff[7].Creation != 0 {fmt.Printf("%s 🎶 ───── %1.2f'inspiration ─── %1.2f'echo ───────── %1.2f'roar ───────────", elbr, You.ElemAff[7].Creation, You.ElemAff[7].Alteration, You.ElemAff[7].Destruction) }
    // if You.ElemAff[8].Creation != 0 {fmt.Printf("%s 🌟 ────── %1.2f'mirage ──────── %1.2f'matter ─────── %1.2f'desintegration ─", elbr, You.ElemAff[8].Creation, You.ElemAff[8].Alteration, You.ElemAff[8].Destruction) }
    // if i==0 {fmt.Printf("%s Totaly not affected by environment ", elbr)}
    fmt.Printf("\nDEBUG [internal elemental state]: ")
    if You.ElemInt[1].Creation != 0 {fmt.Printf("%s ☁️  ──── %1.2f'pressure ────── %1.2f'spreading ──── %1.2f'puncture ───────", elbr, You.ElemInt[1].Creation, You.ElemInt[1].Alteration, You.ElemInt[1].Destruction) }
    if You.ElemInt[2].Creation != 0 {fmt.Printf("%s 🔥 ──── %1.2f'warming ─────── %1.2f'burning ────── %1.2f'detonation ─────", elbr, You.ElemInt[2].Creation, You.ElemInt[2].Alteration, You.ElemInt[2].Destruction) }
    if You.ElemInt[3].Creation != 0 {fmt.Printf("%s ⛰  ──── %1.2f'structure ───── %1.2f'mass ───────── %1.2f'fragility ──────", elbr, You.ElemInt[3].Creation, You.ElemInt[3].Alteration, You.ElemInt[3].Destruction) }
    if You.ElemInt[4].Creation != 0 {fmt.Printf("%s 🧊 ──── %1.2f'form ─────────── %1.2f'direction ─── %1.2f'calm ───────────", elbr, You.ElemInt[4].Creation, You.ElemInt[4].Alteration, You.ElemInt[4].Destruction) }
    if You.ElemInt[5].Creation != 0 {fmt.Printf("%s 🌑 ───── %1.2f'shadows ─────── %1.2f'curse ──────── %1.2f'pain ───────────", elbr, You.ElemInt[5].Creation, You.ElemInt[5].Alteration, You.ElemInt[5].Destruction) }
    if You.ElemInt[6].Creation != 0 {fmt.Printf("%s 🩸 ───── %1.2f'growing ─────── %1.2f'corruption ─── %1.2f'consumption ────", elbr, You.ElemInt[6].Creation, You.ElemInt[6].Alteration, You.ElemInt[6].Destruction) }
    if You.ElemInt[7].Creation != 0 {fmt.Printf("%s 🎶 ───── %1.2f'inspiration ─── %1.2f'echo ───────── %1.2f'roar ───────────", elbr, You.ElemInt[7].Creation, You.ElemInt[7].Alteration, You.ElemInt[7].Destruction) }
    if You.ElemInt[8].Creation != 0 {fmt.Printf("%s 🌟 ────── %1.2f'mirage ──────── %1.2f'matter ─────── %1.2f'desintegration ─", elbr, You.ElemInt[8].Creation, You.ElemInt[8].Alteration, You.ElemInt[8].Destruction) }
  }
}

func RNDElem() string { return AllElements[ rand.New(rand.NewSource(time.Now().UnixNano())).Intn( len(AllElements)-1 ) +1 ] }
func ES(elem string) string { return ElemSigns[ElemToInt(elem)] }
func ElemToInt(elem string) int { for i, each := range AllElements {if elem == each {return i}} ; return -1 }
func RNF() float64 {
  x:= (time.Now().UnixNano())
  // var in_bytes []byte = big.NewInt(x).Bytes()
  in_bytes := make([]byte, 8)
  binary.LittleEndian.PutUint64(in_bytes, uint64(x))

  hsum := sha512.Sum512(in_bytes)
  sum := binary.BigEndian.Uint64(hsum[:])
  return rand.New(rand.NewSource( int64(sum) )).Float64()
}
// func LogMM(a float64, b float64) float64 { return Log1479(math.Max(a,b)-math.Min(a,b)+1) }
func Log1479(a float64) float64 { return math.Log2(math.Abs(a)+1)/math.Log2(1.479) }
func Sign(a float64) float64 { if a != 0 {return a/math.Abs(a)} else {return 0} }
// func Chance(a float64) float64 { return Log1479(a)/Log1479(a+1) } //NU
// func Phi(a int) float64 { return math.Round((math.Pow((math.Sqrt(5)+1)/2, float64(a))-math.Pow((1-math.Sqrt(5))/2, float64(a)))/math.Sqrt(5)) } //NU
func SRNF()float64 { return ( 4*RNF() - 3*RNF() + RNF() - 2*RNF() )/( 2*RNF() - 4*RNF() + 3*RNF() - RNF() ) } //NU
func ChancedRound(a float64) int {
  b,l:=math.Ceil(a),math.Floor(a)
  c:=math.Abs(math.Abs(a)-math.Abs(math.Min(b, l)))
  if a<0 {c = 1-c}
  // if a!= 0 {fmt.Printf("\n %1.2f - between %1.0f and %1.0f - with chance %1.2f",a, b, l, c)}
  if RNF() < c {return int(b)} else {return int(l)}
  return 0
}

func PlayerBorn(class float64) {
  // Nullification and rename
  You = Player{XYZ: [3]float64{1009.8, 7.3, 428.9}}
  fmt.Printf("INPUT[new player]:%s We haven't seen you yet, who are you?.. ", ebr)
  fmt.Scanln(&You.Name)
  if You.Name == "Rhymald" || You.Name == "" {verbose = true}
  You.Health.Current = 1
  // Class randomizing
  if class < 6.5 && class >= 0.5 {
    You.Class = class
  } else {
    playerCountInDB := math.Round(math.Log10(  4917  )+3.5)
    You.Class = class
    for j:=0; j<int(playerCountInDB); j++ {
      You.Class += (RNF()*6+0.499999) /playerCountInDB
    }
  }
  // Stream attachment
  countOfStreams := math.Round(You.Class)
  standart := .9
  standart = math.Cbrt(standart / countOfStreams)
  // fmt.Println(standart)
  You.Health.Max = standart * 16
  empowering := ( - countOfStreams + You.Class )
  if empowering < 0 { empowering = 1 / (1 + math.Abs(empowering)) } else { empowering = (1 + math.Abs(empowering)) }
  empowering = math.Cbrt(empowering)
  if verbose {fmt.Printf("DEBUG [Player creation]: %d streams, %1.0f%% of power\n", int(countOfStreams), empowering*100)}
  // fmt.Printf("DEBUG [Player creation]:%s Count: %d, %1.2f%% of power\n", ebr, int(countOfStreams), empowering*100)
  stringsMatrix := []Stream {}
  lens, wids, pows, geomean := []float64 {}, []float64 {}, []float64 {}, 0.0
  for i:=0; i<int(countOfStreams); i++ {
    leni, widi, powi := 0.1+RNF(), 0.1+RNF(), 0.1+RNF()
    lens, wids, pows = append(lens, leni), append(wids, widi), append(pows, powi)
    geomean := math.Cbrt(leni*widi*powi)
    lens[i], wids[i], pows[i] = lens[i]*(standart/geomean), wids[i]*(standart/geomean), pows[i]*(standart/geomean)
    // fmt.Println(leni, widi, powi, "=>",geomean, "=>", lens[i], wids[i], pows[i], "=>", lens[i]*wids[i]*pows[i])
  }
  geomean = math.Pow(geomean, float64(1/countOfStreams/3))
  for i:=0; i<int(countOfStreams); i++ {
    var strings Stream
    strings.Element     = AllElements[0]
    strings.Creation    = lens[i] * empowering
    strings.Alteration  = wids[i]
    strings.Destruction = pows[i] / empowering
    stringsMatrix = append(stringsMatrix, strings)
    healthBonus := 0.0
    for _, each := range stringsMatrix { healthBonus += Log1479((each.Creation*each.Creation+1)*each.Alteration) }
    You.Health.Max += healthBonus
  }
  fmt.Println("DEBUG: Max health:", You.Health.Max)
  You.StreamStrings = stringsMatrix
  ExtendPools()
  InnerAffinization()
  ListStrings()
}
func ListStrings() {
  var counter Stream
  fmt.Printf("INFO [List strings]:")
  vols := 0.0
  counter.Creation, counter.Alteration, counter.Destruction = 0.0, 0.0, 0.0
  for i, stream := range You.StreamStrings {
    vols += (stream.Alteration)*(stream.Destruction)*(stream.Creation)
    counter.Creation += stream.Creation
    counter.Alteration += stream.Alteration
    counter.Destruction += stream.Destruction
    fmt.Printf("%s %d ─── %s %1.2f'len ─── %1.2f'wid ─── %1.2f'pow ────── Volume: %1.1f ─────", elbr,i+1, ES(stream.Element), stream.Creation, stream.Alteration, stream.Destruction, (stream.Alteration)*(stream.Destruction)*(stream.Creation))
  }
  if verbose == false {PlotEnvAff()}
  fmt.Printf("%s Total: %1.2f'lens + %1.2f'wids + %1.2f'pows = Volume: %1.1f\n", ebr, counter.Creation, counter.Alteration, counter.Destruction, vols)
}
func ExtendPools() {
  fmt.Printf("INFO [Extend dot capacity to maximum]:")
  old := You.Pool.MaxVol
  for _, stream := range You.StreamStrings {
    You.Pool.MaxVol += 32*math.Sqrt(1+stream.Creation)
  }
  You.Pool.MaxVol = math.Round(You.Pool.MaxVol)
  if verbose {
    fmt.Printf("\nDEBUG [Pool]: %1.0f'dots\n", You.Pool.MaxVol)
  } else {
    if old == 0 {old = You.Pool.MaxVol/2}
    fmt.Printf("%s INFO [Pool]: %+2.1f%%'dots\n", ebr, (You.Pool.MaxVol/old-1)*100)
  }
}
func EnergyStatus() {
  sum, mean := 0.0, 0.0
  fmt.Printf("\nINFO [List dots]:%s", elbr)
  count := 0
  span := int(math.Sqrt(2)*math.Sqrt( float64(len(You.Pool.Dots)+1) ))
  if span > 61 {span = 61}
  if verbose {span = 10}
  for e:=0; e<9; e++ {
    for _, dot := range You.Pool.Dots {
      if dot.Element == AllElements[e] {
        if (count)%span == 0 && count != 0 {
          fmt.Printf("%s",elbr)
        }
        if verbose {fmt.Printf("─%5.2f'%s ──", dot.Weight, ES(dot.Element))} else {fmt.Printf("%s",ES(dot.Element))}
        sum += dot.Weight
        mean += 1/dot.Weight
        count++
      }
    }
    // if verbose && count != 0 {fmt.Printf("%s", elbr)}
  }
  // if (count)%span == 0 {
  //   if count != len(You.Pool.Dots) {fmt.Printf("%s",elbr)}
  // }
  // if verbose != true {
  //   for e:=0; e<int(You.Pool.MaxVol)-len(You.Pool.Dots); e++ {
  //     if (count)%span == 0 && count != len(You.Pool.Dots) {
  //       fmt.Printf("%s",elbr)
  //     }
  //     fmt.Printf("◯ ")
  //     count++
  //   }
  // }
  fmt.Printf("\n")
  fmt.Printf("INFO [Energy status]:%s Total energy level: %2.1f%%", ebr, float64(len(You.Pool.Dots))/You.Pool.MaxVol*100)
  if verbose {fmt.Printf(" ─ mean:avg = %2.1f%%, %1.2f / %1.2f ─── Life: %2.1f%%", float64(len(You.Pool.Dots))/mean/(sum/float64(len(You.Pool.Dots)))*100, float64(len(You.Pool.Dots))/mean, sum/float64(len(You.Pool.Dots)), You.Health.Current/You.Health.Max*100)}
}

func GainDot() {
  if len(You.Pool.Dots) >= int(You.Pool.MaxVol) { time.Sleep( time.Millisecond * time.Duration( 4000 )) ; return }
  picker := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(You.StreamStrings))
  element := You.StreamStrings[picker].Element
  weight := Log1479( You.StreamStrings[picker].Alteration ) * (1 + RNF()) / 2
  dot := Dot{Element: element, Weight: weight}
  You.Pool.Dots = append(You.Pool.Dots, dot)
  if You.Health.Current < You.Health.Max {
    heal := math.Sqrt(weight)
    You.Health.Current += heal
    // if verbose {fmt.Printf("  %1.1f %1.1f  ", heal, weight)}
  } else { You.Health.Current = You.Health.Max }
  time.Sleep( time.Millisecond * time.Duration( 1000*Log1479(You.Pool.MaxVol)/math.Sqrt(You.Pool.MaxVol)*math.Sqrt(weight) ))
}
func Regenerate() {
  if len(You.Pool.Dots) >= int(You.Pool.MaxVol) {
    if verbose {fmt.Printf("\nDEBUG [regenerating]: nothing to regenerate. ")}
    time.Sleep( time.Millisecond * time.Duration( 4000 ))
    return
  }
  mana := int( math.Sqrt(You.Pool.MaxVol-float64(len(You.Pool.Dots))) )
  if verbose {fmt.Printf("\nDEBUG [regenerating]: +%d dots. ", mana)}
  for i:=0; i<mana; i++ {
    if len(You.Pool.Dots) >= int(You.Pool.MaxVol) {
      if verbose {fmt.Printf("\nDEBUG [regenerating]: nothing to regenerate. ")}
      time.Sleep( time.Millisecond * time.Duration( 4000 ))
      break
    }
    GainDot()
  }
}
func CrackStream(stream Stream) { // need heat {
  element := stream.Element
  weight := Log1479( stream.Destruction ) * (RNF() + RNF()) / 2
  dot := Dot{Element: element, Weight: weight}
  You.Pool.Dots = append(You.Pool.Dots, dot)
  // return heat[element] = sqrt(sqr(d+1)/sqr(l-1)/sqr(w-1)+1)
}
func EnergeticSurge(doze float64) { // need in time
  fmt.Printf("\n▲ YOU yelling [around]:%s CHEERS! A-ah...", ebr)
  if doze == 0 {
    doze = 1 / You.StreamStrings[0].Destruction
    for _, string := range You.StreamStrings { doze = math.Max(doze, 1 / string.Destruction) }
  }
  for _, string := range You.StreamStrings {
    i := 0.0
    for {
      CrackStream(string)
      i += 1 / doze
      if string.Destruction < i { break }
    }
  }
}
func MinusDot(index int) (string, float64) {
  if index >= len(You.Pool.Dots) { index = rand.New(rand.NewSource(time.Now().UnixNano())).Intn( len(You.Pool.Dots) ) }
  ddelement := You.Pool.Dots[index].Element
  ddweight := You.Pool.Dots[index].Weight
  You.Pool.Dots[index] = You.Pool.Dots[len(You.Pool.Dots)-1]
  You.Pool.Dots = You.Pool.Dots[:len(You.Pool.Dots)-1]
  return ddelement, ddweight
}
func DotTransferIn(e int) {
  if verbose {fmt.Printf("Absorbing dots:")}
  element := AllElements[e]
  if float64(len(You.Pool.Dots)) >= You.Pool.MaxVol+Log1479(You.Pool.MaxVol) { if verbose {fmt.Printf(" Full is energy.\n")} ; time.Sleep( time.Millisecond * time.Duration( 4000 )) ; return }
  weight := Log1479( You.ElemExt[e].Alteration ) * (1 + RNF()) / 2
  dot := Dot{Element: element, Weight: weight}
  You.Pool.Dots = append(You.Pool.Dots, dot)
  step := 32*math.Sqrt(1+math.Abs(You.ElemExt[e].Creation))
  if verbose {fmt.Printf(" +%s'%1.2f", ES(element), weight )}
  if verbose {fmt.Printf(", - dot absorbed for %1.3fs.\n", math.Log2(step)/math.Sqrt(step)*math.Sqrt(weight))}
  time.Sleep( time.Millisecond * time.Duration( 1000* math.Log2(step)/math.Sqrt(step) *math.Sqrt(weight) ))
}
func DotTransferOut(e int) {
  if verbose {fmt.Printf("Losing dots:")}
  element := AllElements[e]
  presense := []int{}
  for i, dot := range You.Pool.Dots { if dot.Element == element {presense = append(presense, i)} }
  if len(presense) == 0 { if verbose{fmt.Printf(" No such dots.\n")} ; time.Sleep( time.Millisecond * time.Duration( 4000 )) ; return }
  killer := presense[rand.New(rand.NewSource(time.Now().UnixNano())).Intn( len(presense) )]
  _, weight := MinusDot(killer)
  step := 32*math.Sqrt(1+math.Abs(You.ElemExt[e].Destruction)) //Log1479(math.Abs(You.ElemExt[e].Destruction)) * (1 + RNF()) / 2
  if verbose {fmt.Printf(" -%s'%1.2f", ES(element), weight)}
  if verbose {fmt.Printf(", - dot is lost for %1.3fs.\n", math.Log2(step)/math.Sqrt(step)/math.Sqrt(weight))}
  time.Sleep( time.Millisecond * time.Duration( 1000* math.Log2(step)/math.Sqrt(step) /math.Sqrt(weight) ))
}
func Transferrence() {
  demand := [9]int{}
  cooldown := 0.0
  for i, source := range You.ElemExt {
    count := 0.0
    if source.Creation < 0 { count = - math.Sqrt(1+math.Abs(source.Destruction)) * (1 + RNF()) / 2 } else { count = math.Sqrt(1+math.Abs(source.Creation)) * (1 + RNF()) / 2 }
    if i == 0 { count = 0 }
    demand[i] = ChancedRound(count * Sign(You.ElemEnv[i].Creation))
    cooldown = math.Max(math.Abs(count) * 500, cooldown)
  }
  if cooldown == 0 { cooldown = 2000 }
  if verbose {fmt.Printf("\nDEBUG [transferrence]: %v dots, cooldown: %1.3fs \n", demand, cooldown/1000)}
  wg := sync.WaitGroup{}
  for e, _ := range demand {
    amount := demand[e]
    if demand[e] > 0 {
      if verbose {fmt.Println("Gaining", amount, AllElements[e])}
      wg.Add(1)
      go func(e int){
        defer wg.Done()
        for j:=0; j<amount; j++ { DotTransferIn(e) }
      }(e)
    } else if demand[e] < 0 {
      amount = 0 - demand[e]
      if verbose {fmt.Println("Loosing", amount, AllElements[e])}
      wg.Add(1)
      go func(e int){
        defer wg.Done()
        for j:=0; j<amount; j++ { DotTransferOut(e) }
      }(e)
    }
  }
  time.Sleep( time.Millisecond * time.Duration( cooldown ))
  wg.Wait()
}
