package midi

import "fmt"

type SystemExclusiveEvent struct {
	deltaTime *DeltaTime
	data      []byte
}

func (e *SystemExclusiveEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *SystemExclusiveEvent) String() string {
	return fmt.Sprintf("&SystemExclusiveEvent{data: %v bytes}", len(e.data))
}
