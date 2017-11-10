package midi

import "fmt"

// NoteOnEvent corresponds to note on event.
type NoteOnEvent struct {
	deltaTime *DeltaTime
	channel   uint8
	note      Note
	velocity  uint8
}

// DeltaTime returns delta time of note on event.
func (e *NoteOnEvent) DeltaTime() *DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &DeltaTime{}
	}
	return e.deltaTime
}

// String returns string representation of note on event.
func (e *NoteOnEvent) String() string {
	return fmt.Sprintf("&NoteOnEvent{channel: %v, note: %v, velocity: %v}", e.channel, e.note, e.velocity)
}

// Serialize serializes note on event.
func (e *NoteOnEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, e.DeltaTime().Quantity().Value()...)
	bs = append(bs, NoteOn+e.channel)
	bs = append(bs, byte(e.note), e.velocity)

	return bs
}

// SetChannel sets channel.
func (e *NoteOnEvent) SetChannel(channel uint8) error {
	if channel > 0x0f {
		return fmt.Errorf("midi: maximum channel number is 15 (0x0f)")
	}
	e.channel = channel

	return nil
}

// Channel returns channel.
func (e *NoteOnEvent) Channel() uint8 {
	return e.channel
}

// SetNote sets note.
func (e *NoteOnEvent) SetNote(note Note) error {
	if note > 0x7f {
		return fmt.Errorf("midi: maximum value of note is 127 (0x7f)")
	}
	e.note = note

	return nil
}

// Note returns note.
func (e *NoteOnEvent) Note() Note {
	return e.note
}

// SetVelocity sets velocity.
func (e *NoteOnEvent) SetVelocity(velocity uint8) error {
	if velocity > 0x7f {
		return fmt.Errorf("midi: maximum value of velocity is 127 (0x7f)")
	}
	e.velocity = velocity

	return nil
}

// Velocity returns velocity.
func (e *NoteOnEvent) Velocity() uint8 {
	return e.velocity
}

// NewNoteOnEvent returns NoteOnEvent with the given parameter.
func NewNoteOnEvent(deltaTime *DeltaTime, channel uint8, note Note, velocity uint8) (*NoteOnEvent, error) {
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
