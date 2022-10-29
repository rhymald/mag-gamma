package primitives

//// Player stats
// streams basics
func LenFromStream(stream Stream) float64 { return stream.Creation*1024 }
func WidFromStream(stream Stream) float64 { return 32*stream.Alteration/Vector(stream.Alteration, stream.Creation) }
func PowFromStream(stream Stream) float64 { return 10*stream.Destruction/Vector(stream.Destruction, stream.Alteration, stream.Creation) }
// elemental state
func ResistanceFromState(state Stream) float64 { return Vector(state.Destruction,state.Destruction,state.Creation) }
