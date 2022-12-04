package player

import "rhymald/mag-gamma/primitives"
import "rhymald/mag-gamma/environment"
import "fmt"
import "math"

//Player things
type ElementalState struct {
  ExternalWells    [9]primitives.Stream // for transference
  Empowered   [9]primitives.Stream // for transfer also
  // ^ both for gains from env
  // Internal    [9]primitives.Stream // for overheat
  ExternalCurses    [9]primitives.Stream // for curses
  Deminished  [9]primitives.Stream // for curses
  Resistances [9]float64 // TBD move to health/def
  // ^ two for resistance supression
}

func PlotElementalState(estate [9]primitives.Stream, header string, verbose bool) {
  i := 0
  if verbose {
    fmt.Println()
    fmt.Printf(" ┌──── DEBUG [Plot elemental state][%s]:", header)
    if estate[0].Creation != 0 {fmt.Printf("\n │ ✳️  ── %1.2f'vitality ── %1.2f'concentration ── %1.2f'power", estate[0].Creation, estate[0].Alteration, estate[0].Destruction); fmt.Printf("\n ├── DEBUG [elemental state]: ")}
  } else {
    fmt.Println()
    fmt.Printf(" ┌──── INFO [%s]:", header)
    summ := estate[0].Creation+estate[0].Alteration+estate[0].Destruction
    if estate[0].Creation != 0 {fmt.Printf("\n │ ✳️  ── %2.1f%%'vitality ── %2.1f%%'concentration ── %2.1f%%'power", estate[0].Creation/summ*100, estate[0].Alteration/summ*100, estate[0].Destruction/summ*100); fmt.Printf("\n ├── DEBUG [elemental state]: ")}
  }
  if verbose {
    if estate[1].Creation != 0 {fmt.Printf("\n │ ☁️  ── %1.2f'pressure ── %1.2f'spreading ── %1.2f'penetration ", estate[1].Creation, estate[1].Alteration, estate[1].Destruction) ; i++}
    if estate[2].Creation != 0 {fmt.Printf("\n │ 🔥 ── %1.2f'warming ── %1.2f'burning ── %1.2f'detonation ", estate[2].Creation, estate[2].Alteration, estate[2].Destruction) ; i++}
    if estate[3].Creation != 0 {fmt.Printf("\n │ ⛰  ── %1.2f'structure ── %1.2f'mass ── %1.2f'fragility ", estate[3].Creation, estate[3].Alteration, estate[3].Destruction) ; i++}
    if estate[4].Creation != 0 {fmt.Printf("\n │ 🧊 ── %1.2f'form ── %1.2f'direction ── %1.2f'calm ", estate[4].Creation, estate[4].Alteration, estate[4].Destruction) ; i++}
    if estate[5].Creation != 0 {fmt.Printf("\n ├────────────────────────────────────────────────────────────────────────────────────────────────────")}
    if estate[5].Creation != 0 {fmt.Printf("\n │ 🌑 ── %1.2f'shadows ── %1.2f'curse ── %1.2f'pain ", estate[5].Creation, estate[5].Alteration, estate[5].Destruction) ; i++}
    if estate[6].Creation != 0 || estate[7].Creation != 0 {fmt.Printf("\n ├────────────────────────────────────────────────────────────────────────────────────────────────────")}
    if estate[6].Creation != 0 {fmt.Printf("\n │ 🩸 ── %1.2f'growing ── %1.2f'corruption ── %1.2f'consumption ", estate[6].Creation, estate[6].Alteration, estate[6].Destruction) ; i++}
    if estate[7].Creation != 0 {fmt.Printf("\n │ 🎶 ── %1.2f'inspiration ── %1.2f'echo ── %1.2f'roar ", estate[7].Creation, estate[7].Alteration, estate[7].Destruction) ; i++}
    if estate[8].Creation != 0 {fmt.Printf("\n ├────────────────────────────────────────────────────────────────────────────────────────────────────")}
    if estate[8].Creation != 0 {fmt.Printf("\n │ 🌟 ── %1.2f'mirage ── %1.2f'matter ── %1.2f'desintegration ", estate[8].Creation, estate[8].Alteration, estate[8].Destruction) ; i++}
    if i==0 {fmt.Printf("Totaly not attuned or affected by environment.")}
  }
  fmt.Println()
  fmt.Printf(" └────────────────────────────────────────────────────────────────────────────────────────────────────")
}

func ReadStatesFromEnv(elementalState *ElementalState, position [3]float64, flock *Streams, location environment.Location) {
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
  for i, element := range AllElements {
    if i!=0 {
      *&flock.InternalElementalState[i] = primitives.StreamSum(element, *&flock.List)
      estate.ExternalWells[i] = primitives.StreamSum(element, affectingPlaces)
    } else {
      *&flock.InternalElementalState[i] = primitives.StreamSum(element, *&flock.List)
    }
  }
  *elementalState = estate
}

