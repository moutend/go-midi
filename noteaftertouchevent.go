package midi

import "fmt"

type NoteAfterTouchEvent struct {
	deltaTime *DeltaTime
	channel   uint8
	note      Note
	velocity  uint8
}

func (e *NoteAfterTouchEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *NoteAfterTouchEvent) String() string {
	return fmt.Sprintf("&NoteAfterTouchEvent{}")
}
