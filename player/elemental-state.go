package player

import "rhymald/mag-gamma/primitives"
import "rhymald/mag-gamma/environment"
import "fmt"
import "math"

//Player things
type ElementalState struct {
  External    [9]primitives.Stream
  Empowered   [9]primitives.Stream
  // ^ both for gains from env
  Internal    [9]primitives.Stream
  Deminished  [9]primitives.Stream
  Resistances [9]float64
  // ^ two for resistance supression
  Total [9]primitives.Stream
  // ^ for casts
}

func PlotElementalState(estate ElementalState, verbose bool) {
  i := 0
  if verbose {
    fmt.Println()
    fmt.Printf(" ┌──── DEBUG [fundamental attribute stats]:")
    if estate.Empowered[0].Creation != 0 {fmt.Printf(" ✳️  ──  %1.2f'vitality ── %1.2f'concentration ── %1.2f'power", estate.Empowered[0].Creation, estate.Empowered[0].Alteration, estate.Empowered[0].Destruction)}
  } else {
    fmt.Println()
    fmt.Printf(" ┌──── INFO [fundamental attribute balance]:")
    summ := estate.Empowered[0].Creation+estate.Empowered[0].Alteration+estate.Empowered[0].Destruction
    if estate.Empowered[0].Creation != 0 {fmt.Printf(" ✳️  ──  %2.1f%%'vitality ── %2.1f%%'concentration ── %2.1f%%'power", estate.Empowered[0].Creation/summ*100, estate.Empowered[0].Alteration/summ*100, estate.Empowered[0].Destruction/summ*100)}
  }
  if verbose {
    fmt.Println()
    fmt.Printf(" ├── DEBUG [surrending elemental state]: ")
    if estate.External[1].Creation != 0 {fmt.Printf("\n │ ☁️  ── %1.2f'pressure ── %1.2f'spreading ── %1.2f'puncture ", estate.External[1].Creation, estate.External[1].Alteration, estate.External[1].Destruction) ; i++}
    if estate.External[2].Creation != 0 {fmt.Printf("\n │ 🔥 ── %1.2f'warming ── %1.2f'burning ── %1.2f'detonation ", estate.External[2].Creation, estate.External[2].Alteration, estate.External[2].Destruction) ; i++}
    if estate.External[3].Creation != 0 {fmt.Printf("\n │ ⛰  ── %1.2f'structure ── %1.2f'mass ── %1.2f'fragility ", estate.External[3].Creation, estate.External[3].Alteration, estate.External[3].Destruction) ; i++}
    if estate.External[4].Creation != 0 {fmt.Printf("\n │ 🧊 ── %1.2f'form ── %1.2f'direction ── %1.2f'calm ", estate.External[4].Creation, estate.External[4].Alteration, estate.External[4].Destruction) ; i++}
    if estate.External[5].Creation != 0 {fmt.Printf("\n │ 🌑 ── %1.2f'shadows ── %1.2f'curse ── %1.2f'pain ", estate.External[5].Creation, estate.External[5].Alteration, estate.External[5].Destruction) ; i++}
    if estate.External[6].Creation != 0 {fmt.Printf("\n │ 🩸 ── %1.2f'growing ── %1.2f'corruption ── %1.2f'consumption ", estate.External[6].Creation, estate.External[6].Alteration, estate.External[6].Destruction) ; i++}
    if estate.External[7].Creation != 0 {fmt.Printf("\n │ 🎶 ── %1.2f'inspiration ── %1.2f'echo ── %1.2f'roar ", estate.External[7].Creation, estate.External[7].Alteration, estate.External[7].Destruction) ; i++}
    if estate.External[8].Creation != 0 {fmt.Printf("\n │ 🌟 ── %1.2f'mirage ── %1.2f'matter ── %1.2f'desintegration ", estate.External[8].Creation, estate.External[8].Alteration, estate.External[8].Destruction) ; i++}
    if i==0 {fmt.Printf("Totaly not affected by environment ")}
    fmt.Println()
    fmt.Printf(" ├── DEBUG [incoming elemental affection]: ")
    if estate.Empowered[1].Creation != 0 {fmt.Printf("\n │ ☁️  ── %1.2f'pressure ── %1.2f'spreading ── %1.2f'puncture ", estate.Empowered[1].Creation, estate.Empowered[1].Alteration, estate.Empowered[1].Destruction) }
    if estate.Empowered[2].Creation != 0 {fmt.Printf("\n │ 🔥 ── %1.2f'warming ── %1.2f'burning ── %1.2f'detonation ", estate.Empowered[2].Creation, estate.Empowered[2].Alteration, estate.Empowered[2].Destruction) }
    if estate.Empowered[3].Creation != 0 {fmt.Printf("\n │ ⛰  ── %1.2f'structure ── %1.2f'mass ── %1.2f'fragility ", estate.Empowered[3].Creation, estate.Empowered[3].Alteration, estate.Empowered[3].Destruction) }
    if estate.Empowered[4].Creation != 0 {fmt.Printf("\n │ 🧊 ── %1.2f'form ── %1.2f'direction ── %1.2f'calm ", estate.Empowered[4].Creation, estate.Empowered[4].Alteration, estate.Empowered[4].Destruction) }
    if estate.Empowered[5].Creation != 0 {fmt.Printf("\n │ 🌑 ── %1.2f'shadows ── %1.2f'curse ── %1.2f'pain ", estate.Empowered[5].Creation, estate.Empowered[5].Alteration, estate.Empowered[5].Destruction) }
    if estate.Empowered[6].Creation != 0 {fmt.Printf("\n │ 🩸 ── %1.2f'growing ── %1.2f'corruption ── %1.2f'consumption ", estate.Empowered[6].Creation, estate.Empowered[6].Alteration, estate.Empowered[6].Destruction) }
    if estate.Empowered[7].Creation != 0 {fmt.Printf("\n │ 🎶 ── %1.2f'inspiration ── %1.2f'echo ── %1.2f'roar ", estate.Empowered[7].Creation, estate.Empowered[7].Alteration, estate.Empowered[7].Destruction) }
    if estate.Empowered[8].Creation != 0 {fmt.Printf("\n │ 🌟 ── %1.2f'mirage ── %1.2f'matter ── %1.2f'desintegration ", estate.Empowered[8].Creation, estate.Empowered[8].Alteration, estate.Empowered[8].Destruction) }
    if i==0 {fmt.Printf("Totaly not affected by environment ")}
    // fmt.Printf("\n │DEBUG [finaly affecting elemental state]: ") -- dat will be curses
    // if You.ElemAff[1].Creation != 0 {fmt.Printf("\n │ ☁️   %1.2f'pressure  %1.2f'spreading  %1.2f'puncture ", You.ElemAff[1].Creation, You.ElemAff[1].Alteration, You.ElemAff[1].Destruction) }
    // if You.ElemAff[2].Creation != 0 {fmt.Printf("\n │ 🔥  %1.2f'warming  %1.2f'burning  %1.2f'detonation ", You.ElemAff[2].Creation, You.ElemAff[2].Alteration, You.ElemAff[2].Destruction) }
    // if You.ElemAff[3].Creation != 0 {fmt.Printf("\n │ ⛰   %1.2f'structure  %1.2f'mass  %1.2f'fragility ", You.ElemAff[3].Creation, You.ElemAff[3].Alteration, You.ElemAff[3].Destruction) }
    // if You.ElemAff[4].Creation != 0 {fmt.Printf("\n │ 🧊  %1.2f'form  %1.2f'direction  %1.2f'calm ", You.ElemAff[4].Creation, You.ElemAff[4].Alteration, You.ElemAff[4].Destruction) }
    // if You.ElemAff[5].Creation != 0 {fmt.Printf("\n │ 🌑  %1.2f'shadows  %1.2f'curse  %1.2f'pain ", You.ElemAff[5].Creation, You.ElemAff[5].Alteration, You.ElemAff[5].Destruction) }
    // if You.ElemAff[6].Creation != 0 {fmt.Printf("\n │ 🩸  %1.2f'growing  %1.2f'corruption  %1.2f'consumption ", You.ElemAff[6].Creation, You.ElemAff[6].Alteration, You.ElemAff[6].Destruction) }
    // if You.ElemAff[7].Creation != 0 {fmt.Printf("\n │ 🎶  %1.2f'inspiration  %1.2f'echo  %1.2f'roar ", You.ElemAff[7].Creation, You.ElemAff[7].Alteration, You.ElemAff[7].Destruction) }
    // if You.ElemAff[8].Creation != 0 {fmt.Printf("\n │ 🌟  %1.2f'mirage  %1.2f'matter  %1.2f'desintegration ", You.ElemAff[8].Creation, You.ElemAff[8].Alteration, You.ElemAff[8].Destruction) }
    // if i==0 {fmt.Printf("\n │ Totaly not affected by environment ", elbr)}
    fmt.Println()
    fmt.Printf(" ├── DEBUG [internal elemental state]: ")
    if estate.Internal[1].Creation != 0 {fmt.Printf("\n │ ☁️  ── %1.2f'pressure ── %1.2f'spreading ── %1.2f'puncture ", estate.Internal[1].Creation, estate.Internal[1].Alteration, estate.Internal[1].Destruction) }
    if estate.Internal[2].Creation != 0 {fmt.Printf("\n │ 🔥 ── %1.2f'warming ── %1.2f'burning ── %1.2f'detonation ", estate.Internal[2].Creation, estate.Internal[2].Alteration, estate.Internal[2].Destruction) }
    if estate.Internal[3].Creation != 0 {fmt.Printf("\n │ ⛰  ── %1.2f'structure ── %1.2f'mass ── %1.2f'fragility ", estate.Internal[3].Creation, estate.Internal[3].Alteration, estate.Internal[3].Destruction) }
    if estate.Internal[4].Creation != 0 {fmt.Printf("\n │ 🧊 ── %1.2f'form ── %1.2f'direction ── %1.2f'calm ", estate.Internal[4].Creation, estate.Internal[4].Alteration, estate.Internal[4].Destruction) }
    if estate.Internal[5].Creation != 0 {fmt.Printf("\n │ 🌑 ── %1.2f'shadows ── %1.2f'curse ── %1.2f'pain ", estate.Internal[5].Creation, estate.Internal[5].Alteration, estate.Internal[5].Destruction) }
    if estate.Internal[6].Creation != 0 {fmt.Printf("\n │ 🩸 ── %1.2f'growing ── %1.2f'corruption ── %1.2f'consumption ", estate.Internal[6].Creation, estate.Internal[6].Alteration, estate.Internal[6].Destruction) }
    if estate.Internal[7].Creation != 0 {fmt.Printf("\n │ 🎶 ── %1.2f'inspiration ── %1.2f'echo ── %1.2f'roar ", estate.Internal[7].Creation, estate.Internal[7].Alteration, estate.Internal[7].Destruction) }
    if estate.Internal[8].Creation != 0 {fmt.Printf("\n │ 🌟 ── %1.2f'mirage ── %1.2f'matter ── %1.2f'desintegration ", estate.Internal[8].Creation, estate.Internal[8].Alteration, estate.Internal[8].Destruction) }
  }
  fmt.Println()
  fmt.Printf(" ├── INFO [resistances]: ")
  fmt.Printf("\n │ %s:%1.2f   ──  ", ElemSigns[0], math.Sqrt( math.Pow(estate.Internal[0].Creation, 2) + math.Pow(estate.Internal[0].Alteration, 2) + math.Pow(estate.Internal[0].Destruction, 2)) )
  for i:=0; i<8; i++ { if estate.Resistances[i] != 0 { fmt.Printf(" %s:%1.2f  ", ElemSigns[i+1], estate.Resistances[i] ) } }
  fmt.Println()
  fmt.Printf(" └────────────────────────────────────────────────────────────────────────────────────────────────────\n")
}

