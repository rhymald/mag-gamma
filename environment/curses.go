package environment

import "rhymald/mag-gamma/primitives"
import "fmt"

// emulation, in future get from db
func Cursing(environment *Location) {
  fmt.Printf(" ┌─ DEBUG [Cursing]: reading bunch of negative power states - start.\n")
  // state where fires are much dangerous
  driesFire := primitives.Stream{Element: "Fire", Creation: 5, Destruction: 8, Alteration: 3}
  Dries     := PowerState{ Area: 22.0, Nature: []primitives.Stream{driesFire}, Description: "Dry fields: dried plants.", Concentrated: false}
  Dries.XYZs = append(Dries.XYZs, [3]float64{1000, 0, 430})
  // compose
  var buffer []PowerState
  buffer = append(buffer, Dries)
  *&environment.Curses = buffer
  fmt.Printf(" └─ DEBUG [Cursing]: reading bunch of negative power states - done.\n")
}
