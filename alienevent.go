package midi

import "fmt"

type AlienEvent struct {
	deltaTime     *DeltaTime
	metaEventType byte
	data          []byte
}

func (e *AlienEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *AlienEvent) String() string {
	return fmt.Sprintf("&AlienEvent{metaEventType: 0x%x data: %v bytes}", e.metaEventType, len(e.data))
}
