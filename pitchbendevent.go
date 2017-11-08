package midi

import "fmt"

type PitchBendEvent struct {
	deltaTime *DeltaTime
	channel   uint8
	note      Note
	velocity  uint16
}

func (e *PitchBendEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *PitchBendEvent) String() string {
	return fmt.Sprintf("&PitchBendEvent{}")
}
