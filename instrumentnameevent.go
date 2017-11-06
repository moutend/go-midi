package midi

import "fmt"

type InstrumentNameEvent struct {
	deltaTime *DeltaTime
	text      []byte
}

func (e *InstrumentNameEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *InstrumentNameEvent) String() string {
	return fmt.Sprintf("&InstrumentNameEvent{}")
}
