package midi

import "fmt"

type ControllerEvent struct {
	deltaTime *DeltaTime
	channel   uint8
	control   uint8
	value     uint8
}

func (e *ControllerEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *ControllerEvent) String() string {
	return fmt.Sprintf("&ControllerEvent{}")
}
