package midi

import "fmt"

type SetTempoEvent struct {
	deltaTime *DeltaTime
	tempo     []byte
}

func (e *SetTempoEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *SetTempoEvent) String() string {
	return fmt.Sprintf("&SetTempoEvent{}")
}
