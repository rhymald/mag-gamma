package player

import "rhymald/mag-gamma/primitives"

//Player things
type ElementalState struct {
  External   [9]primitives.Stream
  Empowered  [9]primitives.Stream
  // ^ both for gains from env
  Internal   [9]primitives.Stream
  Deminished [9]primitives.Stream
  // ^ two for resistance supression
  Total [9]primitives.Stream
  // ^ for casts
}
