package midi

import (
	"fmt"
)

type NoteOffEvent struct {
	deltaTime *DeltaTime
	channel   uint8
	note      uint8
	velocity  uint8
}

func (e *NoteOffEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *NoteOffEvent) String() string {
	return fmt.Sprintf("&NoteOffEvent{}")
}