func InnerAffinization(elementalState *ElementalState, bender float64, herald float64) {
  estate := *elementalState
  // ExternalWells basic
  estate.Empowered[1].Creation    = bender * ( estate.ExternalWells[1].Creation    + estate.ExternalWells[3].Destruction) - estate.ExternalWells[2].Destruction * herald
  estate.Empowered[1].Alteration  = bender * ( estate.ExternalWells[1].Alteration  + estate.ExternalWells[3].Alteration)  - estate.ExternalWells[2].Alteration  * herald
  estate.Empowered[1].Destruction = bender * ( estate.ExternalWells[1].Destruction + estate.ExternalWells[3].Creation)    - estate.ExternalWells[2].Creation    * herald
  estate.Empowered[2].Creation    = bender * ( estate.ExternalWells[2].Creation    + estate.ExternalWells[1].Destruction) - estate.ExternalWells[4].Destruction * herald
  estate.Empowered[2].Alteration  = bender * ( estate.ExternalWells[2].Alteration  + estate.ExternalWells[1].Alteration)  - estate.ExternalWells[4].Alteration  * herald
  estate.Empowered[2].Destruction = bender * ( estate.ExternalWells[2].Destruction + estate.ExternalWells[1].Creation)    - estate.ExternalWells[4].Creation    * herald
  estate.Empowered[3].Creation    = bender * ( estate.ExternalWells[3].Creation    + estate.ExternalWells[4].Destruction) - estate.ExternalWells[1].Destruction * herald
  estate.Empowered[3].Alteration  = bender * ( estate.ExternalWells[3].Alteration  + estate.ExternalWells[4].Alteration)  - estate.ExternalWells[1].Alteration  * herald
  estate.Empowered[3].Destruction = bender * ( estate.ExternalWells[3].Destruction + estate.ExternalWells[4].Creation)    - estate.ExternalWells[1].Creation    * herald
  estate.Empowered[4].Creation    = bender * ( estate.ExternalWells[4].Creation    + estate.ExternalWells[2].Destruction) - estate.ExternalWells[3].Destruction * herald
  estate.Empowered[4].Alteration  = bender * ( estate.ExternalWells[4].Alteration  + estate.ExternalWells[2].Alteration)  - estate.ExternalWells[3].Alteration  * herald
  estate.Empowered[4].Destruction = bender * ( estate.ExternalWells[4].Destruction + estate.ExternalWells[2].Creation)    - estate.ExternalWells[3].Creation    * herald
  // v void - extra consumption
  estate.Empowered[5].Creation    = bender * ( estate.ExternalWells[5].Creation) //    + estate.Empowered[5].Creation)
  estate.Empowered[5].Alteration  = bender * ( estate.ExternalWells[5].Alteration) //  + estate.Empowered[5].Alteration)
  estate.Empowered[5].Destruction = bender * ( estate.ExternalWells[5].Destruction) // + estate.Empowered[5].Destruction)
  // v deviant - extra overheat
  estate.Empowered[6].Creation    = bender * ( estate.ExternalWells[6].Creation    + 2 * math.Sqrt(estate.ExternalWells[4].Creation    * estate.ExternalWells[2].Creation)) // + estate.Empowered[6].Creation    )
  estate.Empowered[6].Alteration  = bender * ( estate.ExternalWells[6].Alteration  + 2 * math.Sqrt(estate.ExternalWells[4].Alteration  * estate.ExternalWells[2].Creation)) // + estate.Empowered[6].Alteration  )
  estate.Empowered[6].Destruction = bender * ( estate.ExternalWells[6].Destruction + 2 * math.Sqrt(estate.ExternalWells[4].Destruction * estate.ExternalWells[2].Creation)) // + estate.Empowered[6].Destruction )
  estate.Empowered[7].Creation    = bender * ( estate.ExternalWells[7].Creation    + 2 * math.Sqrt(estate.ExternalWells[3].Creation    * estate.ExternalWells[1].Creation)) //  + estate.Empowered[7].Creation    )
  estate.Empowered[7].Alteration  = bender * ( estate.ExternalWells[7].Alteration  + 2 * math.Sqrt(estate.ExternalWells[3].Alteration  * estate.ExternalWells[1].Creation)) //  + estate.Empowered[7].Alteration  )
  estate.Empowered[7].Destruction = bender * ( estate.ExternalWells[7].Destruction + 2 * math.Sqrt(estate.ExternalWells[3].Destruction * estate.ExternalWells[1].Creation)) //  + estate.Empowered[7].Destruction )
  // v rarest - extra overheat and consumption
  estate.Empowered[8].Creation    = bender * ( estate.ExternalWells[8].Alteration + estate.ExternalWells[8].Creation   )//  + estate.Empowered[8].Creation)
  estate.Empowered[8].Alteration  = bender * ( estate.ExternalWells[8].Creation   + estate.ExternalWells[8].Destruction)//  + estate.Empowered[8].Alteration)
  estate.Empowered[8].Destruction = bender * ( estate.ExternalWells[8].Alteration + estate.ExternalWells[8].Destruction) // + estate.Empowered[8].Destruction)
  // Finalizing
  // estate.Empowered[0].Creation = bender * estate.ExternalWells[8].Creation - estate.ExternalWells[5].Destruction * herald //+ estate.Internal[0].Creation
  // estate.Empowered[0].Alteration = bender * estate.ExternalWells[8].Alteration - estate.ExternalWells[5].Alteration * herald //+ estate.Internal[0].Alteration
  // estate.Empowered[0].Destruction = bender * estate.ExternalWells[8].Destruction - estate.ExternalWells[5].Creation * herald //+ estate.Internal[0].Destruction
  // for i:=1; i<9; i++ { if estate.Internal[i].Creation*estate.Internal[i].Destruction > 0 { estate.Resistances[i-1] = primitives.InnerAffinization_ResistanceFromState(estate.Internal[i]) }} // math.Sqrt( math.Pow(estate.Internal[i].Destruction, 2)*2 - 1 + math.Pow(estate.Internal[i].Creation, 2)) }}
  *elementalState = estate
}
