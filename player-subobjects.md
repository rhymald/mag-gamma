# Environment
1. **PowerStates**, []PowerState - long term empowerings.
- **Curses**, []Curse - long term catastrophes.

### PowerState
Energetic sources in the world, affecting creation.
1. **Description**, string - player feelings reply.
2. **Nature**, []Stream - positives only.
3. **Area**, float64 - raduis of effect.
4. **XYZs**, [][3]float64 - position coordinates.
5. **Concentrated**, bool - nature of power distribution.

### Curse
Energetic sources in the world, affecting destruction.
1. **Description**, string - player feelings reply.
2. **Nature**, []Stream - negatives only.
3. **Area**, float64 - raduis of effect.
4. **XYZs**, [][3]float64 - position coordinates.
5. **Concentrated**, bool - nature of power distribution.

# Player
1. **Limits**, - maximal abilities, static between upgrades.
2. **XYZ**, [3]float64 - positions.
3. **ElementState**, - env-collided state.
4. **Name**, - name.
5. **Pool**, []Dot - dots.
- **Leveling**, - achieving.
- **Exp**, - current progress.

### Limits
1. **Health**, float64 - current health.
2. **Vitality**, float64 - maximum hp.
3. **Class**, float64 - initial randomized vector of self-developing.
4. **Capacity**, int - max dots capacity.

### ElementState
1. **Pos**, [9]Stream - sources sum
2. **Ext**, [9]Stream - sources complex
- Int, [9]Stream - internal sum
- Neg, [9]Stream - curses sum
- **Resistances**, [9]float64 - internal complex

### Leveling
- **Sprout**, Stream - new sprout stats.
- **Gain**, Stream - new stats.
- **Attunement**, []float64 - possible brandings.

### Exp
- **Strouting**, [3]float64 - creation, alteration, destruction.
- **Gaining**, [3]float64 - creation, alteration, destruction.
- **Branding**, [9]float64 - each element readiness.

# Primitives

### Stream
1. **Creation**, float64 - stream long,
2. **Alteration**, float64 - stream weight,
3. **Destruction**, float64 - stream power.
4. **Element**, string - stream affinity.

### Dot
1. **Weight**, float64 - dot value.
2. **Element**, float64 - dot element.

### Flock
- **Streams**, []int - streams attached.
- **State**, [9]Stream - effect.

### Scheme
- **Chain**
