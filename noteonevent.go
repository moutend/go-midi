package midi

import "fmt"

type NoteOnEvent struct {
	deltaTime *DeltaTime
	channel   uint8
	note      uint8
	velocity  uint8
}

func (e *NoteOnEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *NoteOnEvent) String() string {
	return fmt.Sprintf("&NoteOnEvent{}")
}
