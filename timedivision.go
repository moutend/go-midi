package midi

import "fmt"

// TimeDivision represents the time division used to decode the delta time.
type TimeDivision struct {
	value uint16
}

// String returns string representation of time division.
func (t *TimeDivision) String() string {
	if t.value == 0 {
		return "&TimeDivision{bpm: 120}"
	}
	if t.value < 32768 {
		return fmt.Sprintf("&TimeDivision{bpm: %d}", t.value)
	}

	frames := (t.value & 0x7F00) >> 8
	ticks := t.value & 0x00FF

	return fmt.Sprintf("&TimeDivision{frames: %v, ticks: %v}", frames, ticks)
}

// Serialize serializes time division.
func (t *TimeDivision) Serialize() []byte {
	if t.value == 0 {
		return []byte{0x00, 0x78}
	}

	bs := make([]byte, 2)
	bs[0] = byte(t.value >> 8)
	bs[1] = byte(t.value & 0xff)

	return bs
}

// SetBPM sets time division value as BPM.
func (t *TimeDivision) SetBPM(bpm int) error {
	if bpm >= 0x8000 {
		return fmt.Errorf("midi: BPM must be less than 32768")
	}
	t.value = uint16(bpm)

	return nil
}

// BPM returns time division as beat per minute.
func (t *TimeDivision) BPM() (uint16, error) {
	if t.value == 0 {
		return 120, nil
	}
	if t.value < 0x8000 {
		return t.value, nil
	}

	return 0, fmt.Errorf("midi: cannot retrieve value as BPM (%v)", t.value)
}

// SetFPS sets time division value as frames per second.
func (t *TimeDivision) SetFPS(frames, ticks uint16) error {
	t.value = 0x8000 + (frames << 8) + ticks

	return nil
}

// FPS retuns time division as frames per second
func (t *TimeDivision) FPS() (frames, ticks uint16, err error) {
	if t.value <= 0x8000 {
		return frames, ticks, fmt.Errorf("midi: cannot retrieve time division as FPS (%v)", t.value)
	}
	frames = (t.value >> 8) & 0x7f
	ticks = t.value & 0xff

	return frames, ticks, nil
}
