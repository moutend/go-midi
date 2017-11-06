package midi

import "fmt"

type MIDIPortPrefixEvent struct {
	deltaTime *DeltaTime
	port      uint8
}

func (e *MIDIPortPrefixEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *MIDIPortPrefixEvent) String() string {
	return fmt.Sprintf("&MIDIPortPrefixEvent{port: %v}", e.port)
}
