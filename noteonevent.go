package midi

import "fmt"

// NoteOnEvent corresponds to text event (0x90) in MIDI.
type NoteOnEvent struct {
	deltaTime *DeltaTime
	channel   byte
	note      Note
	velocity  byte
}

// DeltaTime returns delta time of this event as DeltaTime.
func (e *NoteOnEvent) DeltaTime() *DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &DeltaTime{}
	}
	return e.deltaTime
}

// String returns string representation of this event.
func (e *NoteOnEvent) String() string {
	return fmt.Sprintf("&NoteOnEvent{}")
}

// Serialize serializes this event.
func (e *NoteOnEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, e.DeltaTime().Quantity().Value()...)
	bs = append(bs, 0x90+byte(e.channel))
	bs = append(bs, byte(e.note), e.velocity)

	return bs
}

// SetChannel sets channel of this event.
func (e *NoteOnEvent) SetChannel(channel byte) error {
	if channel > 0x0f {
		return fmt.Errorf("midi: maximum channel number is 15 (0x0f)")
	}
	e.channel = channel

	return nil
}

// SetNote sets note for this event.
func (e *NoteOnEvent) SetNote(note Note) error {
	if note > 0x7f {
		return fmt.Errorf("midi: maximum value of note is 127 (0x7f)")
	}
	e.note = note

	return nil
}

// SetVelocity sets velocity of this event.
func (e *NoteOnEvent) SetVelocity(velocity byte) error {
	if velocity > 0x7f {
		return fmt.Errorf("midi: maximum value of velocity is 127 (0x7f)")
	}
	e.velocity = velocity

	return nil
}

// NewNoteOnEvent returns NoteOnEvent with the given parameter.
func NewNoteOnEvent(deltaTime *DeltaTime, channel byte, note Note, velocity byte) (*NoteOnEvent, error) {
	var err error

	event := &NoteOnEvent{}
	event.deltaTime = deltaTime

	err = event.SetChannel(channel)
	if err != nil {
		return nil, err
	}
	err = event.SetNote(note)
	if err != nil {
		return nil, err
	}
	err = event.SetVelocity(velocity)
	if err != nil {
		return nil, err
	}
	return event, nil
}
