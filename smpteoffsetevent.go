package midi

import "fmt"

type SMPTEOffsetEvent struct {
	deltaTime *DeltaTime
	tempo     []byte
}

func (e *SMPTEOffsetEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *SMPTEOffsetEvent) String() string {
	return fmt.Sprintf("&SMPTEOffsetEvent{}")
}
