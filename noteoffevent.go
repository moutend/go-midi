package midi

import (
	"fmt"
)

// NoteOffEvent corresponds to note-off event (0x80) in MIDI.
type NoteOffEvent struct {
	deltaTime *DeltaTime
	channel   uint8
	note      Note
	velocity  uint8
}

// DeltaTime returns delta time of this event.
func (e *NoteOffEvent) DeltaTime() *DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &DeltaTime{}
	}
	return e.deltaTime
}

// String returns string representation of this event.
func (e *NoteOffEvent) String() string {
	return fmt.Sprintf("&NoteOffEvent{}")
}
