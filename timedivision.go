package midi

import "fmt"

// TimeDivision represents time division for MIDI header.
type TimeDivision struct {
	value uint16
}

func (t *TimeDivision) String() string {
	// assumes 120 bpm
	if t.value == 0 {
		return "default (120 bpm)"
	}
	if t.value < 32768 {
		return fmt.Sprintf("%d bpm", t.value)
	}
	fps := t.value & 0x7F00
	ticks := t.value & 0x00FF
	return fmt.Sprintf("%v ticks in %v fps", ticks, fps)
}

func (t *TimeDivision) Set(rawdata uint16) {
	t.value = rawdata
}

func (t *TimeDivision) SetBPM(bpm int) error {
	if bpm >= 0x8000 {
		return fmt.Errorf("midi: invalid BPM %d", bpm)
	}
	t.value = uint16(bpm)

	return nil
}

// Serialize serializes time division.
func (t *TimeDivision) Serialize() []byte {
	bs := make([]byte, 2)
	bs[0] = byte(t.value >> 8)
	bs[1] = byte(t.value & 0xff)

	return bs
}
