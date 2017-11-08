package midi

import "fmt"

type ProgramChangeEvent struct {
	deltaTime *DeltaTime
	channel   byte
	program   uint8
}

func (e *ProgramChangeEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *ProgramChangeEvent) String() string {
	return fmt.Sprintf("&ProgramChangeEvent{}")
}
