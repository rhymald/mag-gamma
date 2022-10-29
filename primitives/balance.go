package primitives
import "math"

//// Player stats
// streams basics
func BendHeraldFromClass(class float64) (float64, float64) { return math.Cbrt(7.5-class), math.Cbrt(0.5+class) }
func LenFromStream(stream Stream) float64 { return stream.Creation*1024 }
func WidFromStream(stream Stream) float64 { return 32*stream.Alteration/Vector(stream.Alteration, stream.Creation) }
func PowFromStream(stream Stream) float64 { return 10*stream.Destruction/Vector(stream.Destruction, stream.Alteration, stream.Creation) }
// elemental state
func ResistanceFromState(state Stream) float64 { return Vector(state.Destruction,state.Destruction,state.Creation) }
