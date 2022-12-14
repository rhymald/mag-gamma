## Streams
```yaml
Class: float64
Herald: float64
Bender: float64
Resistances: [9]float64
Internal State: [9]Stream
List:
  - Stream:
      Creation: float64
      Alteration: float64
      Destruction: float64
      Aexp: [9]float64 #TBD
      Heat:
        Current: float64
        Threshold: float64
  - Stream: ...
```

#### Elementals State
```yaml
External: [9]Stream
Empowered: [9]Stream
```

#### Pool
consumable, transferable.
```yaml
MaxVol: float64
Dots:
  - Dot:
      Element: string
      Weight: float64
  - Dot: ...
```

#### Flocks
```yaml
- Indexes: [all]int
  Fatal: Stream
  Effeciency:
    Cre: float64
    ...: ... # All other
- Indexes: [some]int
  ...: ...
```

#### Skills
```yaml
- Description: Flaming heart
  Chain:
  - Flock 1
  - Socket
  - Flock 2
- Description: Well  
  Chain:
  - Socket in
  - Flock 3  
- Description: Fire channel
  Cain:
  - Flock 1
  - Socket
  - Flock 4
  - Socket out
- ...:
  - ...
```

#### Physical #TBD
```yaml
Position:
  X: float64
  Y: float64
  Z: float64
Body:
  Stream:
    Complexion: float64
    Endurance: float64
    Strength: float64
    Health:
      Max: float64
      Current: float64
Physical State:
  Stream:
    Complexion: float64
    Endurance: float64
    Strength: float64
    Stamina:
      Max: float64
      Current: float64  
Inventory:
  Shoulder: Stream
  Shoulder: Stream
  Back: Stream
  Chest: Stream
  Belt:
    Stream:
      Volume: float64
      Speed: float64
      Burden: float64
      Capacity:
        Max: float64
        Current: float64
```
