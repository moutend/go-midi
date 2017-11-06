package midi

import "fmt"

type MarkerEvent struct {
	deltaTime *DeltaTime
	text      []byte
}

func (e *MarkerEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *MarkerEvent) String() string {
	return fmt.Sprintf("&MarkerEvent{text: \"%v\"}", string(e.text))
}
