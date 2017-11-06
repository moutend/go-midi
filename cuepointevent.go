package midi

import "fmt"

type CuePointEvent struct {
	deltaTime *DeltaTime
	text      []byte
}

func (e *CuePointEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *CuePointEvent) String() string {
	return fmt.Sprintf("&CuePointEvent{}")
}
