package midi

import "fmt"

// NoteAfterTouchEvent corresponds to note after touch event.
type NoteAfterTouchEvent struct {
	deltaTime     *DeltaTime
	runningStatus bool
	channel       uint8
	note          Note
	velocity      uint8
}

// DeltaTime returns delta time of note after touch event.
func (e *NoteAfterTouchEvent) DeltaTime() *DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &DeltaTime{}
	}
	return e.deltaTime
}

// Serialize serializes note after touch event.
func (e *NoteAfterTouchEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, e.DeltaTime().Quantity().Value()...)
	bs = append(bs, NoteAfterTouch+e.channel)
	bs = append(bs, byte(e.note), e.velocity)

	return bs
}

// SetRunningStatus sets running status.
func (e *NoteAfterTouchEvent) SetRunningStatus(status bool) {
	e.runningStatus = status
}

// RunningStatus returns running status.
func (e *NoteAfterTouchEvent) RunningStatus() bool {
	return e.runningStatus
}

// SetChannel sets channel.
func (e *NoteAfterTouchEvent) SetChannel(channel uint8) error {
	if channel > 0x0f {
		return fmt.Errorf("midi: maximum channel number is 15 (0x0f)")
	}
	e.channel = channel

	return nil
}

// Channel returns channel.
func (e *NoteAfterTouchEvent) Channel() uint8 {
	return e.channel
}

// SetNote sets note.
func (e *NoteAfterTouchEvent) SetNote(note Note) error {
	if note > 0x7f {
		return fmt.Errorf("midi: maximum value of note is 127 (0x7f)")
	}
	e.note = note

	return nil
}

// Note returns note.
func (e *NoteAfterTouchEvent) Note() Note {
	return e.note
}

// SetVelocity sets velocity.
func (e *NoteAfterTouchEvent) SetVelocity(velocity uint8) error {
	if velocity > 0x7f {
		return fmt.Errorf("midi: maximum value of velocity is 127 (0x7f)")
	}
	e.velocity = velocity

	return nil
}

// Velocity returns velocity.
func (e *NoteAfterTouchEvent) Velocity() uint8 {
	return e.velocity
}

// String returns string representation of note after touch event.
func (e *NoteAfterTouchEvent) String() string {
	return fmt.Sprintf("&NoteAfterTouchEvent{channel: %v, note: %v, velocity: %v}", e.channel, e.note, e.velocity)
}

// NewNoteAfterTouchEvent returns NoteAfterTouchEvent with the given parameter.
func NewNoteAfterTouchEvent(deltaTime *DeltaTime, channel uint8, note Note, velocity uint8) (*NoteAfterTouchEvent, error) {
	var err error

	event := &NoteAfterTouchEvent{}
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