func ReadStatesFromEnv(elementalState *ElementalState, position [3]float64, flock Streams, location environment.Location) {
  var affectingPlaces []primitives.Stream
  for _, place := range location.Wells {
    for _, being := range place.XYZs {
      distance := math.Sqrt(math.Pow(position[0]-being[0],2)+math.Pow(position[1]-being[1],2)+math.Pow(position[2]-being[2],2))/place.Area
      if distance <= 1 {
        for _, affection := range place.Nature {
          if place.Concentrated {
            buffer := primitives.Stream{
              Element: affection.Element,
              Creation: affection.Creation * math.Pow(1-distance, 2), // creation amount
              Alteration: affection.Alteration * math.Pow(1-distance, 2), //creation quality
              Destruction: affection.Destruction * math.Pow(1-distance, 2), // loose amount
            }
            affectingPlaces = append(affectingPlaces, buffer)
          } else {
            buffer := primitives.Stream{
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
  estate := ElementalState{}
  // for _, affection := range affectingPlaces {
  //   for i:=1;i<9;i++ {
  //     if affection.Element == AllElements[i] {
  //       estate.External[i].Creation += affection.Creation
  //       estate.External[i].Alteration += affection.Alteration
  //       estate.External[i].Destruction += affection.Destruction
  //     }
  //   }
  // }
  for i, element := range AllElements {
    estate.Internal[i] = primitives.StreamSum(element, flock.List)
    if i!=0 {estate.External[i] = primitives.StreamSum(element, affectingPlaces)}
  }
  // for _, each := range flock.List {
  //   estate.Internal[0].Creation += each.Creation // + estate.Empowered[0].Creation)
  //   estate.Internal[0].Alteration += each.Alteration //  + estate.Empowered[0].Alteration)
  //   estate.Internal[0].Destruction += each.Destruction //    + estate.Empowered[0].Destruction)
  //   if each.Element != "Common" {
  //     estate.Internal[primitives.ElemToInt(each.Element)].Creation += each.Creation // + estate.Empowered[0].Creation)
  //     estate.Internal[primitives.ElemToInt(each.Element)].Alteration += each.Alteration //  + estate.Empowered[0].Alteration)
  //     estate.Internal[primitives.ElemToInt(each.Element)].Destruction += each.Destruction //    + estate.Empowered[0].Destruction)
  //   }
  // }
  *elementalState = estate
}

func InnerAffinization(elementalState *ElementalState, bender float64, herald float64) {
  estate := *elementalState
  // External basic
  estate.Empowered[1].Creation    = bender * ( estate.External[1].Creation    + estate.External[3].Destruction) - estate.External[2].Destruction * herald
  estate.Empowered[1].Alteration  = bender * ( estate.External[1].Alteration  + estate.External[3].Alteration)  - estate.External[2].Alteration  * herald
  estate.Empowered[1].Destruction = bender * ( estate.External[1].Destruction + estate.External[3].Creation)    - estate.External[2].Creation    * herald
  estate.Empowered[2].Creation    = bender * ( estate.External[2].Creation    + estate.External[1].Destruction) - estate.External[4].Destruction * herald
  estate.Empowered[2].Alteration  = bender * ( estate.External[2].Alteration  + estate.External[1].Alteration)  - estate.External[4].Alteration  * herald
  estate.Empowered[2].Destruction = bender * ( estate.External[2].Destruction + estate.External[1].Creation)    - estate.External[4].Creation    * herald
  estate.Empowered[3].Creation    = bender * ( estate.External[3].Creation    + estate.External[4].Destruction) - estate.External[1].Destruction * herald
  estate.Empowered[3].Alteration  = bender * ( estate.External[3].Alteration  + estate.External[4].Alteration)  - estate.External[1].Alteration  * herald
  estate.Empowered[3].Destruction = bender * ( estate.External[3].Destruction + estate.External[4].Creation)    - estate.External[1].Creation    * herald
  estate.Empowered[4].Creation    = bender * ( estate.External[4].Creation    + estate.External[2].Destruction) - estate.External[3].Destruction * herald
  estate.Empowered[4].Alteration  = bender * ( estate.External[4].Alteration  + estate.External[2].Alteration)  - estate.External[3].Alteration  * herald
  estate.Empowered[4].Destruction = bender * ( estate.External[4].Destruction + estate.External[2].Creation)    - estate.External[3].Creation    * herald
  // v void - extra consumption
  estate.Empowered[5].Creation    = bender * ( estate.External[5].Creation) //    + estate.Empowered[5].Creation)
  estate.Empowered[5].Alteration  = bender * ( estate.External[5].Alteration) //  + estate.Empowered[5].Alteration)
  estate.Empowered[5].Destruction = bender * ( estate.External[5].Destruction) // + estate.Empowered[5].Destruction)
  // v deviant - extra overheat
  estate.Empowered[6].Creation    = bender * ( estate.External[6].Creation    + 2 * math.Sqrt(estate.External[4].Creation    * estate.External[2].Creation)) // + estate.Empowered[6].Creation    )
  estate.Empowered[6].Alteration  = bender * ( estate.External[6].Alteration  + 2 * math.Sqrt(estate.External[4].Alteration  * estate.External[2].Creation)) // + estate.Empowered[6].Alteration  )
  estate.Empowered[6].Destruction = bender * ( estate.External[6].Destruction + 2 * math.Sqrt(estate.External[4].Destruction * estate.External[2].Creation)) // + estate.Empowered[6].Destruction )
  estate.Empowered[7].Creation    = bender * ( estate.External[7].Creation    + 2 * math.Sqrt(estate.External[3].Creation    * estate.External[1].Creation)) //  + estate.Empowered[7].Creation    )
  estate.Empowered[7].Alteration  = bender * ( estate.External[7].Alteration  + 2 * math.Sqrt(estate.External[3].Alteration  * estate.External[1].Creation)) //  + estate.Empowered[7].Alteration  )
  estate.Empowered[7].Destruction = bender * ( estate.External[7].Destruction + 2 * math.Sqrt(estate.External[3].Destruction * estate.External[1].Creation)) //  + estate.Empowered[7].Destruction )
  // v rarest - extra overheat and consumption
  estate.Empowered[8].Creation    = bender * ( estate.External[8].Alteration + estate.External[8].Creation   )//  + estate.Empowered[8].Creation)
  estate.Empowered[8].Alteration  = bender * ( estate.External[8].Creation   + estate.External[8].Destruction)//  + estate.Empowered[8].Alteration)
  estate.Empowered[8].Destruction = bender * ( estate.External[8].Alteration + estate.External[8].Destruction) // + estate.Empowered[8].Destruction)
  // Finalizing
  estate.Empowered[0].Creation = bender * estate.External[8].Creation - estate.External[5].Destruction * herald + estate.Internal[0].Creation
  estate.Empowered[0].Alteration = bender * estate.External[8].Alteration - estate.External[5].Alteration * herald + estate.Internal[0].Alteration
  estate.Empowered[0].Destruction = bender * estate.External[8].Destruction - estate.External[5].Creation * herald + estate.Internal[0].Destruction
  for i:=1; i<9; i++ { if estate.Internal[i].Creation*estate.Internal[i].Destruction > 0 { estate.Resistances[i-1] = primitives.ResistanceFromState(estate.Internal[i]) }} // math.Sqrt( math.Pow(estate.Internal[i].Destruction, 2)*2 - 1 + math.Pow(estate.Internal[i].Creation, 2)) }}
  *elementalState = estate
}
