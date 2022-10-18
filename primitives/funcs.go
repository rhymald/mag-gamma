package primitives

import (
  "time"
  "crypto/sha512"
  "math/rand"
  "encoding/binary"
)

var (
  AllElements [9]string = [9]string{"Common", "Air", "Fire", "Earth", "Water", "Void", "Mallom", "Noise", "Resonance"}
  ElemSigns [9]string = [9]string{"✳️ ", "☁️ ", "🔥", "⛰ ", "🧊", "🌑", "🩸", "🎶", "🌟"}
)

func RNF() float64 {
  x := (time.Now().UnixNano())
  // var in_bytes []byte = big.NewInt(x).Bytes()
  in_bytes := make([]byte, 8)
  binary.LittleEndian.PutUint64(in_bytes, uint64(x))

  hsum := sha512.Sum512(in_bytes)
  sum := binary.BigEndian.Uint64(hsum[:])
  return rand.New(rand.NewSource( int64(sum) )).Float64()
}
