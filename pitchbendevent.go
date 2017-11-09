package midi

import "fmt"

// PitchBendEvent corresponds to pitch bend event (0xe0) in MIDI.
type PitchBendEvent struct {
	deltaTime *DeltaTime
	channel   uint8
	pitch     uint16
}

// DeltaTime returns delta time of this event.
func (e *PitchBendEvent) DeltaTime() *DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &DeltaTime{}
	}
	return e.deltaTime
}

// String returns string representation of this event.
func (e *PitchBendEvent) String() string {
	return fmt.Sprintf("&PitchBendEvent{channel: %v, pitch: %v}", e.channel, e.pitch)
}

// Serialize serializes this event.
func (e *PitchBendEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, e.DeltaTime().Quantity().Value()...)
	bs = append(bs, PitchBend+e.channel)

	lsb := byte(e.pitch >> 7)
	msb := byte(e.pitch & 0x7f)
	bs = append(bs, msb, lsb)

	return bs
}

// SetChannel sets channel of this event.
func (e *PitchBendEvent) SetChannel(channel uint8) error {
	if channel > 0x0f {
		return fmt.Errorf("midi: maximum channel number is 15 (0x0f)")
	}
	e.channel = channel

	return nil
}

// SetPitch sets pitch of this event.
func (e *PitchBendEvent) SetPitch(pitch uint16) error {
	if pitch > 0x3fff {
		return fmt.Errorf("midi: maximum value of pitch is 16384 (0x3fff)")
	}
	e.pitch = pitch

	return nil
}

// NewPitchBendEvent returns PitchBendEvent with the given parameter.
func NewPitchBendEvent(deltaTime *DeltaTime, channel uint8, pitch uint16) (*PitchBendEvent, error) {
	var err error

	event := &PitchBendEvent{}
	event.deltaTime = deltaTime

	err = event.SetChannel(channel)
	if err != nil {
		return nil, err
	}
	err = event.SetPitch(pitch)
	if err != nil {
		return nil, err
	}
	return event, nil
}
