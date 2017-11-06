package midi

import "fmt"

type AlienEvent struct {
	deltaTime     *DeltaTime
	metaEventType MetaEventType
	data          []byte
}

func (e *AlienEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *AlienEvent) String() string {
	return fmt.Sprintf("&AlienEvent{metaEventType: %x data: %v bytes}", e.metaEventType, len(e.data))
}
