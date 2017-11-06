package midi

import "fmt"

type MIDIChannelPrefixEvent struct {
	deltaTime *DeltaTime
	channel   uint8
}

func (e *MIDIChannelPrefixEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *MIDIChannelPrefixEvent) String() string {
	return fmt.Sprintf("&MIDIChannelPrefixEvent{channel: %v}", e.channel)
}
