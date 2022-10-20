package primitives

import (
  "time"
  "crypto/sha512"
  "math/rand"
  "encoding/binary"
  "math"
)

var (
  AllElements [9]string = [9]string{"Common", "Air", "Fire", "Earth", "Water", "Void", "Mallom", "Noise", "Resonance"}
  ElemSigns [9]string = [9]string{"✳️ ", "☁️ ", "🔥", "⛰ ", "🧊", "🌑", "🩸", "🎶", "🌟"}
)

func RNF() float64 {
  x := (time.Now().UnixNano())
  in_bytes := make([]byte, 8)
  binary.LittleEndian.PutUint64(in_bytes, uint64(x))
  hsum := sha512.Sum512(in_bytes)
  sum := binary.BigEndian.Uint64(hsum[:])
  return rand.New(rand.NewSource( int64(sum) )).Float64()
}

func RNDElem() string { return AllElements[ rand.New(rand.NewSource(time.Now().UnixNano())).Intn( len(AllElements)-1 ) +1 ] }
func ElemToInt(elem string) int { for i, each := range AllElements {if elem == each {return i}} ; return -1 }
func ES(elem string) string { return ElemSigns[ElemToInt(elem)] }
func Log1479(a float64) float64 { return math.Log2(math.Abs(a)+1)/math.Log2(1.479) }
func SRNF() float64 { return ( 4*RNF() - 3*RNF() + RNF() - 2*RNF() )/( 2*RNF() - 4*RNF() + 3*RNF() - RNF() ) } //NU
func Sign(a float64) float64 { if a != 0 {return a/math.Abs(a)} else {return 0} }

func ChancedRound(a float64) int {
  b,l:=math.Ceil(a),math.Floor(a)
  c:=math.Abs(math.Abs(a)-math.Abs(math.Min(b, l)))
  if a<0 {c = 1-c}
  // if a!= 0 {fmt.Printf("\n %1.2f - between %1.0f and %1.0f - with chance %1.2f",a, b, l, c)}
  if RNF() < c {return int(b)} else {return int(l)}
  return 0
}
