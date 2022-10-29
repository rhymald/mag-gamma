package environment

import "rhymald/mag-gamma/primitives"
import "fmt"
import "math"


// emulation, in future get from db
func Welling(environment *Location) {
  fmt.Printf(" ┌──── DEBUG [Location init][Welling]: reading bunch of positive power states - start.\n │ ")
  // campfires
  campFireFire := primitives.Stream{Element: "Fire", Creation: 3, Destruction: 5, Alteration: 8}
  CampFire     := PowerState{ Area: 5.0, Nature: []primitives.Stream{campFireFire}, Description: "Campfire: Warm place to rest at.", Concentrated: false}
  CampFire.XYZs = append(CampFire.XYZs, [3]float64{1014.0, -16.5, 430.0})
  // elementalTree
  powerCore1 := primitives.Stream{Element: primitives.RNDElem(), Creation: 1+math.Abs(primitives.SRNF()), Destruction: 1+math.Abs(primitives.SRNF()), Alteration: 1+math.Abs(primitives.SRNF())}
  powerCore2 := primitives.Stream{Element: primitives.RNDElem(), Creation: 1+math.Abs(primitives.SRNF()), Destruction: 1+math.Abs(primitives.SRNF()), Alteration: 1+math.Abs(primitives.SRNF())}
  PowerCore     := PowerState{ Area: 25.0, Nature: []primitives.Stream{powerCore1, powerCore2}, Description: "Source of pure energy.", Concentrated: true}
  PowerCore.XYZs = append(PowerCore.XYZs, [3]float64{1011.1, -8.5, 430.0})
  fmt.Printf("DEBUG [Location init][Welling]: power core is randomized for emitting %s and %s energy...\n", primitives.ES(powerCore1.Element), primitives.ES(powerCore2.Element))
  // windySpaces
  openSpaceAir  := primitives.Stream{Element: "Air", Creation: 1, Destruction: 2, Alteration: 3}
  OpenSpace     := PowerState{ Area: 931.0, Nature: []primitives.Stream{openSpaceAir}, Description: "Winmdy weather.", Concentrated: false}
  OpenSpace.XYZs = append(OpenSpace.XYZs, [3]float64{600.0, 100.5, 700.0})
  // compose
  var buffer []PowerState
  buffer = append(buffer, CampFire,OpenSpace,PowerCore)
  *&environment.Wells = buffer
  fmt.Printf(" │ DEBUG [Location init][Welling]: reading bucnh of positive power states - done.\n")
  fmt.Printf(" └────────────────────────────────────────────────────────────────────────────────────────────────────\n")
}
