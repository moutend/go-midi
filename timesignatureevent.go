package midi

import "fmt"

type TimeSignatureEvent struct {
	deltaTime *DeltaTime
	tempo     []byte
}

func (e *TimeSignatureEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *TimeSignatureEvent) String() string {
	return fmt.Sprintf("&TimeSignatureEvent{}")
}
